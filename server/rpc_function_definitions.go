package main

import (
	"context"
	proto "grpc-example/proto"
	"log"
)

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
