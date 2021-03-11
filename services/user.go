package services

import (
	"context"
	"github.com/felipeesc/grpc/pb"
)

type UserServiceServer interface {
	AddUser(context.Context, *User) (*User, error)
	mustEmbedUnimplementedUserServiceServer()
}

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	// simulando inserção no database
	return &pb.User {
		Id: "123",
		Name: req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

// instala o evans .. para fazer primeira chamada...
// TODO evans -r repl --host localhost --port 50051