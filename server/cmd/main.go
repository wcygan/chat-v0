package main

import (
	pb "github.com/wcygan/chat-v0/generated/go/chat/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) Chat(stream pb.ChatService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Received message from %s: %s", msg.User, msg.Message)
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})
	log.Printf("Server listening on %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
