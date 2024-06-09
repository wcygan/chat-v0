package main

import (
	pb "github.com/wcygan/chat-v0/generated/go/chat/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type server struct {
	pb.UnimplementedChatServiceServer
	activeStreams []pb.ChatService_ChatServer
	mu            sync.Mutex
}

func (s *server) Chat(stream pb.ChatService_ChatServer) error {
	s.mu.Lock()
	s.activeStreams = append(s.activeStreams, stream)
	s.mu.Unlock()

	for {
		msg, err := stream.Recv()
		if err != nil {
			return err  // Handle disconnection or stream errors
		}
		s.mu.Lock()
		for _, clientStream := range s.activeStreams {
			if err := clientStream.Send(msg); err != nil {
				// Handle failed send, possibly removing the clientStream from activeStreams
			}
		}
		s.mu.Unlock()
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
