package client

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRevertReasonParser(t *testing.T) {
	// 1. Valid format
	s, err := parseRevertReason(
		hexToBytes("0x08c379a00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001a4e6f7420656e6f7567682045746865722070726f76696465642e000000000000"),
	)
	require.NoError(t, err)
	require.Equal(t, "Not enough Ether provided.", s)

	// 2. Empty bytes
	s, err = parseRevertReason(nil)
	require.NoError(t, err)
	require.Equal(t, "", s)

	// 3. Invalid format
	s, err = parseRevertReason([]byte{0})
	require.Error(t, err)
}

func hexToBytes(s string) []byte {
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
	}
	reason, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return reason
}
