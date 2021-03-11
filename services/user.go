package services

import (
	"context"
	"github.com/felipeesc/grpc/pb"
	"time"
)

type UserServiceServer interface {
	AddUser(context.Context, *User) (*User, error)
	mustEmbedUnimplementedUserServiceServer()
	AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)
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

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {

	stream.Send(*pb.UserResultStream{
		Status: "Iniciando",
		User: &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(*pb.UserResultStream{
		Status: "Inserindo",
		User: &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(*pb.UserResultStream{
		Status: "Usuário inserindo",
		User: &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		},
	})
	time.Sleep(time.Second * 3)

	stream.Send(*pb.UserResultStream{
		Status: "Completo",
		User: &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}

// instala o evans .. para fazer primeira chamada...
// TODO evans -r repl --host localhost --port 50051