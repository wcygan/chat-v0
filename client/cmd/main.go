package main

import (
	"bufio"
	"context"
	pb "github.com/wcygan/chat-v0/generated/go/chat/v1"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)

	stream, err := c.Chat(context.Background())
	if err != nil {
		log.Fatalf("could not open stream: %v", err)
	}

	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				log.Fatalf("Failed to receive a message: %v", err)
			}
			log.Printf("Received message from %s: %s", in.User, in.Message)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if err := stream.Send(&pb.ChatMessage{User: "Client", Message: text}); err != nil {
			log.Fatalf("Failed to send a message: %v", err)
		}
	}

	stream.CloseSend()
}
