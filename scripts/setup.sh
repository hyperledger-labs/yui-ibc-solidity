#!/bin/bash
set -ex

TRUFFLE="npx truffle"

function before_common() {
    if [ -n "$NO_GEN_CODE" ]; then
        return
    fi
    ./scripts/solpb.sh
}

function after_common() {
    srcs=(
        "IBCStore"
        "IBCModule"
        "IBFT2Client"
        "SimpleTokenModule"
    )
    if [ -n "$NO_GEN_CODE" ]; then
        return
    fi
    for src in "${srcs[@]}" ; do
        make abi SOURCE=${src}
    done
}

function chain() {
    if [ -z "$network" ]; then
        echo "variable network must be set"
        exit 1
    fi
    if [ -z "$CONF_TPL" ]; then
        echo "variable CONF_TPL must be set"
        exit 1
    fi

    pushd ./chains/besu && docker-compose up -d ${network} && popd
    # XXX Wait for the first block to be created
    sleep 3
    ${TRUFFLE} migrate --reset --network=${network}
    ${TRUFFLE} exec ./scripts/confgen.js --network=${network}
}

function development {
    before_common

    network=development
    export CONF_TPL="./pkg/consts/contract.go:./scripts/template/contract.go.tpl"
    chain

    after_common
}

function testonechain {
    before_common

    network=testchain0
    export CONF_TPL="./tests/e2e/config/chain0/contract.go:./scripts/template/contract.go.tpl"
    chain

    after_common
}

function testtwochainz {
    before_common

    testonechain

    network=testchain1
    export CONF_TPL="./tests/e2e/config/chain1/contract.go:./scripts/template/contract.go.tpl"
    chain

    after_common
}

function down {
    pushd ./chains/besu && docker-compose down && popd
}

subcommand="$1"
shift

case $subcommand in
    development)
        development
        ;;
    testonechain)
        testonechain
        ;;
    testtwochainz)
        testtwochainz
        ;;
    down)
        down
        ;;
    *)
        echo "unknown command '$subcommand'"
        ;;
esac
