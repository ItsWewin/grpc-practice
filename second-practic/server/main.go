package main

import (
	"context"
	pb "grpc-learn/second-practic/chat"

	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.Reply{ReplyWorld: "Hello " + in.Name}, nil
}

func main() {
	l, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
