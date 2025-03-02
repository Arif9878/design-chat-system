package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &proto.HelloRequest{Name: "GoKit"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	fmt.Println(res.Message)
}
