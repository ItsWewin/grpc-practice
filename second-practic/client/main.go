package main

import (
	"context"
	"log"
	"time"

	pb "grpc-learn/second-practic/chat"

	"google.golang.org/grpc"
)

var questions = [...]*pb.Request{
	&pb.Request{Question: "question1"},
	&pb.Request{Question: "question2"},
	&pb.Request{Question: "question3"},
	&pb.Request{Question: "question4"},
	&pb.Request{Question: "question5"},
}

func delClientStream(client pb.ChatClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.QA(ctx)

	for _, question := range questions {
		if err := stream.Send(question); err != nil {
			log.Fatalf("Some error occurred when send message to server, error: %v", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("some error occurred when get response from server, error: %v", err)
	}
	log.Println("Get answer from server: ", res.Answer)
}

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect server error: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatClient(conn)
	delClientStream(client)
}
