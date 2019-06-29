package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpc-learn/first-practic/helloworld"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect server error: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %V", err)
	}

	log.Printf("Greeting: %s", r.Message)
}
