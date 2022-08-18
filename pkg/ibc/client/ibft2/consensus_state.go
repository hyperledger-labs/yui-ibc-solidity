package ibft2

import "github.com/cosmos/ibc-go/v4/modules/core/exported"

var _ exported.ConsensusState = (*ConsensusState)(nil)

func (cons ConsensusState) ClientType() string {
	return IBFT2Client
}

// GetRoot returns the commitment root of the consensus state,
// which is used for key-value pair verification.
func (cons ConsensusState) GetRoot() exported.Root {
	panic("not implemented")
}

// GetTimestamp returns the timestamp (in nanoseconds) of the consensus state
func (cons ConsensusState) GetTimestamp() uint64 {
	return cons.Timestamp
}

func (cons ConsensusState) ValidateBasic() error {
	return nil
}
