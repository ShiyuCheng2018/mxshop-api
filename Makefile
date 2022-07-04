
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install github.com/golang/protobuf/protoc-gen-go@latest
.PHONY: proto
proto:
	  protoc -I ./user-web/proto ./user-web/proto/user.proto --go_out=plugins=grpc:.
