generante-pb: 
	protoc --go_out=./protobuf/helloworld --go_opt=paths=source_relative --go-grpc_out=./protobuf/helloworld --go-grpc_opt=paths=source_relative  ./protobuf/helloworld.proto
