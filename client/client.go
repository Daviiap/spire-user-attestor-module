package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "user_attestor_module/pkg/protos/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const socketPath = "/tmp/grpc_unix.sock"

func main() {
	conn, err := grpc.Dial(
		socketPath,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return net.DialTimeout("unix", addr, time.Second)
		}),
	)
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetUserInfo(ctx, &pb.EmptyMessage{})
	if err != nil {
		log.Fatalf("error calling GetUserInfo: %v", err)
	}

	log.Printf("Received attestation token: %s", response.GetAttestationToken())
	log.Printf("User Info: %v", response.GetUserInfo())
}
