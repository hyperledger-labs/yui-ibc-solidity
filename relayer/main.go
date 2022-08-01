package main

import (
	"log"

	ethereum "github.com/hyperledger-labs/yui-ibc-solidity/pkg/relay/module"
	tendermint "github.com/hyperledger-labs/yui-relayer/chains/tendermint/module"
	"github.com/hyperledger-labs/yui-relayer/cmd"
	mock "github.com/hyperledger-labs/yui-relayer/provers/mock/module"
)

func main() {
	if err := cmd.Execute(
		tendermint.Module{},
		ethereum.Module{},
		mock.Module{},
	); err != nil {
		log.Fatal(err)
	}
}
