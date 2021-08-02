package main

import (
	"context"
	"log"

	pb "github.com/vitor-mariano/first-go-grpc-exercise/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("app:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	client := pb.NewGreeterClient(conn)

	req := &pb.HelloRequest{
		Name: "Vitor Mariano",
	}

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("Did not receive response: %v", err)
	}

	log.Printf("Received: %s", res.GetMessage())
}
