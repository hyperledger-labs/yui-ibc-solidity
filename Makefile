FOUNDRY_PROFILE=default
FORGE=FOUNDRY_PROFILE=$(FOUNDRY_PROFILE) forge
SOLC_VERSION=0.8.24
EVM_VERSION=paris
DOCKER=docker
ABIGEN="$(DOCKER) run -v .:/workspace -w /workspace -it ethereum/client-go:alltools-v1.11.6 abigen"
SOLHINT=npx solhint
SLITHER=slither
DOCKER_COMPOSE=$(DOCKER) compose
E2E_TEST_COMPOSE_FILE=./chains/compose.yml
TEST_BROADCAST_LOG_DIR=./broadcast/Deploy.s.sol
TEST_MNEMONIC="math razor capable expose worth grape metal sunset metal sudden usage scheme"
TEST_UPGRADEABLE=false

######## Development ########

.PHONY: build
build:
	$(FORGE) build --sizes --skip test --use solc:$(SOLC_VERSION)

.PHONY: clean
clean:
	$(FORGE) clean

.PHONY: fmt
fmt:
	$(FORGE) fmt $(FORGE_FMT_OPTS)

.PHONY: lint
lint:
	@$(SOLHINT) 'contracts/**/*.sol'
	@$(MAKE) FORGE_FMT_OPTS=--check fmt

.PHONY: test
test:
	TEST_UPGRADEABLE=$(TEST_UPGRADEABLE) $(FORGE) test -vvvv --gas-report --isolate --use solc:$(SOLC_VERSION)

.PHONY: snapshot
snapshot:
	$(FORGE) snapshot -vvvv --gas-report --isolate --use solc:$(SOLC_VERSION)

.PHONY: coverage
coverage:
	$(FORGE) coverage --use solc:$(SOLC_VERSION)

.PHONY: slither
slither:
	@$(SLITHER) .

######## Protobuf ########

.PHONY: proto-sol
proto-sol:
ifndef SOLPB_DIR
	$(error SOLPB_DIR is not specified)
else
	./scripts/solpb.sh
endif

.PHONY: proto-go
proto-go:
ifndef SOLPB_DIR
	$(error SOLPB_DIR is not specified)
else
	$(DOCKER) run \
		-v $(CURDIR):/workspace \
		-v $(SOLPB_DIR):/solpb \
		-e SOLPB_DIR=/solpb \
		--workdir /workspace \
		tendermintdev/sdk-proto-gen:v0.3 \
		sh ./scripts/protocgen.sh
endif

.PHONY: proto-gen
proto-gen: proto-sol proto-go

######## Abigen ########

.PHONY: abigen
abigen: build
	ABIGEN=$(ABIGEN) ./scripts/abigen.sh

######## E2E test ########

.PHONY: network-ibft2
network-ibft2:
	$(DOCKER_COMPOSE) -f $(E2E_TEST_COMPOSE_FILE) up --detach --wait ibft2-testchain0 ibft2-testchain1
	$(MAKE) deploy

.PHONY: network-qbft
network-qbft:
	$(DOCKER_COMPOSE) -f $(E2E_TEST_COMPOSE_FILE) up --detach --wait qbft-testchain0 qbft-testchain1
	$(MAKE) deploy

.PHONY: deploy
deploy:
	TEST_UPGRADEABLE=$(TEST_UPGRADEABLE) TEST_MNEMONIC=$(TEST_MNEMONIC) $(FORGE) script --legacy --batch-size 5 --use solc:${SOLC_VERSION} --evm-version ${EVM_VERSION} --fork-url http://127.0.0.1:8645 --broadcast \
		./tests/foundry/src/Deploy.s.sol
	TEST_UPGRADEABLE=$(TEST_UPGRADEABLE) TEST_MNEMONIC=$(TEST_MNEMONIC) $(FORGE) script --legacy --batch-size 5 --use solc:${SOLC_VERSION} --evm-version ${EVM_VERSION} --fork-url http://127.0.0.1:8745 --broadcast \
		./tests/foundry/src/Deploy.s.sol

.PHONY: network-down
network-down:
	$(DOCKER_COMPOSE) -f $(E2E_TEST_COMPOSE_FILE) down

.PHONY: e2e-test
e2e-test:
	TEST_UPGRADEABLE=$(TEST_UPGRADEABLE) TEST_MNEMONIC=$(TEST_MNEMONIC) TEST_BROADCAST_LOG_DIR=$(CURDIR)/$(TEST_BROADCAST_LOG_DIR) go test -v ./tests/e2e/... -count=1
