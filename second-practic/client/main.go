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

	r, err := c.Echo(ctx, &pb.Request{Name: "wewin"})
	if err != nil {
		log.Fatalf("could nto greet: %v", err)
	}
	log.Printf("Greeting: %s", r.ReplyWorld)
}
