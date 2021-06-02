# Project Title

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Contributing](../CONTRIBUTING.md)

## About <a name = "about"></a>

There are few steps you need to follow to create a rest API with grpc in golang
Here I am attaching protoc compiler for linux, and googleapi library so that you can directly dive into coding 

## Getting Started <a name = "getting_started"></a>

1. Create three folder client, server and client
2. Create a test.proto file inside proto folder, and paste the content of the test.proto file
3. Now create a go file main.go in server folder.
4. Implement all the unimplemented functions in the main.go file, 
   and then inside main function create a tcp listener and a grpc server, 
   register GreetServer with this grpc server. 
   And make grpc server to listen on the tcp listener.
5. There are many ways of creating grpc client, I have shown one of them inside GRPCClient() function
   StartREST() function will start a rest server, 
   SimpleRPC() function can be used to call normal RPC functions.


## Usage <a name = "usage"></a>

You can use this to get an oversight of building a rest API with grpc in golang
