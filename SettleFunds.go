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

// SettleFunds is the `SettleFunds` instruction.
type SettleFunds struct {

	// [0] = [WRITE] market
	// ··········· market
	//
	// [1] = [WRITE] openOrders
	// ··········· OpenOrders
	//
	// [2] = [SIGNER] owner
	// ··········· the OpenOrders owner
	//
	// [3] = [WRITE] coinVault
	// ··········· coin vault
	//
	// [4] = [WRITE] pcVault
	// ··········· pc vault
	//
	// [5] = [WRITE] coinWallet
	// ··········· coin wallet
	//
	// [6] = [WRITE] pcWallet
	// ··········· pc wallet
	//
	// [7] = [] vaultSigner
	// ··········· vault signer
	//
	// [8] = [] splTokenProgram
	// ··········· spl token program
	//
	// [9] = [WRITE] referrerPcWallet
	// ··········· (optional) referrer pc wallet
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSettleFundsInstructionBuilder creates a new `SettleFunds` instruction builder.
func NewSettleFundsInstructionBuilder() *SettleFunds {
	nd := &SettleFunds{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetMarketAccount sets the "market" account.
// market
func (inst *SettleFunds) SetMarketAccount(market ag_solanago.PublicKey) *SettleFunds {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
// market
func (inst *SettleFunds) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetOpenOrdersAccount sets the "openOrders" account.
// OpenOrders
func (inst *SettleFunds) SetOpenOrdersAccount(openOrders ag_solanago.PublicKey) *SettleFunds {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(openOrders).WRITE()
	return inst
}

// GetOpenOrdersAccount gets the "openOrders" account.
// OpenOrders
func (inst *SettleFunds) GetOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetOwnerAccount sets the "owner" account.
// the OpenOrders owner
func (inst *SettleFunds) SetOwnerAccount(owner ag_solanago.PublicKey) *SettleFunds {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
// the OpenOrders owner
func (inst *SettleFunds) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetCoinVaultAccount sets the "coinVault" account.
// coin vault
func (inst *SettleFunds) SetCoinVaultAccount(coinVault ag_solanago.PublicKey) *SettleFunds {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(coinVault).WRITE()
	return inst
}

// GetCoinVaultAccount gets the "coinVault" account.
// coin vault
func (inst *SettleFunds) GetCoinVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetPcVaultAccount sets the "pcVault" account.
// pc vault
func (inst *SettleFunds) SetPcVaultAccount(pcVault ag_solanago.PublicKey) *SettleFunds {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(pcVault).WRITE()
	return inst
}

// GetPcVaultAccount gets the "pcVault" account.
// pc vault
func (inst *SettleFunds) GetPcVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetCoinWalletAccount sets the "coinWallet" account.
// coin wallet
func (inst *SettleFunds) SetCoinWalletAccount(coinWallet ag_solanago.PublicKey) *SettleFunds {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(coinWallet).WRITE()
	return inst
}

// GetCoinWalletAccount gets the "coinWallet" account.
// coin wallet
func (inst *SettleFunds) GetCoinWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetPcWalletAccount sets the "pcWallet" account.
// pc wallet
func (inst *SettleFunds) SetPcWalletAccount(pcWallet ag_solanago.PublicKey) *SettleFunds {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(pcWallet).WRITE()
	return inst
}

// GetPcWalletAccount gets the "pcWallet" account.
// pc wallet
func (inst *SettleFunds) GetPcWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetVaultSignerAccount sets the "vaultSigner" account.
// vault signer
func (inst *SettleFunds) SetVaultSignerAccount(vaultSigner ag_solanago.PublicKey) *SettleFunds {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(vaultSigner)
	return inst
}

// GetVaultSignerAccount gets the "vaultSigner" account.
// vault signer
func (inst *SettleFunds) GetVaultSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetSplTokenProgramAccount sets the "splTokenProgram" account.
// spl token program
func (inst *SettleFunds) SetSplTokenProgramAccount(splTokenProgram ag_solanago.PublicKey) *SettleFunds {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(splTokenProgram)
	return inst
}

// GetSplTokenProgramAccount gets the "splTokenProgram" account.
// spl token program
func (inst *SettleFunds) GetSplTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetReferrerPcWalletAccount sets the "referrerPcWallet" account.
// (optional) referrer pc wallet
func (inst *SettleFunds) SetReferrerPcWalletAccount(referrerPcWallet ag_solanago.PublicKey) *SettleFunds {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(referrerPcWallet).WRITE()
	return inst
}

// GetReferrerPcWalletAccount gets the "referrerPcWallet" account.
// (optional) referrer pc wallet
func (inst *SettleFunds) GetReferrerPcWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

func (inst SettleFunds) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_SettleFunds, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SettleFunds) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SettleFunds) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OpenOrders is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.CoinVault is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.PcVault is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.CoinWallet is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PcWallet is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.VaultSigner is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SplTokenProgram is not set")
		}

		// [9] = ReferrerPcWallet is optional

	}
	return nil
}

func (inst *SettleFunds) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SettleFunds")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          market", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("      openOrders", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("           owner", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("       coinVault", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("         pcVault", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("      coinWallet", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("        pcWallet", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("     vaultSigner", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta(" splTokenProgram", inst.AccountMetaSlice[8]))
						if len(inst.AccountMetaSlice) > 9 {
							accountsBranch.Child(ag_format.Meta("referrerPcWallet", inst.AccountMetaSlice[9]))
						}
					})
				})
		})
}

func (obj SettleFunds) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SettleFunds) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSettleFundsInstruction declares a new SettleFunds instruction with the provided parameters and accounts.
func NewSettleFundsInstruction(
	// Accounts:
	market ag_solanago.PublicKey,
	openOrders ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	coinVault ag_solanago.PublicKey,
	pcVault ag_solanago.PublicKey,
	coinWallet ag_solanago.PublicKey,
	pcWallet ag_solanago.PublicKey,
	vaultSigner ag_solanago.PublicKey,
	splTokenProgram ag_solanago.PublicKey,
	referrerPcWallet ag_solanago.PublicKey) *SettleFunds {
	return NewSettleFundsInstructionBuilder().
		SetMarketAccount(market).
		SetOpenOrdersAccount(openOrders).
		SetOwnerAccount(owner).
		SetCoinVaultAccount(coinVault).
		SetPcVaultAccount(pcVault).
		SetCoinWalletAccount(coinWallet).
		SetPcWalletAccount(pcWallet).
		SetVaultSignerAccount(vaultSigner).
		SetSplTokenProgramAccount(splTokenProgram).
		SetReferrerPcWalletAccount(referrerPcWallet)
}
