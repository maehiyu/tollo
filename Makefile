.PHONY: all proto-gen-go

all: proto-gen-go

proto-gen-go:
	@echo "Generating Go protobuf code..."
	@mkdir -p gen/go
	protoc --proto_path=protos \
		--go_out=gen/go --go_opt=module=github.com/maehiyu/tollo \
    	--go-grpc_out=gen/go --go-grpc_opt=module=github.com/maehiyu/tollo \
		$(shell find protos -name '*.proto')
  