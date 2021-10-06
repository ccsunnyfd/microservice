package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"log"
	"microservice/api/message/service/v1"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint("127.0.0.1:9000"),
	)
	if err != nil {
		log.Fatalf("rpc connect error: %v\n", err)
	}

	eClient := v1.NewEmailClient(conn)
	reply, err := eClient.SendEmail(ctx, &v1.EmailRequest{
		Email:   "chenzheyuan@rorke.com.cn",
		Message: "test007",
	})
	if err != nil {
		log.Fatalf("sendEmail error: %v\n", err)
	}
	log.Printf("sendEmail succeeded: %v\n", reply.Success)
}
