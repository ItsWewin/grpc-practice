package main

import (
	"context"
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

	c := pb.NewChatClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.QA(ctx, &pb.Request{Question: "wewin"})
	if err != nil {
		log.Fatalf("Some error occurred when get data from server, %v", err)
	}
	log.Printf("Answer: %s", r.Answer)
}
