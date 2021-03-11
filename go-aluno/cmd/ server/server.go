package main

import (
	"log"
	"net"
	"github.com/felipeesc/grpc/pb"
	"github.com/felipeesc/grpc/services"
	"google.golang.org/grpc"
)
// para subir o servidor GRPC
// TODO go run cmd/server/server.go

func main()  {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Não conseguiu conectar: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService)

	if err := grpcServer.Server(lis); err != nil {
		log.Fatalf("Não conectou servidor: %v", err)
	}
}
