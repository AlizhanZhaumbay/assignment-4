package main

import (
	"assignment-4/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := protos.NewUserServiceClient(conn)

	addUserResp, err := client.AddUser(context.Background(), &protos.UserCreateRequest{Name: "Alice", Email: "alice@example.com"})
	if err != nil {
		log.Fatalf("AddUser call failed: %v", err)
	}
	fmt.Printf("User ID: %d\n", addUserResp.Id)

	getUserResp, err := client.GetUser(context.Background(), &protos.UserGetRequest{Id: 1})
	if err != nil {
		log.Fatalf("GetUser call failed: %v", err)
	}
	fmt.Printf("User: %+v\n", getUserResp)

	listUsersResp, err := client.ListUsers(context.Background(), &protos.UserGetAllRequest{})
	if err != nil {
		log.Fatalf("ListUsers call failed: %v", err)
	}
	fmt.Println("Users:")
	for _, user := range listUsersResp.Users {
		fmt.Printf(" - %+v\n", user)
	}
}
