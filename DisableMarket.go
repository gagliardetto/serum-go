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

// DisableMarket is the `DisableMarket` instruction.
type DisableMarket struct {

	// [0] = [WRITE] market
	// ··········· market
	//
	// [1] = [SIGNER] disableAuthority
	// ··········· disable authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDisableMarketInstructionBuilder creates a new `DisableMarket` instruction builder.
func NewDisableMarketInstructionBuilder() *DisableMarket {
	nd := &DisableMarket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetMarketAccount sets the "market" account.
// market
func (inst *DisableMarket) SetMarketAccount(market ag_solanago.PublicKey) *DisableMarket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
// market
func (inst *DisableMarket) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetDisableAuthorityAccount sets the "disableAuthority" account.
// disable authority
func (inst *DisableMarket) SetDisableAuthorityAccount(disableAuthority ag_solanago.PublicKey) *DisableMarket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(disableAuthority).SIGNER()
	return inst
}

// GetDisableAuthorityAccount gets the "disableAuthority" account.
// disable authority
func (inst *DisableMarket) GetDisableAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst DisableMarket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_DisableMarket, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DisableMarket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DisableMarket) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.DisableAuthority is not set")
		}
	}
	return nil
}

func (inst *DisableMarket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DisableMarket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          market", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("disableAuthority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj DisableMarket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *DisableMarket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewDisableMarketInstruction declares a new DisableMarket instruction with the provided parameters and accounts.
func NewDisableMarketInstruction(
	// Accounts:
	market ag_solanago.PublicKey,
	disableAuthority ag_solanago.PublicKey) *DisableMarket {
	return NewDisableMarketInstructionBuilder().
		SetMarketAccount(market).
		SetDisableAuthorityAccount(disableAuthority)
}
