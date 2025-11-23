PROJECT_NAME=$(basename $(pwd))
BUILD_DIR=./config/examples

CONFIG_PATH=$(BUILD_DIR)/config.yaml

generate: conf_prepare
	swag --version
	swag fmt ./...
	go generate -ldflags="$(ldflags)" ./...
	go mod tidy

lint:
	go vet ./...
	golangci-lint run $(GOLINT_ARGS)

run: conf_prepare
	go run -tags=jsoniter -ldflags="$(ldflags)" ./cmd -c "$(CONFIG_PATH)"

conf_prepare:
	mkdir -p "$(BUILD_DIR)" && (cp -n ./examples/configs/config.yaml "$(CONFIG_PATH)" 2>/dev/null || true)
