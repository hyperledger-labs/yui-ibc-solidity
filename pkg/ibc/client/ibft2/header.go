package ibft2

import (
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v4/modules/core/exported"
)

var _ exported.Header = (*Header)(nil)

func (h Header) ClientType() string {
	return IBFT2Client
}

func (h Header) GetHeight() exported.Height {
	return clienttypes.Height(h.TrustedHeight)
}

func (h Header) ValidateBasic() error {
	// TODO implement this
	return nil
}
