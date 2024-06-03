.PHONY: all clean build run test proto

all: proto build

clean:
	rm -f bin/usersearchgo
	rm -rf proto/*.pb.go

proto:
	protoc --go_out=. --go-grpc_out=. proto/user.proto

build:
	go build -o bin/usersearchgo main.go grpc_server.go

run: build
	./bin/usersearchgo

test:
	go test -v ./...

generate:
	go generate ./...

