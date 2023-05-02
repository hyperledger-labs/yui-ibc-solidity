FORGE ?= forge
SOLC_VERSION ?= 0.8.18
ABIGEN ?= "docker run -v .:/workspace -w /workspace -it ethereum/client-go:alltools-v1.11.6 abigen"
DOCKER_COMPOSE ?= docker compose
INTEGRATION_TEST_COMPOSE_FILE ?= ./chains/docker-compose.yml
TEST_BROADCAST_LOG_DIR ?= ./broadcast/Deploy.s.sol
TEST_MNEMONIC ?= "math razor capable expose worth grape metal sunset metal sudden usage scheme"

######## Development ########

.PHONY: fmt
fmt:
	@$(FORGE) fmt $(FORGE_FMT_OPTS) \
		./contracts/core \
		./contracts/apps \
		./contracts/clients \
		./tests/foundry/src

.PHONY: lint
lint:
	@npx solhint 'contracts/{apps,clients,core}/**/*.sol' 'tests/foundry/src/**/*.sol'
	@$(MAKE) FORGE_FMT_OPTS=--check fmt

.PHONY: build
build:
	@forge build --sizes --skip test --use solc:$(SOLC_VERSION)

.PHONY: test
test:
	@forge snapshot -vvvv --gas-report --use solc:$(SOLC_VERSION) $(FORGE_SNAPSHOT_OPTION)

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
	docker run \
		-v $(CURDIR):/workspace \
		-v $(SOLPB_DIR):/solpb \
		-e SOLPB_DIR=/solpb \
		--workdir /workspace \
		tendermintdev/sdk-proto-gen:v0.3 \
		sh ./scripts/protocgen.sh
endif

.PHONY: proto-gen
proto-gen: proto-sol proto-go

######## Integration test ########

.PHONY: network-development
network-development:
	TEST_MNEMONIC=$(TEST_MNEMONIC) $(DOCKER_COMPOSE) -f $(INTEGRATION_TEST_COMPOSE_FILE) up --detach --wait development
	TEST_MNEMONIC=$(TEST_MNEMONIC) $(FORGE) script --use solc:${SOLC_VERSION} --fork-url http://127.0.0.1:8545 --broadcast \
		./tests/foundry/src/Deploy.s.sol

.PHONY: network-e2e
network-e2e:
	$(DOCKER_COMPOSE) -f $(INTEGRATION_TEST_COMPOSE_FILE) up --detach --wait testchain0 testchain1
	TEST_MNEMONIC=$(TEST_MNEMONIC) $(FORGE) script --legacy --use solc:${SOLC_VERSION} --fork-url http://127.0.0.1:8645 --broadcast \
		./tests/foundry/src/Deploy.s.sol
	TEST_MNEMONIC=$(TEST_MNEMONIC) $(FORGE) script --legacy --use solc:${SOLC_VERSION} --fork-url http://127.0.0.1:8745 --broadcast \
		./tests/foundry/src/Deploy.s.sol

.PHONY: network-down
network-down:
	$(DOCKER_COMPOSE) -f $(INTEGRATION_TEST_COMPOSE_FILE) down

.PHONY: integration-test
integration-test:
	TEST_MNEMONIC=$(TEST_MNEMONIC) TEST_BROADCAST_LOG_DIR=$(CURDIR)/$(TEST_BROADCAST_LOG_DIR) go test -v ./tests/integration/... -count=1

.PHONY: e2e-test
e2e-test:
	TEST_MNEMONIC=$(TEST_MNEMONIC) TEST_BROADCAST_LOG_DIR=$(CURDIR)/$(TEST_BROADCAST_LOG_DIR) go test -v ./tests/e2e/... -count=1

.PHONY: abigen
abigen:
	ABIGEN=$(ABIGEN) ./scripts/abigen.sh
