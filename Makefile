NETWORK ?= development
TRUFFLE ?= npx truffle
ABIGEN ?= abigen

.PHONY: build
build:
	go build -o ./build/cmd/ibcsol ./cmd

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
	protoc --go_out=. ./proto/*.proto

.PHONY: test
test:
	go test -v ./tests/... -count=1

.PHONY: setup
setup:
	./scripts/setup.sh development

.PHONY: down
down:
	./scripts/setup.sh down

.PHONY: proto-gen
proto-gen:
	@echo "Generating Protobuf files"
	docker run -v $(CURDIR):/workspace --workdir /workspace tendermintdev/sdk-proto-gen sh ./scripts/protocgen.sh

.PHONY: e2e-test
e2e-test:
	go test -v ./tests/e2e/...
