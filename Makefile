GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
SERVICE_NAME=$(shell a=`basename $$PWD` && echo $$a)
SERVICE_NAME_UPPER=$(shell echo $(SERVICE_NAME) | cut -b 1 | tr [a-z] [A-Z])
SERVICE_NAME_UPPER1=$(shell echo $(SERVICE_NAME) | cut -b 2-)
SERVICE_NAME_UPPERAll=$(shell echo $(SERVICE_NAME_UPPER)$(SERVICE_NAME_UPPER1) )
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
DOCKER_IMAGE=$(shell echo $(APP_NAME) |awk -F '@' '{print "kratos-crud-layout/" $$0 ":0.1.0"}')
TEMPLATE_SERVICE_NAME=serviceName
TEMPLATE_SERVICE_NAME_UPPER=ServiceName

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
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
# docker编译
.PHONY: docker
docker:
	docker build -f ./Dockerfile --build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) -t $(DOCKER_IMAGE) .

# 新服务初始化
.PHONY: newServiceInit
newServiceInit:
	@echo $(SERVICE_NAME)
	@echo $(TEMPLATE_SERVICE_NAME)
	@echo $(TEMPLATE_SERVICE_NAME_UPPER)
	@echo $(SERVICE_NAME_UPPERAll)
	# 更改文件名称
	@ cp ./api/$(TEMPLATE_SERVICE_NAME)/v1/$(TEMPLATE_SERVICE_NAME).proto ./api/$(TEMPLATE_SERVICE_NAME)/v1/$(SERVICE_NAME).proto
	# 更新文件夹
	@ mv ./api/$(TEMPLATE_SERVICE_NAME) ./api/$(SERVICE_NAME)
	# 替换proto文件内容
	@ sed -i "s/$(TEMPLATE_SERVICE_NAME_UPPER)/$(SERVICE_NAME_UPPERAll)/g" ./api/$(SERVICE_NAME)/v1/$(SERVICE_NAME).proto
	@ sed -i "s/$(TEMPLATE_SERVICE_NAME)/$(SERVICE_NAME)/g" ./api/$(SERVICE_NAME)/v1/$(SERVICE_NAME).proto
	# 生成客户端文件
	@ kratos proto client ./api/$(SERVICE_NAME)/v1/$(SERVICE_NAME).proto
	# 删除模板文件
	@ rm ./api/$(SERVICE_NAME)/v1/serviceName*
	#替换每一个文件下的内容 (server)
	@ sed -i "s/$(TEMPLATE_SERVICE_NAME_UPPER)/$(SERVICE_NAME_UPPERAll)/g" ./cmd/$(SERVICE_NAME)/wire_gen.go
	#替换每一个文件下的内容 (service)
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME_UPPER)/$(SERVICE_NAME_UPPERAll)/g" ./internal/service/service.go
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME)/$(SERVICE_NAME)/g" ./internal/service/$(TEMPLATE_SERVICE_NAME).go
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME_UPPER)/$(SERVICE_NAME_UPPERAll)/g" ./internal/service/$(TEMPLATE_SERVICE_NAME).go
	# 更改文件
	@ mv ./internal/service/$(TEMPLATE_SERVICE_NAME).go ./internal/service/$(SERVICE_NAME).go
	#替换每一个文件下的内容 (grpc)
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME)/$(SERVICE_NAME)/g" ./internal/server/grpc.go
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME_UPPER)/$(SERVICE_NAME_UPPERAll)/g" ./internal/server/grpc.go
	#替换每一个文件下的内容 (biz)
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME_UPPER)/$(SERVICE_NAME_UPPERAll)/g" ./internal/biz/biz.go
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME)/$(SERVICE_NAME)/g" ./internal/biz/$(TEMPLATE_SERVICE_NAME).go
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME_UPPER)/$(SERVICE_NAME_UPPERAll)/g" ./internal/biz/$(TEMPLATE_SERVICE_NAME).go
	# 更改文件
	@ mv ./internal/biz/$(TEMPLATE_SERVICE_NAME).go ./internal/biz/$(SERVICE_NAME).go
	#替换每一个文件下的内容 (data)
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME_UPPER)/$(SERVICE_NAME_UPPERAll)/g" ./internal/data/data.go
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME)/$(SERVICE_NAME)/g" ./internal/data/$(TEMPLATE_SERVICE_NAME).go
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME_UPPER)/$(SERVICE_NAME_UPPERAll)/g" ./internal/data/$(TEMPLATE_SERVICE_NAME).go
#	# 更改文件
	@ mv ./internal/data/$(TEMPLATE_SERVICE_NAME).go ./internal/data/$(SERVICE_NAME).go
#	#替换每一个文件下的内容 (model)
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME)/$(SERVICE_NAME)/g" ./internal/data/entity/$(TEMPLATE_SERVICE_NAME).go
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME_UPPER)/$(SERVICE_NAME_UPPERAll)/g" ./internal/data/entity/$(TEMPLATE_SERVICE_NAME).go
	# 更改文件
	@ mv ./internal/data/entity/$(TEMPLATE_SERVICE_NAME).go ./internal/data/entity/$(SERVICE_NAME).go
#	#更改配置文件
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME)/$(SERVICE_NAME)/g" ./configs/config.yaml
	@ sed -i  "s/$(TEMPLATE_SERVICE_NAME_UPPER)/$(SERVICE_NAME)/g" ./configs/config.yaml
#	# 更新配置文件
	make config
#	# 拉取引用包
	go mod tidy
