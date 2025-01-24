package main

import (
	"context"
	"log"
	"time"

	pb "user_attestor_module/proto/user_attestor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	const socketPath = "/tmp/user_attestor_module.sock"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.NewClient(
		"unix://"+socketPath,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAttestationServiceClient(conn)

	res, err := client.GetUserAttestation(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Could not get attestation: %v", err)
	}

	log.Printf("Token: %s", res.Token)
	log.Printf("Name: %s", res.UserInfo.Name)
	log.Printf("Secret: %s", res.UserInfo.Secret)
	log.Printf("UserID: %s", res.UserInfo.SystemInfo.UserId)
	log.Printf("Username: %s", res.UserInfo.SystemInfo.Username)
	log.Printf("GroupID: %s", res.UserInfo.SystemInfo.GroupId)
	log.Printf("GroupName: %s", res.UserInfo.SystemInfo.GroupName)

	for _, group := range res.UserInfo.SystemInfo.SupplementaryGroups {
		log.Printf("Supplementary GroupID: %s", group.GroupId)
		log.Printf("Supplementary GroupName: %s", group.GroupName)
	}
}
