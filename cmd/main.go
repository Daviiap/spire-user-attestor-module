package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"user_attestor_module/internal/interfaceadapters"
	"user_attestor_module/internal/usecases"
	pb "user_attestor_module/pkg/protos/user"
)

const socketPath = "/tmp/grpc_unix.sock"

func main() {
	if err := os.RemoveAll(socketPath); err != nil {
		log.Fatalf("failed to remove existing unix socket file: %v", err)
	}

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("failed to listen on unix socket: %v", err)
	}
	defer listener.Close()

	userInteractor := &usecases.UserInteractor{}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, interfaceadapters.NewUserServer(userInteractor))

	log.Printf("gRPC server is listening on unix socket: %s", socketPath)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
