// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package solmate_cba

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/SolmateDev/solana-go"
)

type BidAuth struct {
	User  ag_solanago.PublicKey
	Start int64
	Nonce int64
}

func (obj BidAuth) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `Start` param:
	err = encoder.Encode(obj.Start)
	if err != nil {
		return err
	}
	// Serialize `Nonce` param:
	err = encoder.Encode(obj.Nonce)
	if err != nil {
		return err
	}
	return nil
}

func (obj *BidAuth) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `Start`:
	err = decoder.Decode(&obj.Start)
	if err != nil {
		return err
	}
	// Deserialize `Nonce`:
	err = decoder.Decode(&obj.Nonce)
	if err != nil {
		return err
	}
	return nil
}

type Bid struct {
	IsBlank             bool
	User                ag_solanago.PublicKey
	Refund              ag_solanago.PublicKey
	Deposit             uint64
	BandwidthAllocation uint64
}

func (obj Bid) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `IsBlank` param:
	err = encoder.Encode(obj.IsBlank)
	if err != nil {
		return err
	}
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `Refund` param:
	err = encoder.Encode(obj.Refund)
	if err != nil {
		return err
	}
	// Serialize `Deposit` param:
	err = encoder.Encode(obj.Deposit)
	if err != nil {
		return err
	}
	// Serialize `BandwidthAllocation` param:
	err = encoder.Encode(obj.BandwidthAllocation)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Bid) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `IsBlank`:
	err = decoder.Decode(&obj.IsBlank)
	if err != nil {
		return err
	}
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `Refund`:
	err = decoder.Decode(&obj.Refund)
	if err != nil {
		return err
	}
	// Deserialize `Deposit`:
	err = decoder.Decode(&obj.Deposit)
	if err != nil {
		return err
	}
	// Deserialize `BandwidthAllocation`:
	err = decoder.Decode(&obj.BandwidthAllocation)
	if err != nil {
		return err
	}
	return nil
}

type Period struct {
	IsBlank   bool
	Withhold  uint64
	Start     uint64
	Length    uint64
	DecayRate [2]uint64
}

func (obj Period) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `IsBlank` param:
	err = encoder.Encode(obj.IsBlank)
	if err != nil {
		return err
	}
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
	// Serialize `DecayRate` param:
	err = encoder.Encode(obj.DecayRate)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Period) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `IsBlank`:
	err = decoder.Decode(&obj.IsBlank)
	if err != nil {
		return err
	}
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
	// Deserialize `DecayRate`:
	err = decoder.Decode(&obj.DecayRate)
	if err != nil {
		return err
	}
	return nil
}

type ProxyAddress struct {
	Url []byte
}

func (obj ProxyAddress) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Url` param:
	err = encoder.Encode(obj.Url)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ProxyAddress) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Url`:
	err = decoder.Decode(&obj.Url)
	if err != nil {
		return err
	}
	return nil
}

type PeriodStatus ag_binary.BorshEnum

const (
	PeriodStatusHasFinished PeriodStatus = iota
	PeriodStatusInProgress
	PeriodStatusWillStart
)

func (value PeriodStatus) String() string {
	switch value {
	case PeriodStatusHasFinished:
		return "HasFinished"
	case PeriodStatusInProgress:
		return "InProgress"
	case PeriodStatusWillStart:
		return "WillStart"
	default:
		return ""
	}
}
