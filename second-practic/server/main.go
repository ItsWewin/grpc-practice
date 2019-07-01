package main

import (
	"bytes"
	"fmt"
	pb "grpc-learn/second-practic/chat"
	"io"

	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

var answer1 = pb.Response{Answer: "answer1"}
var answer2 = pb.Response{Answer: "answer2"}
var answer3 = pb.Response{Answer: "answer3"}
var answers = [...]*pb.Response{&answer1, &answer2, &answer3}

func (s *server) QA(stream pb.Chat_QAServer) error {
	var buffer bytes.Buffer
	for {
		question, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Response{
				Answer: buffer.String(),
			})
		}
		if err != nil {
			return err
		}
		log.Println("question: ", question.Question)
		buffer.WriteString(fmt.Sprintf("%s ", question.Question))
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
