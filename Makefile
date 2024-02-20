generate_go:
	mkdir -p los/generated
	protoc --proto_path=./los/protobuf --go_out=./los/generated --go-grpc_out=require_unimplemented_servers=false:./los/generated ./los/protobuf/*.proto

	mkdir -p mambu/generated
	protoc --proto_path=./mambu/protobuf --go_out=./mambu/generated --go-grpc_out=require_unimplemented_servers=false:./mambu/generated ./mambu/protobuf/*.proto
