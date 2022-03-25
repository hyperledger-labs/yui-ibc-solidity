NETWORK ?= development
TRUFFLE ?= npx truffle
ABIGEN ?= abigen

.PHONY: config
config:
	export CONF_TPL="./pkg/consts/contract.go:./scripts/template/contract.go.tpl" && $(TRUFFLE) exec ./scripts/confgen.js --network=$(NETWORK)

.PHONY: abi
abi:
ifdef SOURCE
	$(eval TARGET := $(shell echo ${SOURCE} | tr A-Z a-z))
	@mkdir -p ./build/abi ./pkg/contract
	@mkdir -p ./pkg/contract/$(TARGET)
	@cat ./build/contracts/${SOURCE}.json | jq ".abi" > ./build/abi/${SOURCE}.abi
	$(ABIGEN) --abi ./build/abi/${SOURCE}.abi --pkg $(TARGET) --out ./pkg/contract/$(TARGET)/$(TARGET).go
else
	@echo "'SOURCE={SOURCE}' is required"
endif

.PHONY: proto
proto:
ifndef SOLPB_DIR
	$(error SOLPB_DIR is not specified)
else
	protoc --go_out=. \
		-I./proto \
		-I./third_party/proto \
		-I$(SOLPB_DIR)/protobuf-solidity/src/protoc/include \
		./proto/**/*.proto
endif

.PHONY: test
test:
	go test -v ./tests/... -count=1

.PHONY: setup
setup:
	./scripts/setup.sh development

.PHONY: setup-e2e
setup-e2e:
	./scripts/setup.sh testtwochainz

.PHONY: down
down:
	./scripts/setup.sh down

.PHONY: proto-gen
proto-gen:
	@echo "Generating Protobuf files"
ifndef SOLPB_DIR
	$(error SOLPB_DIR is not specified)
else
	docker run \
		-v $(CURDIR):/workspace \
		-v $(SOLPB_DIR):/solpb \
		-e SOLPB_DIR=/solpb \
		--workdir /workspace \
		tendermintdev/sdk-proto-gen \
		sh ./scripts/protocgen.sh
endif

.PHONY: integration-test
integration-test:
	go test -v ./tests/integration/... -count=1

.PHONY: e2e-test
e2e-test:
	go test -v ./tests/e2e/... -count=1
