package serum_dex

import (
	"encoding/base64"
	"encoding/binary"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	SetProgramID(solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin"))
	os.Exit(m.Run())
}

func TestInstructionDecode_CancelOrderV2(t *testing.T) {
	CancelOrdersAndSettleMsg :=
		"AQADDQpz5x/t0hNl7QruhPzk4rIGR/001ey9oRXwI9JjP4d4MZCYjBIPabdFGOFRg4aCuYIvpX0/Up2H0gsFDOvAKxptYTUM2" +
			"1S5Es4XParbLdSX6tunfXUP9th1OPASGbtY2PiI2PpWo39ogKOkNY5k/Nfr0Wt/9KEDGrcUmrxuhRiSbi3T/D9J3OFx/jgQJR" +
			"Q+s2+9GBNxoSwkDUIUrgnJSVYCOaxQQvf9r8PyaRmnlkSN6BQs0u7SDcCOz4CLebCdaSgViGOph6mtTc/i5/V7k/Malr7/YxBe" +
			"zcjPsCTq/LNI8RqttgK/juTyfqv8uSiqbkGWGETaSk2Kb5xsFcNlQQWSbGlLqfiZ41yfAAH/BysfePaWsZCjnts5xkiSRTT14" +
			"A0x4LQTafOL58gL65FcTUQQaY0+tE9OQfrFVtgS5kpxhQ8tbgKkevgk0Jq2ncQtcMsoy/okn7fuV7nSVsEnYu+pgdiCRuMX9X" +
			"DIP87eCqZqkDn14UiDO4FzWxHK14PVPQbd9uHXZaGT2cvhRs7reawctIXtX1s3kTqM9YV+/wCpRT/gXmftFbCh8k/tkXlqLbwa" +
			"AlSoCOY59nR43pTxwWwFCgYBAgMEAAUZAAsAAAAAAAAANC/f//////89DAAAAAAAAAoGAQIDBAAFGQALAAAAAAAAADMv3/////" +
			"//HgwAAAAAAAAKBgECAwQABRkACwAAAAAAAAAyL9////////4LAAAAAAAACgYBAgMEAAUZAAsAAAAAAAAAMS/f///////eCwAA" +
			"AAAAAAoJAQQABgcICQsMBQAFAAAA"

	data, err := base64.StdEncoding.DecodeString(CancelOrdersAndSettleMsg)
	require.NoError(t, err)

	msg := new(solana.Message)

	err = msg.UnmarshalWithDecoder(bin.NewBorshDecoder(data))
	require.NoError(t, err)
	require.Equal(t, 5, len(msg.Instructions))

	in := msg.Instructions[1]

	inst, err := DecodeInstruction(in.ResolveInstructionAccounts(msg), in.Data)
	require.NoError(t, err)

	spew.Dump(inst)

	ix, ok := inst.Impl.(*CancelOrderV2)
	if !ok {
		t.Errorf("the instruction is not a *CancelOrderV2")
	}
	_ = ix

	require.NotNil(t, ix.GetOpenOrdersAccount())
	require.NotNil(t, ix.GetOwnerAccount())
	require.NotNil(t, ix.GetMarketAccount())
	require.NotNil(t, ix.GetBidsAccount())
	require.NotNil(t, ix.GetAsksAccount())
	require.NotNil(t, ix.GetEventQueueAccount())

	require.Equal(t, solana.MustPublicKeyFromBase58("8R6NtLSyVpgG6f8Z9cLKFcuaJtt3mqHypKGsmRixwyJV"), ix.GetOpenOrdersAccount().PublicKey)
	require.Equal(t, solana.MustPublicKeyFromBase58("hoakwpFB8UoLnPpLC56gsjpY7XbVwaCuRQRMQzN5TVh"), ix.GetOwnerAccount().PublicKey)
	require.Equal(t, solana.MustPublicKeyFromBase58("4LUro5jaPaTurXK737QAxgJywdhABnFAMQkXX4ZyqqaZ"), ix.GetMarketAccount().PublicKey)
	require.Equal(t, solana.MustPublicKeyFromBase58("9gpfTc4zsndJSdpnpXQbey16L5jW2GWcKeY3PLixqU4"), ix.GetEventQueueAccount().PublicKey)
	require.Equal(t, solana.MustPublicKeyFromBase58("8MyQkxux1NnpNqpBbPeiQHYeDbZvdvs7CHmGpciSMWvs"), ix.GetBidsAccount().PublicKey)
	require.Equal(t, solana.MustPublicKeyFromBase58("HjB8zKe9xezDrgqXCSjCb5F7dMC9WMwtZoT7yKYEhZYV"), ix.GetAsksAccount().PublicKey)

	require.True(t, ix.Args.Side == SideBid)

	orderID, err := decimal.NewFromString("57240246860720736513843")
	require.NoError(t, err)
	spew.Dump(orderID)

	{
		dec := bin.NewBorshDecoder(in.Data)
		{
			got, err := dec.ReadUint8()
			require.NoError(t, err)
			spew.Dump(got)
			require.Equal(t, uint8(0), got)
		}
		got, err := dec.ReadUint32(binary.LittleEndian)
		require.NoError(t, err)
		spew.Dump(got)
		require.Equal(t, Instruction_CancelOrderV2, got)

		{
			got, err := dec.ReadUint32(binary.LittleEndian)
			require.NoError(t, err)
			require.Equal(t, SideBid, Side(got))
			spew.Dump(got)
			// dec.SkipBytes(3)
		}

		instructionParamData, err := dec.ReadNBytes(dec.Remaining())
		require.NoError(t, err)
		spew.Dump([]byte(instructionParamData))
		spew.Dump(bin.FormatByteSlice([]byte(instructionParamData)))
		oo := orderID.BigInt().FillBytes(make([]byte, 16))
		bin.ReverseBytes(oo)
		spew.Dump(bin.FormatByteSlice(oo))
		require.Equal(t, oo, instructionParamData)
	}

	require.Equal(t, orderID.BigInt(), ix.Args.OrderId.BigInt())
	spew.Dump(ix.Args.OrderId.Hi)
	spew.Dump(ix.Args.OrderId.Lo)
	require.Equal(t, uint64(18446744073707401011), ix.Args.OrderId.Lo)
	require.Equal(t, uint64(3102), ix.Args.OrderId.Hi)
}

func TestInstructionDecode_CreateOpenOrdersNewOrdersAndSettleMsg(t *testing.T) {
	CreateOpenOrdersNewOrdersAndSettleMsg :=
		"AgAFEApz5x/t0hNl7QruhPzk4rIGR/001ey9oRXwI9JjP4d4bi3T/D9J3OFx/jgQJRQ+s2+9GBNxoSwkDUIUrgnJSVYxkJiME" +
			"g9pt0UY4VGDhoK5gi+lfT9SnYfSCwUM68ArGpr33IM/qQm6+EFjv3WCl9+nICT59ibJWro20VS5uUQ9AjmsUEL3/a/D8mkZp5" +
			"ZEjegULNLu0g3Ajs+Ai3mwnWltYTUM21S5Es4XParbLdSX6tunfXUP9th1OPASGbtY2PiI2PpWo39ogKOkNY5k/Nfr0Wt/9KE" +
			"DGrcUmrxuhRiSDTHgtBNp84vnyAvrkVxNRBBpjT60T05B+sVW2BLmSnEoFYhjqYeprU3P4uf1e5PzGpa+/2MQXs3Iz7Ak6vyz" +
			"SPEarbYCv47k8n6r/Lkoqm5BlhhE2kpNim+cbBXDZUEFkmxpS6n4meNcnwAB/wcrH3j2lrGQo57bOcZIkkU09eAAAAAAAAAAA" +
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAan1RcZLFxRIYzJTD1K8X9Y2u4Im6H9ROPb2YoAAAAAhQ8tbgKkevgk0Jq2ncQtcM" +
			"soy/okn7fuV7nSVsEnYu8G3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqamB2IJG4xf1cMg/zt4KpmqQOfXhSIM7gXN" +
			"bEcrXg9U9sKLMetn7/1UUmSgiOvObf6ncpgx9byhg2wueQRNGOvYICwIAATQAAAAAQGlkAQAAAACcDAAAAAAAAIUPLW4CpHr4" +
			"JNCatp3ELXDLKMv6JJ+37le50lbBJ2LvDQQBAAIMBQAPAAAADQwCAQMEBQYHAAgJDgwzAAoAAAAAAAAAXQwAAAAAAAAKAAAAA" +
			"AAAAEhLMAAAAAAAAQAAAAAAAAABAAAAAAAAAP//DQwCAQMEBQYHAAgJDgwzAAoAAAAAAAAAPQwAAAAAAAAWAAAAAAAAADgsaQ" +
			"AAAAAAAQAAAAAAAAABAAAAAAAAAP//DQwCAQMEBQYHAAgJDgwzAAoAAAAAAAAAHgwAAAAAAAAjAAAAAAAAACiqpQAAAAAAAQA" +
			"AAAAAAAABAAAAAAAAAP//DQwCAQMEBQYHAAgJDgwzAAoAAAAAAAAA/gsAAAAAAAAvAAAAAAAAAEgr3AAAAAAAAQAAAAAAAAAB" +
			"AAAAAAAAAP//DQwCAQMEBQYHAAgJDgwzAAoAAAAAAAAA3gsAAAAAAAA8AAAAAAAAACAjFgEAAAAAAQAAAAAAAAABAAAAAAAAA" +
			"P//DQkCAQAICQoHDw4FAAUAAAA="

	data, err := base64.StdEncoding.DecodeString(CreateOpenOrdersNewOrdersAndSettleMsg)
	require.NoError(t, err)

	msg := new(solana.Message)

	err = msg.UnmarshalWithDecoder(bin.NewBorshDecoder(data))
	require.NoError(t, err)
	require.Equal(t, 8, len(msg.Instructions))

	{
		in := msg.Instructions[1]
		inst, err := DecodeInstruction(in.ResolveInstructionAccounts(msg), in.Data)
		require.NoError(t, err)

		spew.Dump(inst)

		ix, ok := inst.Impl.(*InitOpenOrders)
		if !ok {
			t.Errorf("the instruction is not a *InitOpenOrders")
		}

		require.NotNil(t, ix.GetOpenOrdersAccount())
		require.NotNil(t, ix.GetOwnerAccount())
		require.NotNil(t, ix.GetMarketAccount())

		require.Equal(t, solana.MustPublicKeyFromBase58("8R6NtLSyVpgG6f8Z9cLKFcuaJtt3mqHypKGsmRixwyJV"), ix.GetOpenOrdersAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("hoakwpFB8UoLnPpLC56gsjpY7XbVwaCuRQRMQzN5TVh"), ix.GetOwnerAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("4LUro5jaPaTurXK737QAxgJywdhABnFAMQkXX4ZyqqaZ"), ix.GetMarketAccount().PublicKey)
	}
	{
		in := msg.Instructions[2]
		inst, err := DecodeInstruction(in.ResolveInstructionAccounts(msg), in.Data)
		require.NoError(t, err)

		spew.Dump(inst)

		ix, ok := inst.Impl.(*NewOrderV3)
		if !ok {
			t.Errorf("the instruction is not a *NewOrderV3")
		}

		require.Equal(t, solana.MustPublicKeyFromBase58("8R6NtLSyVpgG6f8Z9cLKFcuaJtt3mqHypKGsmRixwyJV"), ix.GetOpenOrdersAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("hoakwpFB8UoLnPpLC56gsjpY7XbVwaCuRQRMQzN5TVh"), ix.GetOwnerAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("4LUro5jaPaTurXK737QAxgJywdhABnFAMQkXX4ZyqqaZ"), ix.GetMarketAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("BRvzjEjphBLVMMq8tEvh4G5o9TTNJ4PSu23CAAdJDKsr"), ix.GetRequestQueueAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("9gpfTc4zsndJSdpnpXQbey16L5jW2GWcKeY3PLixqU4"), ix.GetEventQueueAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("8MyQkxux1NnpNqpBbPeiQHYeDbZvdvs7CHmGpciSMWvs"), ix.GetBidsAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("HjB8zKe9xezDrgqXCSjCb5F7dMC9WMwtZoT7yKYEhZYV"), ix.GetAsksAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("tWQuevB8Rou1HS9a76fjYSQPrDixZMbVzXe2Q1kY5ma"), ix.GetOrderPayerAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("3hUMPnn3WNUbhTBoyXH3wHkWyq85MEZx9LWLTdEEaTef"), ix.GetCoinVaultAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("HEArHmgm9mnsj2u98Ldr4iWSwWvPPUUg8L9fwxT1cTyv"), ix.GetPcVaultAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"), ix.GetSplTokenProgramAccount().PublicKey)
		require.Equal(t, SideBid, ix.Args.Side)
		spew.Dump(ix.Args)
		require.Equal(t, uint64(3165), ix.Args.LimitPrice)
		require.Equal(t, uint64(10), ix.Args.MaxCoinQty)
		require.Equal(t, uint64(3165000), ix.Args.MaxNativePcQtyIncludingFees)
		require.Equal(t, SelfTradeBehaviorCancelProvide, ix.Args.SelfTradeBehavior)
		require.Equal(t, OrderTypeLimit, ix.Args.OrderType)
	}
	{
		in := msg.Instructions[7]
		inst, err := DecodeInstruction(in.ResolveInstructionAccounts(msg), in.Data)
		require.NoError(t, err)

		spew.Dump(inst)

		ix, ok := inst.Impl.(*SettleFunds)
		if !ok {
			t.Errorf("the instruction is not a *SettleFunds")
		}

		require.Equal(t, solana.MustPublicKeyFromBase58("8R6NtLSyVpgG6f8Z9cLKFcuaJtt3mqHypKGsmRixwyJV"), ix.GetOpenOrdersAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("hoakwpFB8UoLnPpLC56gsjpY7XbVwaCuRQRMQzN5TVh"), ix.GetOwnerAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("4LUro5jaPaTurXK737QAxgJywdhABnFAMQkXX4ZyqqaZ"), ix.GetMarketAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("3hUMPnn3WNUbhTBoyXH3wHkWyq85MEZx9LWLTdEEaTef"), ix.GetCoinVaultAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("HEArHmgm9mnsj2u98Ldr4iWSwWvPPUUg8L9fwxT1cTyv"), ix.GetPcVaultAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"), ix.GetSplTokenProgramAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("CQgjkmuDXXJ2WpF6bK9VWZko9T5hwVxAVgvmbV3gkfVe"), ix.GetVaultSignerAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("AraQPzSsE31pdzeTe6Dkvu6g8PvreFW429DAYhsfKYRd"), ix.GetCoinWalletAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("tWQuevB8Rou1HS9a76fjYSQPrDixZMbVzXe2Q1kY5ma"), ix.GetPcWalletAccount().PublicKey)
	}
}

func TestInstructionDecode_CloseOpenOrdersMsg(t *testing.T) {
	CloseOpenOrdersMsg :=
		"AQACBApz5x/t0hNl7QruhPzk4rIGR/001ey9oRXwI9JjP4d4bi3T/D9J3OFx/jgQJRQ+s2+9GBNxoSwkDUIUrgnJSVYxkJiME" +
			"g9pt0UY4VGDhoK5gi+lfT9SnYfSCwUM68ArGoUPLW4CpHr4JNCatp3ELXDLKMv6JJ+37le50lbBJ2LvRT/gXmftFbCh8k/tkX" +
			"lqLbwaAlSoCOY59nR43pTxwWwBAwQBAAACBQAOAAAA"

	data, err := base64.StdEncoding.DecodeString(CloseOpenOrdersMsg)
	require.NoError(t, err)

	msg := new(solana.Message)

	err = msg.UnmarshalWithDecoder(bin.NewBorshDecoder(data))
	require.NoError(t, err)
	require.Equal(t, 1, len(msg.Instructions))

	{
		in := msg.Instructions[0]
		inst, err := DecodeInstruction(in.ResolveInstructionAccounts(msg), in.Data)
		require.NoError(t, err)

		spew.Dump(inst)

		ix, ok := inst.Impl.(*CloseOpenOrders)
		if !ok {
			t.Errorf("the instruction is not a *CloseOpenOrders")
		}

		require.Equal(t, solana.MustPublicKeyFromBase58("8R6NtLSyVpgG6f8Z9cLKFcuaJtt3mqHypKGsmRixwyJV"), ix.GetOpenOrdersAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("hoakwpFB8UoLnPpLC56gsjpY7XbVwaCuRQRMQzN5TVh"), ix.GetOwnerAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("hoakwpFB8UoLnPpLC56gsjpY7XbVwaCuRQRMQzN5TVh"), ix.GetDestinationAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("4LUro5jaPaTurXK737QAxgJywdhABnFAMQkXX4ZyqqaZ"), ix.GetMarketAccount().PublicKey)
	}
}

func TestInstructionDecode_InitOpenOrdersAuthorityMsg(t *testing.T) {
	InitOpenOrdersAuthorityMsg :=
		"AgEDBgpz5x/t0hNl7QruhPzk4rIGR/001ey9oRXwI9JjP4d43BTiskuLijN0Im9GJbaYltfYM7EWlK4B1p3CrtS9LytnTR0fu" +
			"hJxv9yK2x3vBF9FKqGptmtvvRKj5Z2Yt80B+jGQmIwSD2m3RRjhUYOGgrmCL6V9P1Kdh9ILBQzrwCsaBqfVFxksXFEhjMlMPU" +
			"rxf1ja7gibof1E49vZigAAAACFDy1uAqR6+CTQmradxC1wyyjL+iSft+5XudJWwSdi7zCsqcBb8R316Oq2E/s6bacxNqaZZYv" +
			"YBqQrYp2V4ZVJAQUFAgADBAEFAA8AAAA="

	data, err := base64.StdEncoding.DecodeString(InitOpenOrdersAuthorityMsg)
	require.NoError(t, err)

	msg := new(solana.Message)

	err = msg.UnmarshalWithDecoder(bin.NewBorshDecoder(data))
	require.NoError(t, err)
	require.Equal(t, 1, len(msg.Instructions))

	{
		in := msg.Instructions[0]
		inst, err := DecodeInstruction(in.ResolveInstructionAccounts(msg), in.Data)
		require.NoError(t, err)

		spew.Dump(inst)

		ix, ok := inst.Impl.(*InitOpenOrders)
		if !ok {
			t.Errorf("the instruction is not a *InitOpenOrders")
		}

		require.Equal(t, solana.MustPublicKeyFromBase58("7xFCBA6F9xLg56hCMRBeDboJQUFNE5KHfjXGuFukavau"), ix.GetOpenOrdersAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("hoakwpFB8UoLnPpLC56gsjpY7XbVwaCuRQRMQzN5TVh"), ix.GetOwnerAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("4LUro5jaPaTurXK737QAxgJywdhABnFAMQkXX4ZyqqaZ"), ix.GetMarketAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("Fp7AZKKazcqU4N83opDzEZMTnQ6CtRxT118CGNZpziSe"), ix.GetMarketAuthorityAccount().PublicKey)
	}
}

func TestInstructionDecode_CancelByClientIdAndConsumeEventsMsg(t *testing.T) {
	CancelByClientIdAndConsumeEventsMsg :=
		"AgABCQpz5x/t0hNl7QruhPzk4rIGR/001ey9oRXwI9JjP4d4AN+Ypnsy9CYUTPWS+d6mQ62AG7sIO9zzX4rvqJX0TU0xkJiME" +
			"g9pt0UY4VGDhoK5gi+lfT9SnYfSCwUM68ArGm1hNQzbVLkSzhc9qtst1Jfq26d9dQ/22HU48BIZu1jY+IjY+lajf2iAo6Q1jm" +
			"T81+vRa3/0oQMatxSavG6FGJICOaxQQvf9r8PyaRmnlkSN6BQs0u7SDcCOz4CLebCdaZJsaUup+JnjXJ8AAf8HKx949paxkKO" +
			"e2znGSJJFNPXgDTHgtBNp84vnyAvrkVxNRBBpjT60T05B+sVW2BLmSnGFDy1uAqR6+CTQmradxC1wyyjL+iSft+5XudJWwSdi" +
			"7zXpjGB3tXakV2utaBIsepEZdxt2dSrG41gLDbgqbkt4AggGAgMEAQAFDQAMAAAAoIYBAAAAAAAIBQECBQYHBwADAAAA//8="

	data, err := base64.StdEncoding.DecodeString(CancelByClientIdAndConsumeEventsMsg)
	require.NoError(t, err)

	msg := new(solana.Message)

	err = msg.UnmarshalWithDecoder(bin.NewBorshDecoder(data))
	require.NoError(t, err)
	require.Equal(t, 2, len(msg.Instructions))

	{
		in := msg.Instructions[0]
		inst, err := DecodeInstruction(in.ResolveInstructionAccounts(msg), in.Data)
		require.NoError(t, err)

		spew.Dump(inst)

		ix, ok := inst.Impl.(*CancelOrderByClientIdV2)
		if !ok {
			t.Errorf("the instruction is not a *CancelOrderByClientIdV2")
		}

		require.Equal(t, solana.MustPublicKeyFromBase58("14QkUy2jkZU2coqY8mTjL3uGiicTTng1VXrYKK1xk7yW"), ix.GetOpenOrdersAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("hoakwpFB8UoLnPpLC56gsjpY7XbVwaCuRQRMQzN5TVh"), ix.GetOwnerAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("4LUro5jaPaTurXK737QAxgJywdhABnFAMQkXX4ZyqqaZ"), ix.GetMarketAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("8MyQkxux1NnpNqpBbPeiQHYeDbZvdvs7CHmGpciSMWvs"), ix.GetBidsAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("HjB8zKe9xezDrgqXCSjCb5F7dMC9WMwtZoT7yKYEhZYV"), ix.GetAsksAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("9gpfTc4zsndJSdpnpXQbey16L5jW2GWcKeY3PLixqU4"), ix.GetEventQueueAccount().PublicKey)

		require.Equal(t, uint64(100000), *ix.ClientOrderId)
	}
	{
		in := msg.Instructions[1]
		inst, err := DecodeInstruction(in.ResolveInstructionAccounts(msg), in.Data)
		require.NoError(t, err)

		spew.Dump(inst)

		ix, ok := inst.Impl.(*ConsumeEvents)
		if !ok {
			t.Errorf("the instruction is not a *ConsumeEvents")
		}

		require.Equal(t, solana.MustPublicKeyFromBase58("4LUro5jaPaTurXK737QAxgJywdhABnFAMQkXX4ZyqqaZ"), ix.GetMarketAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("9gpfTc4zsndJSdpnpXQbey16L5jW2GWcKeY3PLixqU4"), ix.GetEventQueueAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("AraQPzSsE31pdzeTe6Dkvu6g8PvreFW429DAYhsfKYRd"), ix.GetCoinFeeReceivableAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("tWQuevB8Rou1HS9a76fjYSQPrDixZMbVzXe2Q1kY5ma"), ix.GetPcFeeReceivableAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("14QkUy2jkZU2coqY8mTjL3uGiicTTng1VXrYKK1xk7yW"), ix.GetOpenOrdersAccount().PublicKey)
	}
}

func TestInstructionDecode_PruneMsg(t *testing.T) {
	PruneMsg :=
		"AgECCApz5x/t0hNl7QruhPzk4rIGR/001ey9oRXwI9JjP4d4uHiJfn+GplTj4talKGCSg9UtU/dYC8VBFhP5wi/VDboxkJiMEg" +
			"9pt0UY4VGDhoK5gi+lfT9SnYfSCwUM68ArGm1hNQzbVLkSzhc9qtst1Jfq26d9dQ/22HU48BIZu1jY+IjY+lajf2iAo6Q1jmT8" +
			"1+vRa3/0oQMatxSavG6FGJICOaxQQvf9r8PyaRmnlkSN6BQs0u7SDcCOz4CLebCdaQZjqIBBstRlvGKXFTn/swMxl19ZWB2KVy" +
			"TAWVJDJv12hQ8tbgKkevgk0Jq2ncQtcMsoy/okn7fuV7nSVsEnYu9gYhiIP+KUcsxGrD3ryAixnmdYImbKDZVmwFIINIPFJgEH" +
			"BwIDBAEGAAUHABAAAAD//w=="

	data, err := base64.StdEncoding.DecodeString(PruneMsg)
	require.NoError(t, err)

	msg := new(solana.Message)

	err = msg.UnmarshalWithDecoder(bin.NewBorshDecoder(data))
	require.NoError(t, err)
	require.Equal(t, 1, len(msg.Instructions))

	{
		in := msg.Instructions[0]
		inst, err := DecodeInstruction(in.ResolveInstructionAccounts(msg), in.Data)
		require.NoError(t, err)

		spew.Dump(inst)

		ix, ok := inst.Impl.(*Prune)
		if !ok {
			t.Errorf("the instruction is not a *Prune")
		}

		require.Equal(t, solana.MustPublicKeyFromBase58("4LUro5jaPaTurXK737QAxgJywdhABnFAMQkXX4ZyqqaZ"), ix.GetMarketAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("8MyQkxux1NnpNqpBbPeiQHYeDbZvdvs7CHmGpciSMWvs"), ix.GetBidsAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("HjB8zKe9xezDrgqXCSjCb5F7dMC9WMwtZoT7yKYEhZYV"), ix.GetAsksAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("9gpfTc4zsndJSdpnpXQbey16L5jW2GWcKeY3PLixqU4"), ix.GetEventQueueAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("hoakwpFB8UoLnPpLC56gsjpY7XbVwaCuRQRMQzN5TVh"), ix.GetOwnerAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("RwatxRLAiLNqjzBxB8hWXArgtJ7G84zq7xxoYnrUzZB"), ix.GetOpenOrdersAccount().PublicKey)
		require.Equal(t, solana.MustPublicKeyFromBase58("DR6cxg8D7dXYhtCfuMqBmQFSTRAgcYXUd9SZK7dFgsD3"), ix.GetPruneAuthorityAccount().PublicKey)
	}
}
