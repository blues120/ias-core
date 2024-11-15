GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	CONFIG_PROTO_FILES=$(shell $(Git_Bash) -c "find conf -name *.proto")
	ERROR_PROTO_FILES=$(shell $(Git_Bash) -c "find errors -name *.proto")
else
	CONFIG_PROTO_FILES=$(shell find conf -name *.proto)
	ERROR_PROTO_FILES=$(shell find errors -name *.proto)
endif

.PHONY: error
# generate api proto
error:
	protoc --proto_path=./errors \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./errors \
 	       --go-errors_out=paths=source_relative:./errors \
	       $(ERROR_PROTO_FILES)

.PHONY: config
# generate config proto
config:
	protoc --proto_path=./conf \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./conf \
	       $(CONFIG_PROTO_FILES)

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# generate all
all:
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
