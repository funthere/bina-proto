proto_generate:
	mkdir -p src/los/generate
	protoc --proto_path=./src/los/protobuf --go_out=./src/los/generate --go-grpc_out=require_unimplemented_servers=false:./src/los/generate ./src/los/protobuf/*.proto

	mkdir -p src/mambu/generate
	protoc --proto_path=./src/mambu/protobuf --go_out=./src/mambu/generate --go-grpc_out=require_unimplemented_servers=false:./src/mambu/generate ./src/mambu/protobuf/*.proto
