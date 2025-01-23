package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	// Define the Unix socket path
	socketPath := "/tmp/grpc.sock"

	// Remove the socket file if it already exists
	if err := os.RemoveAll(socketPath); err != nil {
		log.Fatalf("Failed to remove existing socket: %v", err)
	}

	// Create a Unix domain socket listener
	unixListener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to listen on Unix socket: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register your gRPC services here
	// pb.RegisterYourServiceServer(grpcServer, &yourServiceServer{})

	// Start serving on the Unix socket
	log.Printf("gRPC server listening on %s", socketPath)
	if err := grpcServer.Serve(unixListener); err != nil {
		log.Fatalf("Failed to serve gRPC server over Unix socket: %v", err)
	}
}
