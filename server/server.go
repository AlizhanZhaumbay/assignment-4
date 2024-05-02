package main

import (
	"assignment-4/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserServiceServerImpl struct {
	protos.UnimplementedUserServiceServer
}

func (s *UserServiceServerImpl) AddUser(ctx context.Context, req *protos.UserCreateRequest) (*protos.UserIdResponse, error) {
	fmt.Printf("Received request to add user with name: %s, email: %s\n", req.Name, req.Email)

	userID := int32(123)

	return &protos.UserIdResponse{Id: userID}, nil
}

func (s *UserServiceServerImpl) GetUser(ctx context.Context, req *protos.UserGetRequest) (*protos.User, error) {
	return &protos.User{
		Id:    req.Id,
		Name:  "John Doe",
		Email: "john@example.com",
	}, nil
}

func (s *UserServiceServerImpl) ListUsers(ctx context.Context, req *protos.UserGetAllRequest) (*protos.UserGetAllResponse, error) {

	users := []*protos.User{
		{Id: 1, Name: "Alice", Email: "alice@example.com"},
		{Id: 2, Name: "Bob", Email: "bob@example.com"},
		{Id: 3, Name: "Charlie", Email: "charlie@example.com"},
	}

	return &protos.UserGetAllResponse{Users: users}, nil
}

func main() {
	server := grpc.NewServer()

	protos.RegisterUserServiceServer(server, &UserServiceServerImpl{})

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Server started, listening on port 8080")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
