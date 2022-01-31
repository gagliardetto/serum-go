package serum_dex

import (
	"encoding/base64"
	"encoding/binary"
	"testing"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestInstructionDecode_CancelOrderV2(t *testing.T) {
	SetProgramID(solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin"))
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
	// SetProgramID(solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin"))
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
	}
}
