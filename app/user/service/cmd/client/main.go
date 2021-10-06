package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"log"
	"microservice/api/user/service/v1"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint("127.0.0.1:9001"),
	)
	if err != nil {
		log.Fatalf("rpc connect error: %v\n", err)
	}

	uClient := v1.NewUserClient(conn)
	reply1, err := uClient.RegisterUser(ctx, &v1.RegisterUserRequest{
		UserInfo: &v1.UserInfo{
			Username: "David",
			Password: "scot@123.com",
			RealName: "chen",
			Mobile:   "13784949884",
			Email:    "03949@fglgkl.com",
		},
	})
	if err != nil {
		log.Fatalf("register user error: %v\n", err)
	}
	log.Printf("register user succeeded: %v\n", reply1.Success)

	reply2, err := uClient.GetUserByName(ctx, &v1.GetUserByNameRequest{
		Name: "David",
	})
	if err != nil {
		log.Fatalf("get user by name error: %v\n", err)
	}
	log.Printf("get user by name succeeded: %v\n", *reply2.GetUserInfo())

	reply3, err := uClient.GetUserByID(ctx, &v1.GetUserByIDRequest{
		ID: reply2.GetUserInfo().GetID(),
	})
	if err != nil {
		log.Fatalf("get user by ID error: %v\n", err)
	}
	log.Printf("get user by ID succeeded: %v\n", *reply3.GetUserInfo())
}
