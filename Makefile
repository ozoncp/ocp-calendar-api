.PHONY: build
build: vendor-proto .generate .build

.PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ocp-calendar-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-calendar-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-calendar-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-calendar-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-calendar-api \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				api/ocp-calendar-api/ocp-calendar-api.proto
		mv pkg/ocp-calendar-api/github.com/ozoncp/ocp-calendar-api/* pkg/ocp-calendar-api/
		rm -rf pkg/ocp-calendar-api/github.com
		mkdir -p cmd/ocp-calendar-api

.PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-calendar-api cmd/ocp-calendar-api/main.go

.PHONY: install
install: build .install

.PHONY: .install
install:
		go install cmd/grpc-server/main.go

.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-calendar-api
		cp api/ocp-calendar-api/ocp-calendar-api.proto vendor.protogen/api/ocp-calendar-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u google.golang.org/grpc
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go get github.com/envoyproxy/protoc-gen-validate
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
