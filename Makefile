.PHONY: build
build:
	go run ./cmd/main.go --config="./config/local.yaml"

.DEFAULT_GOAL: build