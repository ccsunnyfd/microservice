package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"log"
	"microservice/api/course/service/v1"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint("127.0.0.1:10903"),
	)
	if err != nil {
		log.Fatalf("rpc connect error: %v\n", err)
	}

	cClient := v1.NewCourseClient(conn)
	reply1, err := cClient.ListCourse(ctx, &v1.ListCourseRequest{})
	if err != nil {
		log.Fatalf("list courss error: %v\n", err)
	}
	log.Printf("list courss succeeded: %v\n", reply1.GetCourses())
}
