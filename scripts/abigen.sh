#!/usr/bin/env bash
set -e

if [ -z "$ABIGEN" ]; then
  echo 'Please set the "ABIGEN" environment variable and try again.' >&2
  exit 1
fi

function gen_code() {
    local source=$1;
    local target=$(echo ${source} | tr A-Z a-z)

    mkdir -p ./build/abi ./pkg/contract/${target}
	jq -r '.abi' ./out/${source}.sol/${source}.json > ./build/abi/${source}.abi
	${ABIGEN} --abi ./build/abi/${source}.abi --pkg ${target} --out ./pkg/contract/${target}/${target}.go
}

function main() {
    local srcs=(
        "IBCHandler"
        "SimpleToken"
        "ICS20TransferBank"
        "ICS20Bank"
        "IBCCommitmentTestHelper"
    )
    for src in "${srcs[@]}" ; do
        gen_code ${src}
    done
}

main
