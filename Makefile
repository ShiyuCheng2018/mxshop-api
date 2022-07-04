

.PHONY: proto
proto:
	  protoc -I ./user-web/proto ./user-web/proto/user.proto --go_out=plugins=grpc:.
