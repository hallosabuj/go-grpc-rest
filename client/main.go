package main

import (
	"context"
	"fmt"
	"grpc-example/proto"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

// This is only for creating a grpc client
func GRPCClient(port int) (*grpc.ClientConn, error) {
	return grpc.DialContext(
		context.Background(),
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
}

// This function will be used for calling the API
func StartREST(conn *grpc.ClientConn) {
	/////////////////////////////////////////////////////////////////////
	// Now connecting to the google api from using the REST server
	/////////////////////////////////////////////////////////////////////
	mux := runtime.NewServeMux()
	err := proto.RegisterGreeterHandler(context.Background(), mux, conn)
	if err != nil {
		fmt.Println("Failed initializing grpc gateway handler, error: ", err)
	}

	handlers := http.NewServeMux()
	handlers.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if strings.HasPrefix(request.URL.Path, "/api") {
			mux.ServeHTTP(writer, request)
			return
		}
	})

	log.Println(http.ListenAndServe(":8080", handlers))
}

// This function is user for calling rpc functions directly
func SimpleRPC(conn *grpc.ClientConn) {
	c := proto.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: name})
	if err != nil {
		fmt.Println("Could not greet :", err)
	}
	fmt.Println("Greeting :", r.GetMessage())
}

func main() {
	// Creating a grpc client by using grpc.DialConetxt()
	conn, err := GRPCClient(1234)
	if err != nil {
		fmt.Println("Did not connect :", err)
	}
	defer conn.Close()

	// This function can be used for calling simple rpc functions available in the server side.
	SimpleRPC(conn)

	// This can be user for making API calls
	StartREST(conn)
}
