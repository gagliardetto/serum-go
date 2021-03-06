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

// Removes all orders for a given open orders account from the orderbook.
type Prune struct {
	Limit *uint16

	// [0] = [WRITE] market
	// ··········· market
	//
	// [1] = [WRITE] bids
	// ··········· bids
	//
	// [2] = [WRITE] asks
	// ··········· asks
	//
	// [3] = [SIGNER] pruneAuthority
	// ··········· prune authority
	//
	// [4] = [WRITE] openOrders
	// ··········· open orders.
	//
	// [5] = [] owner
	// ··········· open orders owner.
	//
	// [6] = [WRITE] eventQueue
	// ··········· event queue.
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewPruneInstructionBuilder creates a new `Prune` instruction builder.
func NewPruneInstructionBuilder() *Prune {
	nd := &Prune{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetLimit sets the "limit" parameter.
func (inst *Prune) SetLimit(limit uint16) *Prune {
	inst.Limit = &limit
	return inst
}

// SetMarketAccount sets the "market" account.
// market
func (inst *Prune) SetMarketAccount(market ag_solanago.PublicKey) *Prune {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
// market
func (inst *Prune) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetBidsAccount sets the "bids" account.
// bids
func (inst *Prune) SetBidsAccount(bids ag_solanago.PublicKey) *Prune {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
// bids
func (inst *Prune) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAsksAccount sets the "asks" account.
// asks
func (inst *Prune) SetAsksAccount(asks ag_solanago.PublicKey) *Prune {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(asks).WRITE()
	return inst
}

// GetAsksAccount gets the "asks" account.
// asks
func (inst *Prune) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPruneAuthorityAccount sets the "pruneAuthority" account.
// prune authority
func (inst *Prune) SetPruneAuthorityAccount(pruneAuthority ag_solanago.PublicKey) *Prune {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(pruneAuthority).SIGNER()
	return inst
}

// GetPruneAuthorityAccount gets the "pruneAuthority" account.
// prune authority
func (inst *Prune) GetPruneAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetOpenOrdersAccount sets the "openOrders" account.
// open orders.
func (inst *Prune) SetOpenOrdersAccount(openOrders ag_solanago.PublicKey) *Prune {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(openOrders).WRITE()
	return inst
}

// GetOpenOrdersAccount gets the "openOrders" account.
// open orders.
func (inst *Prune) GetOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetOwnerAccount sets the "owner" account.
// open orders owner.
func (inst *Prune) SetOwnerAccount(owner ag_solanago.PublicKey) *Prune {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(owner)
	return inst
}

// GetOwnerAccount gets the "owner" account.
// open orders owner.
func (inst *Prune) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetEventQueueAccount sets the "eventQueue" account.
// event queue.
func (inst *Prune) SetEventQueueAccount(eventQueue ag_solanago.PublicKey) *Prune {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(eventQueue).WRITE()
	return inst
}

// GetEventQueueAccount gets the "eventQueue" account.
// event queue.
func (inst *Prune) GetEventQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst Prune) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_Prune, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Prune) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Prune) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Limit == nil {
			return errors.New("Limit parameter is not set")
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
			return errors.New("accounts.PruneAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.OpenOrders is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.EventQueue is not set")
		}
	}
	return nil
}

func (inst *Prune) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Prune")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Limit", *inst.Limit))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        market", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          bids", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          asks", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("pruneAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    openOrders", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         owner", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("    eventQueue", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj Prune) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Limit` param:
	err = encoder.Encode(obj.Limit)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Prune) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Limit`:
	err = decoder.Decode(&obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

// NewPruneInstruction declares a new Prune instruction with the provided parameters and accounts.
func NewPruneInstruction(
	// Parameters:
	limit uint16,
	// Accounts:
	market ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	asks ag_solanago.PublicKey,
	pruneAuthority ag_solanago.PublicKey,
	openOrders ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	eventQueue ag_solanago.PublicKey) *Prune {
	return NewPruneInstructionBuilder().
		SetLimit(limit).
		SetMarketAccount(market).
		SetBidsAccount(bids).
		SetAsksAccount(asks).
		SetPruneAuthorityAccount(pruneAuthority).
		SetOpenOrdersAccount(openOrders).
		SetOwnerAccount(owner).
		SetEventQueueAccount(eventQueue)
}
