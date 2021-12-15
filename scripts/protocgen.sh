#!/usr/bin/env bash

set -eo pipefail

if [ -z "$SOLPB_DIR" ]; then
    echo "variable SOLPB_DIR must be set"
    exit 1
fi

protoc_gen_gocosmos() {
  if ! grep "github.com/gogo/protobuf => github.com/regen-network/protobuf" go.mod &>/dev/null ; then
    echo -e "\tPlease run this command from somewhere inside the cosmos-sdk folder."
    return 1
  fi

  go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos 2>/dev/null
}

protoc_gen_gocosmos

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  buf protoc \
  -I "proto" \
  -I "third_party/proto" \
  -I "$SOLPB_DIR/protobuf-solidity/src/protoc/include" \
  --gocosmos_out=plugins=interfacetype+grpc,\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')

  # command to generate gRPC gateway (*.pb.gw.go in respective modules) files
  buf protoc \
  -I "proto" \
  -I "third_party/proto" \
  -I "$SOLPB_DIR/protobuf-solidity/src/protoc/include" \
  --grpc-gateway_out=logtostderr=true:. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')

done

cp -r github.com/hyperledger-labs/yui-ibc-solidity/* ./
rm -rf github.com
