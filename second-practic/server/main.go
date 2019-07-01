package main

import (
	"context"
	pb "grpc-learn/second-practic/chat"

	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) QA(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Answer: "Search Answer of " + in.Question}, nil
}

func main() {
	l, err := net.Listen("tcp", "localhost:50001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
