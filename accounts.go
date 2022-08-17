// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package solmate_cba

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/SolmateDev/solana-go"
)

type BidList struct {
	Pipeline             ag_solanago.PublicKey
	Book                 []Bid
	LastPeriodStart      uint64
	CrankFeeRate         [2]uint64
	BandwidthDenominator uint64
	TotalDeposits        uint64
}

var BidListDiscriminator = [8]byte{233, 127, 13, 29, 123, 209, 192, 79}

func (obj BidList) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(BidListDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pipeline` param:
	err = encoder.Encode(obj.Pipeline)
	if err != nil {
		return err
	}
	// Serialize `Book` param:
	err = encoder.Encode(obj.Book)
	if err != nil {
		return err
	}
	// Serialize `LastPeriodStart` param:
	err = encoder.Encode(obj.LastPeriodStart)
	if err != nil {
		return err
	}
	// Serialize `CrankFeeRate` param:
	err = encoder.Encode(obj.CrankFeeRate)
	if err != nil {
		return err
	}
	// Serialize `BandwidthDenominator` param:
	err = encoder.Encode(obj.BandwidthDenominator)
	if err != nil {
		return err
	}
	// Serialize `TotalDeposits` param:
	err = encoder.Encode(obj.TotalDeposits)
	if err != nil {
		return err
	}
	return nil
}

func (obj *BidList) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(BidListDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[233 127 13 29 123 209 192 79]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pipeline`:
	err = decoder.Decode(&obj.Pipeline)
	if err != nil {
		return err
	}
	// Deserialize `Book`:
	err = decoder.Decode(&obj.Book)
	if err != nil {
		return err
	}
	// Deserialize `LastPeriodStart`:
	err = decoder.Decode(&obj.LastPeriodStart)
	if err != nil {
		return err
	}
	// Deserialize `CrankFeeRate`:
	err = decoder.Decode(&obj.CrankFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `BandwidthDenominator`:
	err = decoder.Decode(&obj.BandwidthDenominator)
	if err != nil {
		return err
	}
	// Deserialize `TotalDeposits`:
	err = decoder.Decode(&obj.TotalDeposits)
	if err != nil {
		return err
	}
	return nil
}

type Controller struct {
	ControllerBump         uint8
	PcVaultBump            uint8
	Admin                  ag_solanago.PublicKey
	CrankAuthority         ag_solanago.PublicKey
	PcVault                ag_solanago.PublicKey
	PcMint                 ag_solanago.PublicKey
	PcVaultAllocatedAmount uint64
}

var ControllerDiscriminator = [8]byte{184, 79, 171, 0, 183, 43, 113, 110}

func (obj Controller) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ControllerDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `ControllerBump` param:
	err = encoder.Encode(obj.ControllerBump)
	if err != nil {
		return err
	}
	// Serialize `PcVaultBump` param:
	err = encoder.Encode(obj.PcVaultBump)
	if err != nil {
		return err
	}
	// Serialize `Admin` param:
	err = encoder.Encode(obj.Admin)
	if err != nil {
		return err
	}
	// Serialize `CrankAuthority` param:
	err = encoder.Encode(obj.CrankAuthority)
	if err != nil {
		return err
	}
	// Serialize `PcVault` param:
	err = encoder.Encode(obj.PcVault)
	if err != nil {
		return err
	}
	// Serialize `PcMint` param:
	err = encoder.Encode(obj.PcMint)
	if err != nil {
		return err
	}
	// Serialize `PcVaultAllocatedAmount` param:
	err = encoder.Encode(obj.PcVaultAllocatedAmount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Controller) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ControllerDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[184 79 171 0 183 43 113 110]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `ControllerBump`:
	err = decoder.Decode(&obj.ControllerBump)
	if err != nil {
		return err
	}
	// Deserialize `PcVaultBump`:
	err = decoder.Decode(&obj.PcVaultBump)
	if err != nil {
		return err
	}
	// Deserialize `Admin`:
	err = decoder.Decode(&obj.Admin)
	if err != nil {
		return err
	}
	// Deserialize `CrankAuthority`:
	err = decoder.Decode(&obj.CrankAuthority)
	if err != nil {
		return err
	}
	// Deserialize `PcVault`:
	err = decoder.Decode(&obj.PcVault)
	if err != nil {
		return err
	}
	// Deserialize `PcMint`:
	err = decoder.Decode(&obj.PcMint)
	if err != nil {
		return err
	}
	// Deserialize `PcVaultAllocatedAmount`:
	err = decoder.Decode(&obj.PcVaultAllocatedAmount)
	if err != nil {
		return err
	}
	return nil
}

type PeriodRing struct {
	Pipeline ag_solanago.PublicKey
	Ring     []Period
	Start    uint16
	Length   uint16
}

var PeriodRingDiscriminator = [8]byte{61, 191, 59, 143, 226, 235, 104, 26}

func (obj PeriodRing) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PeriodRingDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pipeline` param:
	err = encoder.Encode(obj.Pipeline)
	if err != nil {
		return err
	}
	// Serialize `Ring` param:
	err = encoder.Encode(obj.Ring)
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
	return nil
}

func (obj *PeriodRing) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PeriodRingDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[61 191 59 143 226 235 104 26]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pipeline`:
	err = decoder.Decode(&obj.Pipeline)
	if err != nil {
		return err
	}
	// Deserialize `Ring`:
	err = decoder.Decode(&obj.Ring)
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
	return nil
}

type Pipeline struct {
	Controller   ag_solanago.PublicKey
	Validator    ag_solanago.PublicKey
	PipelineBump uint8
	Admin        ag_solanago.PublicKey
	CrankFeeRate [2]uint64
	Address      ProxyAddress
	Periods      ag_solanago.PublicKey
	Bids         ag_solanago.PublicKey
}

var PipelineDiscriminator = [8]byte{30, 82, 16, 218, 196, 77, 115, 224}

func (obj Pipeline) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PipelineDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Controller` param:
	err = encoder.Encode(obj.Controller)
	if err != nil {
		return err
	}
	// Serialize `Validator` param:
	err = encoder.Encode(obj.Validator)
	if err != nil {
		return err
	}
	// Serialize `PipelineBump` param:
	err = encoder.Encode(obj.PipelineBump)
	if err != nil {
		return err
	}
	// Serialize `Admin` param:
	err = encoder.Encode(obj.Admin)
	if err != nil {
		return err
	}
	// Serialize `CrankFeeRate` param:
	err = encoder.Encode(obj.CrankFeeRate)
	if err != nil {
		return err
	}
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `Periods` param:
	err = encoder.Encode(obj.Periods)
	if err != nil {
		return err
	}
	// Serialize `Bids` param:
	err = encoder.Encode(obj.Bids)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Pipeline) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PipelineDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[30 82 16 218 196 77 115 224]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Controller`:
	err = decoder.Decode(&obj.Controller)
	if err != nil {
		return err
	}
	// Deserialize `Validator`:
	err = decoder.Decode(&obj.Validator)
	if err != nil {
		return err
	}
	// Deserialize `PipelineBump`:
	err = decoder.Decode(&obj.PipelineBump)
	if err != nil {
		return err
	}
	// Deserialize `Admin`:
	err = decoder.Decode(&obj.Admin)
	if err != nil {
		return err
	}
	// Deserialize `CrankFeeRate`:
	err = decoder.Decode(&obj.CrankFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `Periods`:
	err = decoder.Decode(&obj.Periods)
	if err != nil {
		return err
	}
	// Deserialize `Bids`:
	err = decoder.Decode(&obj.Bids)
	if err != nil {
		return err
	}
	return nil
}
