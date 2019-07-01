package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "grpc-learn/second-practic/chat"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect server error: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.QA(ctx, &pb.Request{Question: "wewin"})
	if err != nil {
		log.Fatalf("Some error occurred when get data from server, %v", err)
	}
	for {
		answer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Some error occurred when get data: %v", err)
		}
		log.Println("Answer: ", answer.Answer)
	}
}
