// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package solmate_cba

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/SolmateDev/solana-go"
	ag_format "github.com/SolmateDev/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InsertBid is the `insertBid` instruction.
type InsertBid struct {
	Bid *Bid

	// [0] = [] controller
	//
	// [1] = [WRITE] pcVault
	//
	// [2] = [] validatorPipeline
	//
	// [3] = [WRITE] periods
	//
	// [4] = [WRITE] bids
	//
	// [5] = [WRITE] userFund
	//
	// [6] = [WRITE, SIGNER] user
	//
	// [7] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInsertBidInstructionBuilder creates a new `InsertBid` instruction builder.
func NewInsertBidInstructionBuilder() *InsertBid {
	nd := &InsertBid{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetBid sets the "bid" parameter.
func (inst *InsertBid) SetBid(bid Bid) *InsertBid {
	inst.Bid = &bid
	return inst
}

// SetControllerAccount sets the "controller" account.
func (inst *InsertBid) SetControllerAccount(controller ag_solanago.PublicKey) *InsertBid {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(controller)
	return inst
}

// GetControllerAccount gets the "controller" account.
func (inst *InsertBid) GetControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPcVaultAccount sets the "pcVault" account.
func (inst *InsertBid) SetPcVaultAccount(pcVault ag_solanago.PublicKey) *InsertBid {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(pcVault).WRITE()
	return inst
}

// GetPcVaultAccount gets the "pcVault" account.
func (inst *InsertBid) GetPcVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetValidatorPipelineAccount sets the "validatorPipeline" account.
func (inst *InsertBid) SetValidatorPipelineAccount(validatorPipeline ag_solanago.PublicKey) *InsertBid {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(validatorPipeline)
	return inst
}

// GetValidatorPipelineAccount gets the "validatorPipeline" account.
func (inst *InsertBid) GetValidatorPipelineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPeriodsAccount sets the "periods" account.
func (inst *InsertBid) SetPeriodsAccount(periods ag_solanago.PublicKey) *InsertBid {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(periods).WRITE()
	return inst
}

// GetPeriodsAccount gets the "periods" account.
func (inst *InsertBid) GetPeriodsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetBidsAccount sets the "bids" account.
func (inst *InsertBid) SetBidsAccount(bids ag_solanago.PublicKey) *InsertBid {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
func (inst *InsertBid) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetUserFundAccount sets the "userFund" account.
func (inst *InsertBid) SetUserFundAccount(userFund ag_solanago.PublicKey) *InsertBid {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(userFund).WRITE()
	return inst
}

// GetUserFundAccount gets the "userFund" account.
func (inst *InsertBid) GetUserFundAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetUserAccount sets the "user" account.
func (inst *InsertBid) SetUserAccount(user ag_solanago.PublicKey) *InsertBid {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *InsertBid) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *InsertBid) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *InsertBid {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *InsertBid) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst InsertBid) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InsertBid,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InsertBid) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InsertBid) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Bid == nil {
			return errors.New("Bid parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Controller is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.PcVault is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ValidatorPipeline is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Periods is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.UserFund is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *InsertBid) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InsertBid")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Bid", *inst.Bid))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       controller", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          pcVault", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("validatorPipeline", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          periods", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             bids", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         userFund", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("             user", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("     tokenProgram", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj InsertBid) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bid` param:
	err = encoder.Encode(obj.Bid)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InsertBid) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bid`:
	err = decoder.Decode(&obj.Bid)
	if err != nil {
		return err
	}
	return nil
}

// NewInsertBidInstruction declares a new InsertBid instruction with the provided parameters and accounts.
func NewInsertBidInstruction(
	// Parameters:
	bid Bid,
	// Accounts:
	controller ag_solanago.PublicKey,
	pcVault ag_solanago.PublicKey,
	validatorPipeline ag_solanago.PublicKey,
	periods ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	userFund ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *InsertBid {
	return NewInsertBidInstructionBuilder().
		SetBid(bid).
		SetControllerAccount(controller).
		SetPcVaultAccount(pcVault).
		SetValidatorPipelineAccount(validatorPipeline).
		SetPeriodsAccount(periods).
		SetBidsAccount(bids).
		SetUserFundAccount(userFund).
		SetUserAccount(user).
		SetTokenProgramAccount(tokenProgram)
}
