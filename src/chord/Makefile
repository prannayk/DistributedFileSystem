install :
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/grpclog
	go get -u github.com/golang/protobuf/protoc-gen-go

prebuild : install
	protoc --go_out=paths=source_relative:. node/node.proto

build : prebuild install
	go build

