package main

import (
	pb "github.com/chrisdamba/usersearchgo/proto"
	"github.com/chrisdamba/usersearchgo/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := &UserServiceServer{
		service: services.NewUserService(),
	}
	pb.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
