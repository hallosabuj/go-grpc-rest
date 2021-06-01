package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "grpc-example/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	proto.UnimplementedGreeterServer
}

// SayHello()
func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	// pass
	log.Println("Received :", in.GetName())
	return &proto.HelloReply{Message: "Hello Sabuj"}, nil
}

// EchoGET()
func (s *server) EchoGET(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello User"}, nil
}

// EchoPOST()
func (s *server) EchoPOST(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello " + in.Name}, nil
}

// EchoName()
func (s *server) EchoName(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	fmt.Println("Hello server")
	// Creating tcp listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen :", err)
	}
	// Creating a grpc server
	s := grpc.NewServer()
	// Registering the GreetServer with the grpc server
	proto.RegisterGreeterServer(s, &server{})
	// Starting the grpc server and allowed to listen on tcp listener
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve :", err)
	}
}
