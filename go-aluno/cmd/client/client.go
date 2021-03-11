package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("não conseguiu conectar com gRPC server %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id: "0",
		Name: "Carvalho",
		Email: "felipe@carvalho.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("não conseguiu fazer requisição com gRPC %v", err)
	}

	fmt.Println(res)

//	TODO   go run cmd/client/client.go
}