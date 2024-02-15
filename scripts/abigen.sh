#!/usr/bin/env bash
set -e

if [ -z "$ABIGEN" ]; then
  echo 'Please set the "ABIGEN" environment variable and try again.' >&2
  exit 1
fi

function gen_code() {
  local source=$1;
  if [ -z $2 ]; then
    local target=$(echo ${source} | tr A-Z a-z)
  else
    local target=$2
  fi

  mkdir -p ./build/abi ./pkg/contract/${target}
  forge inspect ${source} abi > ./build/abi/${source}.abi
  ${ABIGEN} --abi ./build/abi/${source}.abi --pkg ${target} --out ./pkg/contract/${target}/${target}.go
}

function main() {
  local srcs=(
    "IBCHandler"
    "ERC20"
    "ICS20TransferBank"
    "ICS20Bank"
    "IBFT2Client"
    "IBCMockApp"
    "IBCCommitmentTestHelper"
  )
  for src in "${srcs[@]}" ; do
    gen_code ${src}
  done
}

main
