#!/usr/bin/env bash
set -e

if [ -z "$SOLPB_DIR" ]; then
    echo "variable CONF_TPL must be set"
    exit 1
fi

for file in proto/*
do
  echo "Generating "$file
  protoc -I$(pwd)/proto -I${SOLPB_DIR}/protobuf-solidity/src/protoc/include --plugin=protoc-gen-sol=${SOLPB_DIR}/protobuf-solidity/src/protoc/plugin/gen_sol.py --"sol_out=gen_runtime=ProtoBufRuntime.sol&solc_version=0.6.8:$(pwd)/contracts/core/types/" $(pwd)/$file
done

npx truffle compile
