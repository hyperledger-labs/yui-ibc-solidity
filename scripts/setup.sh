#!/usr/bin/env bash
set -ex

TRUFFLE="npx truffle"

function launch_chain() {
    if [ -z "$network" ]; then
        echo "variable network must be set"
        exit 1
    fi
    if [ -z "$CONF_TPL" ]; then
        echo "variable CONF_TPL must be set"
        exit 1
    fi

    pushd ./chains && docker compose up -d ${network} && popd
    ${TRUFFLE} compile
    ${TRUFFLE} migrate --reset --compile-none --network=${network}
    ${TRUFFLE} exec ./scripts/confgen.js --network=${network}
}

function development {
    network=development
    export CONF_TPL="./pkg/consts/contract.go:./scripts/template/contract.go.tpl"
    launch_chain
}

function testonechain {
    network=testchain0
    export CONF_TPL="./tests/e2e/config/chain0/contract.go:./scripts/template/contract.go.tpl"
    launch_chain
}

function testtwochainz {
    testonechain
    network=testchain1
    export CONF_TPL="./tests/e2e/config/chain1/contract.go:./scripts/template/contract.go.tpl"
    launch_chain
}

function down {
    pushd ./chains && docker compose down && popd
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
