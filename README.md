# Simple rest API in golang with grpc

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)

## About <a name = "about"></a>

There are few steps you need to follow to create a rest API with grpc in golang
Here I am attaching protoc compiler for linux, and googleapi library so that you can directly dive into coding 

## Getting Started <a name = "getting_started"></a>

1. Create three folder client, server and client
2. Create a test.proto file inside proto folder, and paste the content of the test.proto file
3. Create a main.go file insede server folder. 
   In this file define all the rpc functions.
   In the main.go write main function inside main function create a tcp listener and a grpc server, register GreeterServer with this grpc server. 
   And make grpc server to listen on the tcp listener.
4. There are many ways of creating grpc client, I have shown one of them inside GRPCClient() function
   StartREST() function will start a rest server, 
   SimpleRPC() function can be used to call normal RPC functions.

## 🔧 Running the tests

For testing I have used Thunder Client extension in VScode.

### Break down into end to end tests

First install the Thunder Client in vscode.

```
Then import the json collection 'thunder-collection_go-rest-server.json'
Then you can run a single request once at a time, or as collection.
```
