once :
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/grpclog
	go get -u github.com/golang/protobuf/protoc-gen-go

prebuild : 
	protoc --go_out=paths=source_relative:. node/node.proto

