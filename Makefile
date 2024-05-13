FORGE ?= forge
SOLC_VERSION ?= 0.8.24
ABIGEN ?= "docker run -v .:/workspace -w /workspace -it ethereum/client-go:alltools-v1.11.6 abigen"
DOCKER_COMPOSE ?= docker compose
E2E_TEST_COMPOSE_FILE ?= ./chains/docker-compose.yml
TEST_BROADCAST_LOG_DIR ?= ./broadcast/Deploy.s.sol
TEST_MNEMONIC ?= "math razor capable expose worth grape metal sunset metal sudden usage scheme"

######## Development ########

.PHONY: build
build:
	@forge build --sizes --skip test --use solc:$(SOLC_VERSION)

.PHONY: fmt
fmt:
	@$(FORGE) fmt $(FORGE_FMT_OPTS) \
		./contracts/apps \
		./contracts/clients \
		./contracts/core \
		./contracts/helpers \
		./tests/foundry/src

.PHONY: lint
lint:
	@npx solhint 'contracts/{apps,clients,core}/**/*.sol' 'tests/foundry/src/**/*.sol'
	@$(MAKE) FORGE_FMT_OPTS=--check fmt

.PHONY: test
test:
	@forge snapshot -vvvv --gas-report --use solc:$(SOLC_VERSION) $(FORGE_SNAPSHOT_OPTION)

.PHONY: coverage
coverage:
	@forge coverage --use solc:$(SOLC_VERSION)

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
	TEST_MNEMONIC=$(TEST_MNEMONIC) $(FORGE) script --legacy --batch-size 5 --use solc:${SOLC_VERSION} --fork-url http://127.0.0.1:8645 --broadcast \
		./tests/foundry/src/Deploy.s.sol
	TEST_MNEMONIC=$(TEST_MNEMONIC) $(FORGE) script --legacy --batch-size 5 --use solc:${SOLC_VERSION} --fork-url http://127.0.0.1:8745 --broadcast \
		./tests/foundry/src/Deploy.s.sol

.PHONY: network-down
network-down:
	$(DOCKER_COMPOSE) -f $(E2E_TEST_COMPOSE_FILE) down

.PHONY: e2e-test
e2e-test:
	TEST_MNEMONIC=$(TEST_MNEMONIC) TEST_BROADCAST_LOG_DIR=$(CURDIR)/$(TEST_BROADCAST_LOG_DIR) go test -v ./tests/e2e/... -count=1
