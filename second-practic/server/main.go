package main

import (
	"fmt"
	pb "grpc-learn/second-practic/chat"

	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

var answer1 = pb.Response{Answer: "answer1"}
var answer2 = pb.Response{Answer: "answer2"}
var answer3 = pb.Response{Answer: "answer3"}
var answers = [...]*pb.Response{&answer1, &answer2, &answer3}

func (s *server) QA(req *pb.Request, stream pb.Chat_QAServer) error {
	for _, answer := range answers {
		fmt.Printf("Send data: %v\n", answer.Answer)
		if err := stream.Send(answer); err != nil {
			fmt.Printf("Some error occurred when send data: %v", err)
			return err
		}
	}
	return nil
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
