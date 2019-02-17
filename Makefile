-include .env

test:
    env GO111MODULE=on go test -v ./...
