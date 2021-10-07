GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: init
# init env
init:
	go get -u github.com/go-kratos/kratos/cmd/kratos/v2
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2

.PHONY: errors
# generate errors code
errors:
	protoc --proto_path=. \
               --proto_path=./third_party \
               --go_out=paths=source_relative:. \
               --go-errors_out=paths=source_relative:. \
               $(API_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	make build-message
	make build-user
	make build-course
	make build-apigateway
	make build-edgeUser
	make build-edgeCourse

.PHONY: build-message
# build-message
build-message:
	mkdir -p message/bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./message/bin/ ./app/message/...

.PHONY: build-user
# build-user
build-user:
	mkdir -p user/bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./user/bin/ ./app/user/...

.PHONY: build-course
# build-course
build-course:
	mkdir -p course/bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./course/bin/ ./app/course/...

.PHONY: build-apigateway
# build-apigateway
build-apigateway:
	mkdir -p apigateway/bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./apigateway/bin/ ./app/api-gateway/...

.PHONY: build-edgeUser
# build-edgeUser
build-edgeUser:
	mkdir -p edgeUser/bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./edgeUser/bin/ ./app/edge/user/...

.PHONY: build-edgeCourse
# build-edgeCourse
build-edgeCourse:
	mkdir -p edgeCourse/bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./edgeCourse/bin/ ./app/edge/course/...


.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
	make errors;
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
