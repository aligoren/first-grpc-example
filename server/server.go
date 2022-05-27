package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)
import pb "grpc_project/gen/proto"

type helloApiServer struct {
	pb.UnimplementedHelloApiServer
}

func (h *helloApiServer) WriteHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("You said! %s", request.Message),
	}, nil
}

func (h helloApiServer) GetDate(ctx context.Context, empty *emptypb.Empty) (*pb.DateResponse, error) {
	return &pb.DateResponse{
		Date: timestamppb.Now(),
	}, nil
}

func main() {

	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterHelloApiServer(grpcServer, &helloApiServer{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Error while serving gRPC server: %v", err)
	}
}
