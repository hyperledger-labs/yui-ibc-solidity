FORGE ?= forge
SOLC_VERSION ?= 0.8.9
ABIGEN ?= "docker run -v .:/workspace -w /workspace -it ethereum/client-go:alltools-v1.11.6 abigen"
FORGE_SNAPSHOT_OPTION ?= --diff

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

.PHONY: setup
setup:
	./scripts/setup.sh development

.PHONY: setup-e2e
setup-e2e:
	./scripts/setup.sh testtwochainz

.PHONY: down
down:
	./scripts/setup.sh down

.PHONY: integration-test
integration-test:
	go test -v ./tests/integration/... -count=1

.PHONY: e2e-test
e2e-test:
	go test -v ./tests/e2e/... -count=1

.PHONY: abigen
abigen:
	ABIGEN=$(ABIGEN) ./scripts/abigen.sh

.PHONY: yrly
yrly:
	go build -o build/yrly ./relayer
