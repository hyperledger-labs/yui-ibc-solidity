package client

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRevertReasonParser(t *testing.T) {
	erepo := NewErrorsRepository()
	s, args, err := erepo.ParseError(
		hexToBytes("0x08c379a00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001a4e6f7420656e6f7567682045746865722070726f76696465642e000000000000"),
	)
	require.NoError(t, err)
	require.Equal(t, "Error(string)", s)
	require.Equal(t, []interface{}{"Not enough Ether provided."}, args)
}

func hexToBytes(s string) []byte {
	reason, err := hex.DecodeString(strings.TrimPrefix(s, "0x"))
	if err != nil {
		panic(err)
	}
	return reason
}
