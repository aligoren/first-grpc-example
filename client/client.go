package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "grpc_project/gen/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Couldn't connect to gRPC server %v", err)
	}

	client := pb.NewHelloApiClient(conn)

	resp, _ := client.WriteHello(context.Background(), &pb.HelloRequest{
		Message: "Hello world!",
	})

	fmt.Printf("Response Message: ----> %s\n", resp.Message)

	timeStampResponse, _ := client.GetDate(context.Background(), &emptypb.Empty{})

	fmt.Printf("Response timestamp ----> %s", timeStampResponse.Date.AsTime().String())
}
