package main

import (
	"context"
	"log"
	"time"

	pb "grpc-learn/second-practic/chat"

	"google.golang.org/grpc"
)

func unaryRPCs(client pb.ChatClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.QA(ctx, &pb.Request{Question: "question1"})
	if err != nil {
		log.Fatalf("Some error occurred when get data from server, %v", err)
	}
	log.Printf("Answer: %s", res.Answer)
}

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect server error: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatClient(conn)

	unaryRPCs(client)
}
