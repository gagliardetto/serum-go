// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package serumgo

import (
	"encoding/binary"
	"errors"

	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CancelOrderByClientIdV2 is the `CancelOrderByClientIdV2` instruction.
type CancelOrderByClientIdV2 struct {
	ClientOrderId *uint64

	// [0] = [WRITE] market
	// ··········· market
	//
	// [1] = [WRITE] bids
	// ··········· bids
	//
	// [2] = [WRITE] asks
	// ··········· asks
	//
	// [3] = [WRITE] openOrders
	// ··········· OpenOrders
	//
	// [4] = [SIGNER] owner
	// ··········· the OpenOrders owner
	//
	// [5] = [WRITE] eventQueue
	// ··········· event_q
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCancelOrderByClientIdV2InstructionBuilder creates a new `CancelOrderByClientIdV2` instruction builder.
func NewCancelOrderByClientIdV2InstructionBuilder() *CancelOrderByClientIdV2 {
	nd := &CancelOrderByClientIdV2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetClientOrderId sets the "client_order_id" parameter.
func (inst *CancelOrderByClientIdV2) SetClientOrderId(client_order_id uint64) *CancelOrderByClientIdV2 {
	inst.ClientOrderId = &client_order_id
	return inst
}

// SetMarketAccount sets the "market" account.
// market
func (inst *CancelOrderByClientIdV2) SetMarketAccount(market ag_solanago.PublicKey) *CancelOrderByClientIdV2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
// market
func (inst *CancelOrderByClientIdV2) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetBidsAccount sets the "bids" account.
// bids
func (inst *CancelOrderByClientIdV2) SetBidsAccount(bids ag_solanago.PublicKey) *CancelOrderByClientIdV2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
// bids
func (inst *CancelOrderByClientIdV2) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAsksAccount sets the "asks" account.
// asks
func (inst *CancelOrderByClientIdV2) SetAsksAccount(asks ag_solanago.PublicKey) *CancelOrderByClientIdV2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(asks).WRITE()
	return inst
}

// GetAsksAccount gets the "asks" account.
// asks
func (inst *CancelOrderByClientIdV2) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOpenOrdersAccount sets the "openOrders" account.
// OpenOrders
func (inst *CancelOrderByClientIdV2) SetOpenOrdersAccount(openOrders ag_solanago.PublicKey) *CancelOrderByClientIdV2 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(openOrders).WRITE()
	return inst
}

// GetOpenOrdersAccount gets the "openOrders" account.
// OpenOrders
func (inst *CancelOrderByClientIdV2) GetOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetOwnerAccount sets the "owner" account.
// the OpenOrders owner
func (inst *CancelOrderByClientIdV2) SetOwnerAccount(owner ag_solanago.PublicKey) *CancelOrderByClientIdV2 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
// the OpenOrders owner
func (inst *CancelOrderByClientIdV2) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetEventQueueAccount sets the "eventQueue" account.
// event_q
func (inst *CancelOrderByClientIdV2) SetEventQueueAccount(eventQueue ag_solanago.PublicKey) *CancelOrderByClientIdV2 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(eventQueue).WRITE()
	return inst
}

// GetEventQueueAccount gets the "eventQueue" account.
// event_q
func (inst *CancelOrderByClientIdV2) GetEventQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst CancelOrderByClientIdV2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_CancelOrderByClientIdV2, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CancelOrderByClientIdV2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CancelOrderByClientIdV2) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ClientOrderId == nil {
			return errors.New("ClientOrderId parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Asks is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.OpenOrders is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.EventQueue is not set")
		}
	}
	return nil
}

func (inst *CancelOrderByClientIdV2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CancelOrderByClientIdV2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  ClientOrderId", *inst.ClientOrderId))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    market", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      bids", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      asks", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("openOrders", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     owner", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("eventQueue", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj CancelOrderByClientIdV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ClientOrderId` param:
	err = encoder.Encode(obj.ClientOrderId)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CancelOrderByClientIdV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ClientOrderId`:
	err = decoder.Decode(&obj.ClientOrderId)
	if err != nil {
		return err
	}
	return nil
}

// NewCancelOrderByClientIdV2Instruction declares a new CancelOrderByClientIdV2 instruction with the provided parameters and accounts.
func NewCancelOrderByClientIdV2Instruction(
	// Parameters:
	client_order_id uint64,
	// Accounts:
	market ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	asks ag_solanago.PublicKey,
	openOrders ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	eventQueue ag_solanago.PublicKey) *CancelOrderByClientIdV2 {
	return NewCancelOrderByClientIdV2InstructionBuilder().
		SetClientOrderId(client_order_id).
		SetMarketAccount(market).
		SetBidsAccount(bids).
		SetAsksAccount(asks).
		SetOpenOrdersAccount(openOrders).
		SetOwnerAccount(owner).
		SetEventQueueAccount(eventQueue)
}
