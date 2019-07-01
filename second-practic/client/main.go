package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "grpc-learn/second-practic/chat"

	"google.golang.org/grpc"
)

var questions = [...]*pb.Request{
	&pb.Request{Question: "question1"},
	&pb.Request{Question: "question2"},
	&pb.Request{Question: "question3"},
}

func delBidirectionalStream(client pb.ChatClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.QA(ctx)
	if err != nil {
		log.Fatalf("Some error occurred when get data from server, %v", err)
	}

	// 无缓冲 Channel
	// 阻塞主goroutine，防止提前退出
	wait := make(chan struct{})
	// 使用 goroutine 处理服务端响应
	go func() {
		for {
			answer, err := stream.Recv()
			if err == io.EOF {
				close(wait)
				return
			}
			if err != nil {
				return
			}
			log.Println("Answer: ", answer.Answer)
		}
	}()

	// 发送请求到服务端
	for _, question := range questions {
		if err := stream.Send(question); err != nil {
			log.Fatalf("Some error occurred when sent data to server, error;l %v", err)
		}
	}
	stream.CloseSend()
	<-wait
}

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect server error: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatClient(conn)

	delBidirectionalStream(client)
}
