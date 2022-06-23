package main

import (
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
