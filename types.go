// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package serumgo

import (
	"encoding/binary"

	ag_binary "github.com/gagliardetto/binary"
	bin "github.com/gagliardetto/binary"
)

type InitializeMarketInstruction struct {
	// In the matching engine, all prices and balances are integers.
	// This only works if the smallest representable quantity of the coin
	// is at least a few orders of magnitude larger than the smallest representable
	// quantity of the price currency. The internal representation also relies on
	// on the assumption that every order will have a (quantity x price) value that
	// fits into a u64.
	//
	// If these assumptions are problematic, rejigger the lot sizes.
	CoinLotSize      uint64
	PcLotSize        uint64
	FeeRateBps       uint16
	VaultSignerNonce uint64
	PcDustThreshold  uint64
}

func (obj InitializeMarketInstruction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CoinLotSize` param:
	err = encoder.Encode(obj.CoinLotSize)
	if err != nil {
		return err
	}
	// Serialize `PcLotSize` param:
	err = encoder.Encode(obj.PcLotSize)
	if err != nil {
		return err
	}
	// Serialize `FeeRateBps` param:
	err = encoder.Encode(obj.FeeRateBps)
	if err != nil {
		return err
	}
	// Serialize `VaultSignerNonce` param:
	err = encoder.Encode(obj.VaultSignerNonce)
	if err != nil {
		return err
	}
	// Serialize `PcDustThreshold` param:
	err = encoder.Encode(obj.PcDustThreshold)
	if err != nil {
		return err
	}
	return nil
}

func (obj *InitializeMarketInstruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CoinLotSize`:
	err = decoder.Decode(&obj.CoinLotSize)
	if err != nil {
		return err
	}
	// Deserialize `PcLotSize`:
	err = decoder.Decode(&obj.PcLotSize)
	if err != nil {
		return err
	}
	// Deserialize `FeeRateBps`:
	err = decoder.Decode(&obj.FeeRateBps)
	if err != nil {
		return err
	}
	// Deserialize `VaultSignerNonce`:
	err = decoder.Decode(&obj.VaultSignerNonce)
	if err != nil {
		return err
	}
	// Deserialize `PcDustThreshold`:
	err = decoder.Decode(&obj.PcDustThreshold)
	if err != nil {
		return err
	}
	return nil
}

type NewOrderInstructionV1 struct {
	Side       Side
	LimitPrice uint64
	MaxQty     uint64
	OrderType  OrderType
	ClientId   uint64
}

func (obj NewOrderInstructionV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Side` param:
	err = encoder.WriteUint32(uint32(obj.Side), bin.LE)
	if err != nil {
		return err
	}
	// Serialize `LimitPrice` param:
	err = encoder.Encode(obj.LimitPrice)
	if err != nil {
		return err
	}
	// Serialize `MaxQty` param:
	err = encoder.Encode(obj.MaxQty)
	if err != nil {
		return err
	}
	// Serialize `OrderType` param:
	err = encoder.WriteUint32(uint32(obj.OrderType), bin.LE)
	if err != nil {
		return err
	}
	// Serialize `ClientId` param:
	err = encoder.Encode(obj.ClientId)
	if err != nil {
		return err
	}
	return nil
}

func (obj *NewOrderInstructionV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Side`:
	{
		tmp, err := decoder.ReadUint32(bin.LE)
		if err != nil {
			return err
		}
		obj.Side = Side(tmp)
	}
	// Deserialize `LimitPrice`:
	err = decoder.Decode(&obj.LimitPrice)
	if err != nil {
		return err
	}
	// Deserialize `MaxQty`:
	err = decoder.Decode(&obj.MaxQty)
	if err != nil {
		return err
	}
	// Deserialize `OrderType`:
	{
		tmp, err := decoder.ReadUint32(bin.LE)
		if err != nil {
			return err
		}
		obj.OrderType = OrderType(tmp)
	}
	// Deserialize `ClientId`:
	err = decoder.Decode(&obj.ClientId)
	if err != nil {
		return err
	}
	return nil
}

type CancelOrderInstructionV2 struct {
	Side    Side
	OrderId ag_binary.Uint128
}

func (obj CancelOrderInstructionV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Side` param:
	err = encoder.WriteUint32(uint32(obj.Side), bin.LE)
	if err != nil {
		return err
	}
	// Serialize `OrderId` param:
	err = encoder.Encode(obj.OrderId)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CancelOrderInstructionV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Side`:
	{
		tmp, err := decoder.ReadUint32(bin.LE)
		if err != nil {
			return err
		}
		obj.Side = Side(tmp)
	}
	// Deserialize `OrderId`:
	obj.OrderId, err = decoder.ReadUint128(binary.LittleEndian)
	if err != nil {
		return err
	}
	return nil
}

type NewOrderInstructionV2 struct {
	Side              Side
	LimitPrice        uint64
	MaxQty            uint64
	OrderType         OrderType
	ClientId          uint64
	SelfTradeBehavior SelfTradeBehavior
}

func (obj NewOrderInstructionV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Side` param:
	err = encoder.WriteUint32(uint32(obj.Side), bin.LE)
	if err != nil {
		return err
	}
	// Serialize `LimitPrice` param:
	err = encoder.Encode(obj.LimitPrice)
	if err != nil {
		return err
	}
	// Serialize `MaxQty` param:
	err = encoder.Encode(obj.MaxQty)
	if err != nil {
		return err
	}
	// Serialize `OrderType` param:
	err = encoder.WriteUint32(uint32(obj.OrderType), bin.LE)
	if err != nil {
		return err
	}
	// Serialize `ClientId` param:
	err = encoder.Encode(obj.ClientId)
	if err != nil {
		return err
	}
	// Serialize `SelfTradeBehavior` param:
	err = encoder.WriteUint32(uint32(obj.SelfTradeBehavior), bin.LE)
	if err != nil {
		return err
	}
	return nil
}

func (obj *NewOrderInstructionV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Side`:
	{
		tmp, err := decoder.ReadUint32(bin.LE)
		if err != nil {
			return err
		}
		obj.Side = Side(tmp)
	}
	// Deserialize `LimitPrice`:
	err = decoder.Decode(&obj.LimitPrice)
	if err != nil {
		return err
	}
	// Deserialize `MaxQty`:
	err = decoder.Decode(&obj.MaxQty)
	if err != nil {
		return err
	}
	// Deserialize `OrderType`:
	{
		tmp, err := decoder.ReadUint32(bin.LE)
		if err != nil {
			return err
		}
		obj.OrderType = OrderType(tmp)
	}
	// Deserialize `ClientId`:
	err = decoder.Decode(&obj.ClientId)
	if err != nil {
		return err
	}
	// Deserialize `SelfTradeBehavior`:
	{
		tmp, err := decoder.ReadUint32(bin.LE)
		if err != nil {
			return err
		}
		obj.SelfTradeBehavior = SelfTradeBehavior(tmp)
	}
	return nil
}

type NewOrderInstructionV3 struct {
	Side                        Side
	LimitPrice                  uint64
	MaxCoinQty                  uint64
	MaxNativePcQtyIncludingFees uint64
	SelfTradeBehavior           SelfTradeBehavior
	OrderType                   OrderType
	ClientOrderId               uint64
	Limit                       uint16
}

func (obj NewOrderInstructionV3) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Side` param:
	err = encoder.WriteUint32(uint32(obj.Side), bin.LE)
	if err != nil {
		return err
	}
	// Serialize `LimitPrice` param:
	err = encoder.Encode(obj.LimitPrice)
	if err != nil {
		return err
	}
	// Serialize `MaxCoinQty` param:
	err = encoder.Encode(obj.MaxCoinQty)
	if err != nil {
		return err
	}
	// Serialize `MaxNativePcQtyIncludingFees` param:
	err = encoder.Encode(obj.MaxNativePcQtyIncludingFees)
	if err != nil {
		return err
	}
	// Serialize `SelfTradeBehavior` param:
	err = encoder.WriteUint32(uint32(obj.SelfTradeBehavior), bin.LE)
	if err != nil {
		return err
	}
	// Serialize `OrderType` param:
	err = encoder.WriteUint32(uint32(obj.OrderType), bin.LE)
	if err != nil {
		return err
	}
	// Serialize `ClientOrderId` param:
	err = encoder.Encode(obj.ClientOrderId)
	if err != nil {
		return err
	}
	// Serialize `Limit` param:
	err = encoder.Encode(obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

func (obj *NewOrderInstructionV3) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Side`:
	{
		tmp, err := decoder.ReadUint32(bin.LE)
		if err != nil {
			return err
		}
		obj.Side = Side(tmp)
	}
	// Deserialize `LimitPrice`:
	err = decoder.Decode(&obj.LimitPrice)
	if err != nil {
		return err
	}
	// Deserialize `MaxCoinQty`:
	err = decoder.Decode(&obj.MaxCoinQty)
	if err != nil {
		return err
	}
	// Deserialize `MaxNativePcQtyIncludingFees`:
	err = decoder.Decode(&obj.MaxNativePcQtyIncludingFees)
	if err != nil {
		return err
	}
	// Deserialize `SelfTradeBehavior`:
	{
		tmp, err := decoder.ReadUint32(bin.LE)
		if err != nil {
			return err
		}
		obj.SelfTradeBehavior = SelfTradeBehavior(tmp)
	}
	// Deserialize `OrderType`:
	{
		tmp, err := decoder.ReadUint32(bin.LE)
		if err != nil {
			return err
		}
		obj.OrderType = OrderType(tmp)
	}
	// Deserialize `ClientOrderId`:
	err = decoder.Decode(&obj.ClientOrderId)
	if err != nil {
		return err
	}
	// Deserialize `Limit`:
	err = decoder.Decode(&obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

type SendTakeInstruction struct {
	Side                        Side
	LimitPrice                  uint64
	MaxCoinQty                  uint64
	MaxNativePcQtyIncludingFees uint64
	MinCoinQty                  uint64
	MinNativePcQty              uint64
	Limit                       uint16
}

func (obj SendTakeInstruction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Side` param:
	err = encoder.WriteUint32(uint32(obj.Side), bin.LE)
	if err != nil {
		return err
	}
	// Serialize `LimitPrice` param:
	err = encoder.Encode(obj.LimitPrice)
	if err != nil {
		return err
	}
	// Serialize `MaxCoinQty` param:
	err = encoder.Encode(obj.MaxCoinQty)
	if err != nil {
		return err
	}
	// Serialize `MaxNativePcQtyIncludingFees` param:
	err = encoder.Encode(obj.MaxNativePcQtyIncludingFees)
	if err != nil {
		return err
	}
	// Serialize `MinCoinQty` param:
	err = encoder.Encode(obj.MinCoinQty)
	if err != nil {
		return err
	}
	// Serialize `MinNativePcQty` param:
	err = encoder.Encode(obj.MinNativePcQty)
	if err != nil {
		return err
	}
	// Serialize `Limit` param:
	err = encoder.Encode(obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SendTakeInstruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Side`:
	{
		tmp, err := decoder.ReadUint32(bin.LE)
		if err != nil {
			return err
		}
		obj.Side = Side(tmp)
	}
	// Deserialize `LimitPrice`:
	err = decoder.Decode(&obj.LimitPrice)
	if err != nil {
		return err
	}
	// Deserialize `MaxCoinQty`:
	err = decoder.Decode(&obj.MaxCoinQty)
	if err != nil {
		return err
	}
	// Deserialize `MaxNativePcQtyIncludingFees`:
	err = decoder.Decode(&obj.MaxNativePcQtyIncludingFees)
	if err != nil {
		return err
	}
	// Deserialize `MinCoinQty`:
	err = decoder.Decode(&obj.MinCoinQty)
	if err != nil {
		return err
	}
	// Deserialize `MinNativePcQty`:
	err = decoder.Decode(&obj.MinNativePcQty)
	if err != nil {
		return err
	}
	// Deserialize `Limit`:
	err = decoder.Decode(&obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

type Side ag_binary.BorshEnum

const (
	SideBid Side = iota
	SideAsk
)

func (value Side) String() string {
	switch value {
	case SideBid:
		return "Bid"
	case SideAsk:
		return "Ask"
	default:
		return ""
	}
}

type OrderType ag_binary.BorshEnum

const (
	OrderTypeLimit OrderType = iota
	OrderTypeImmediateOrCancel
	OrderTypePostOnly
)

func (value OrderType) String() string {
	switch value {
	case OrderTypeLimit:
		return "Limit"
	case OrderTypeImmediateOrCancel:
		return "ImmediateOrCancel"
	case OrderTypePostOnly:
		return "PostOnly"
	default:
		return ""
	}
}

type SelfTradeBehavior ag_binary.BorshEnum

const (
	SelfTradeBehaviorDecrementTake SelfTradeBehavior = iota
	SelfTradeBehaviorCancelProvide
	SelfTradeBehaviorAbortTransaction
)

func (value SelfTradeBehavior) String() string {
	switch value {
	case SelfTradeBehaviorDecrementTake:
		return "DecrementTake"
	case SelfTradeBehaviorCancelProvide:
		return "CancelProvide"
	case SelfTradeBehaviorAbortTransaction:
		return "AbortTransaction"
	default:
		return ""
	}
}

type FeeTier ag_binary.BorshEnum

const (
	FeeTierBase FeeTier = iota
	FeeTierSRM2
	FeeTierSRM3
	FeeTierSRM4
	FeeTierSRM5
	FeeTierSRM6
	FeeTierMSRM
	FeeTierStable
)

func (value FeeTier) String() string {
	switch value {
	case FeeTierBase:
		return "Base"
	case FeeTierSRM2:
		return "SRM2"
	case FeeTierSRM3:
		return "SRM3"
	case FeeTierSRM4:
		return "SRM4"
	case FeeTierSRM5:
		return "SRM5"
	case FeeTierSRM6:
		return "SRM6"
	case FeeTierMSRM:
		return "MSRM"
	case FeeTierStable:
		return "Stable"
	default:
		return ""
	}
}
