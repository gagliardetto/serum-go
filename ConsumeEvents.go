// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package serum_dex

import (
	"encoding/binary"
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ConsumeEvents is the `ConsumeEvents` instruction.
type ConsumeEvents struct {
	Limit *uint16

	// [0] = [WRITE] openOrders
	// ··········· OpenOrders; TODO: this is an array of accounts
	//
	// [1] = [WRITE] market
	// ··········· market
	//
	// [2] = [WRITE] eventQueue
	// ··········· event queue
	//
	// [3] = [WRITE] coinFeeReceivable
	//
	// [4] = [WRITE] pcFeeReceivable
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewConsumeEventsInstructionBuilder creates a new `ConsumeEvents` instruction builder.
func NewConsumeEventsInstructionBuilder() *ConsumeEvents {
	nd := &ConsumeEvents{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetLimit sets the "limit" parameter.
func (inst *ConsumeEvents) SetLimit(limit uint16) *ConsumeEvents {
	inst.Limit = &limit
	return inst
}

// SetOpenOrdersAccount sets the "openOrders" account.
// OpenOrders; TODO: this is an array of accounts
func (inst *ConsumeEvents) SetOpenOrdersAccount(openOrders ag_solanago.PublicKey) *ConsumeEvents {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(openOrders).WRITE()
	return inst
}

// GetOpenOrdersAccount gets the "openOrders" account.
// OpenOrders; TODO: this is an array of accounts
func (inst *ConsumeEvents) GetOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetMarketAccount sets the "market" account.
// market
func (inst *ConsumeEvents) SetMarketAccount(market ag_solanago.PublicKey) *ConsumeEvents {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
// market
func (inst *ConsumeEvents) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetEventQueueAccount sets the "eventQueue" account.
// event queue
func (inst *ConsumeEvents) SetEventQueueAccount(eventQueue ag_solanago.PublicKey) *ConsumeEvents {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(eventQueue).WRITE()
	return inst
}

// GetEventQueueAccount gets the "eventQueue" account.
// event queue
func (inst *ConsumeEvents) GetEventQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetCoinFeeReceivableAccount sets the "coinFeeReceivable" account.
func (inst *ConsumeEvents) SetCoinFeeReceivableAccount(coinFeeReceivable ag_solanago.PublicKey) *ConsumeEvents {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(coinFeeReceivable).WRITE()
	return inst
}

// GetCoinFeeReceivableAccount gets the "coinFeeReceivable" account.
func (inst *ConsumeEvents) GetCoinFeeReceivableAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetPcFeeReceivableAccount sets the "pcFeeReceivable" account.
func (inst *ConsumeEvents) SetPcFeeReceivableAccount(pcFeeReceivable ag_solanago.PublicKey) *ConsumeEvents {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(pcFeeReceivable).WRITE()
	return inst
}

// GetPcFeeReceivableAccount gets the "pcFeeReceivable" account.
func (inst *ConsumeEvents) GetPcFeeReceivableAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

func (inst ConsumeEvents) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_ConsumeEvents, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ConsumeEvents) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ConsumeEvents) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Limit == nil {
			return errors.New("Limit parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.OpenOrders is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.EventQueue is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.CoinFeeReceivable is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.PcFeeReceivable is not set")
		}
	}
	return nil
}

func (inst *ConsumeEvents) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ConsumeEvents")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Limit", *inst.Limit))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       openOrders", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("           market", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("       eventQueue", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("coinFeeReceivable", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("  pcFeeReceivable", inst.AccountMetaSlice[4]))
					})
				})
		})
}

func (obj ConsumeEvents) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Limit` param:
	err = encoder.Encode(obj.Limit)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ConsumeEvents) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Limit`:
	err = decoder.Decode(&obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

// NewConsumeEventsInstruction declares a new ConsumeEvents instruction with the provided parameters and accounts.
func NewConsumeEventsInstruction(
	// Parameters:
	limit uint16,
	// Accounts:
	openOrders ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	eventQueue ag_solanago.PublicKey,
	coinFeeReceivable ag_solanago.PublicKey,
	pcFeeReceivable ag_solanago.PublicKey) *ConsumeEvents {
	return NewConsumeEventsInstructionBuilder().
		SetLimit(limit).
		SetOpenOrdersAccount(openOrders).
		SetMarketAccount(market).
		SetEventQueueAccount(eventQueue).
		SetCoinFeeReceivableAccount(coinFeeReceivable).
		SetPcFeeReceivableAccount(pcFeeReceivable)
}
