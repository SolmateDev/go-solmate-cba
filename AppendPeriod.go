// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package solmate_cba

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/SolmateDev/solana-go"
	ag_format "github.com/SolmateDev/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AppendPeriod is the `appendPeriod` instruction.
type AppendPeriod struct {
	Withhold     *uint64
	Start        *uint64
	Length       *uint64
	DecayRateNum *uint64
	DecayRateDen *uint64

	// [0] = [] controller
	//
	// [1] = [WRITE] validatorPipeline
	//
	// [2] = [WRITE] periods
	//
	// [3] = [WRITE, SIGNER] admin
	//
	// [4] = [] tokenProgram
	//
	// [5] = [] clock
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAppendPeriodInstructionBuilder creates a new `AppendPeriod` instruction builder.
func NewAppendPeriodInstructionBuilder() *AppendPeriod {
	nd := &AppendPeriod{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetWithhold sets the "withhold" parameter.
func (inst *AppendPeriod) SetWithhold(withhold uint64) *AppendPeriod {
	inst.Withhold = &withhold
	return inst
}

// SetStart sets the "start" parameter.
func (inst *AppendPeriod) SetStart(start uint64) *AppendPeriod {
	inst.Start = &start
	return inst
}

// SetLength sets the "length" parameter.
func (inst *AppendPeriod) SetLength(length uint64) *AppendPeriod {
	inst.Length = &length
	return inst
}

// SetDecayRateNum sets the "decayRateNum" parameter.
func (inst *AppendPeriod) SetDecayRateNum(decayRateNum uint64) *AppendPeriod {
	inst.DecayRateNum = &decayRateNum
	return inst
}

// SetDecayRateDen sets the "decayRateDen" parameter.
func (inst *AppendPeriod) SetDecayRateDen(decayRateDen uint64) *AppendPeriod {
	inst.DecayRateDen = &decayRateDen
	return inst
}

// SetControllerAccount sets the "controller" account.
func (inst *AppendPeriod) SetControllerAccount(controller ag_solanago.PublicKey) *AppendPeriod {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(controller)
	return inst
}

// GetControllerAccount gets the "controller" account.
func (inst *AppendPeriod) GetControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetValidatorPipelineAccount sets the "validatorPipeline" account.
func (inst *AppendPeriod) SetValidatorPipelineAccount(validatorPipeline ag_solanago.PublicKey) *AppendPeriod {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(validatorPipeline).WRITE()
	return inst
}

// GetValidatorPipelineAccount gets the "validatorPipeline" account.
func (inst *AppendPeriod) GetValidatorPipelineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPeriodsAccount sets the "periods" account.
func (inst *AppendPeriod) SetPeriodsAccount(periods ag_solanago.PublicKey) *AppendPeriod {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(periods).WRITE()
	return inst
}

// GetPeriodsAccount gets the "periods" account.
func (inst *AppendPeriod) GetPeriodsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAdminAccount sets the "admin" account.
func (inst *AppendPeriod) SetAdminAccount(admin ag_solanago.PublicKey) *AppendPeriod {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *AppendPeriod) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *AppendPeriod) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *AppendPeriod {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *AppendPeriod) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetClockAccount sets the "clock" account.
func (inst *AppendPeriod) SetClockAccount(clock ag_solanago.PublicKey) *AppendPeriod {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(clock)
	return inst
}

// GetClockAccount gets the "clock" account.
func (inst *AppendPeriod) GetClockAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst AppendPeriod) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AppendPeriod,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AppendPeriod) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AppendPeriod) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Withhold == nil {
			return errors.New("Withhold parameter is not set")
		}
		if inst.Start == nil {
			return errors.New("Start parameter is not set")
		}
		if inst.Length == nil {
			return errors.New("Length parameter is not set")
		}
		if inst.DecayRateNum == nil {
			return errors.New("DecayRateNum parameter is not set")
		}
		if inst.DecayRateDen == nil {
			return errors.New("DecayRateDen parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Controller is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ValidatorPipeline is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Periods is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Clock is not set")
		}
	}
	return nil
}

func (inst *AppendPeriod) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AppendPeriod")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    Withhold", *inst.Withhold))
						paramsBranch.Child(ag_format.Param("       Start", *inst.Start))
						paramsBranch.Child(ag_format.Param("      Length", *inst.Length))
						paramsBranch.Child(ag_format.Param("DecayRateNum", *inst.DecayRateNum))
						paramsBranch.Child(ag_format.Param("DecayRateDen", *inst.DecayRateDen))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       controller", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("validatorPipeline", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          periods", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            admin", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     tokenProgram", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("            clock", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj AppendPeriod) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Withhold` param:
	err = encoder.Encode(obj.Withhold)
	if err != nil {
		return err
	}
	// Serialize `Start` param:
	err = encoder.Encode(obj.Start)
	if err != nil {
		return err
	}
	// Serialize `Length` param:
	err = encoder.Encode(obj.Length)
	if err != nil {
		return err
	}
	// Serialize `DecayRateNum` param:
	err = encoder.Encode(obj.DecayRateNum)
	if err != nil {
		return err
	}
	// Serialize `DecayRateDen` param:
	err = encoder.Encode(obj.DecayRateDen)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AppendPeriod) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Withhold`:
	err = decoder.Decode(&obj.Withhold)
	if err != nil {
		return err
	}
	// Deserialize `Start`:
	err = decoder.Decode(&obj.Start)
	if err != nil {
		return err
	}
	// Deserialize `Length`:
	err = decoder.Decode(&obj.Length)
	if err != nil {
		return err
	}
	// Deserialize `DecayRateNum`:
	err = decoder.Decode(&obj.DecayRateNum)
	if err != nil {
		return err
	}
	// Deserialize `DecayRateDen`:
	err = decoder.Decode(&obj.DecayRateDen)
	if err != nil {
		return err
	}
	return nil
}

// NewAppendPeriodInstruction declares a new AppendPeriod instruction with the provided parameters and accounts.
func NewAppendPeriodInstruction(
	// Parameters:
	withhold uint64,
	start uint64,
	length uint64,
	decayRateNum uint64,
	decayRateDen uint64,
	// Accounts:
	controller ag_solanago.PublicKey,
	validatorPipeline ag_solanago.PublicKey,
	periods ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	clock ag_solanago.PublicKey) *AppendPeriod {
	return NewAppendPeriodInstructionBuilder().
		SetWithhold(withhold).
		SetStart(start).
		SetLength(length).
		SetDecayRateNum(decayRateNum).
		SetDecayRateDen(decayRateDen).
		SetControllerAccount(controller).
		SetValidatorPipelineAccount(validatorPipeline).
		SetPeriodsAccount(periods).
		SetAdminAccount(admin).
		SetTokenProgramAccount(tokenProgram).
		SetClockAccount(clock)
}
