package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "user_attestor_module/proto/user_attestor"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAttestationServiceServer
}

func (s *server) GetUserAttestation(ctx context.Context, in *pb.Empty) (*pb.UserAttestation, error) {
	fmt.Println("Received request")

	// TODO: Get real user attestation data
	return &pb.UserAttestation{
		Token: "sample_token",
		UserInfo: &pb.UserInfo{
			Name:   "John Doe",
			Secret: "supersecret",
			SystemInfo: &pb.SystemInfo{
				UserId:    "12345",
				Username:  "jdoe",
				GroupId:   "group1",
				GroupName: "developers",
				SupplementaryGroups: []*pb.GroupInfo{
					{
						GroupId:   "group2",
						GroupName: "admins",
					},
				},
			},
		},
	}, nil
}

func main() {
	const socketPath = "/tmp/user_attestor_module.sock"

	// Remove the socket file if it already exists
	if err := os.RemoveAll(socketPath); err != nil {
		log.Fatalf("Failed to remove socket file: %v", err)
	}

	lis, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAttestationServiceServer(grpcServer, &server{})

	log.Printf("Server listening on unix://%s", socketPath)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
