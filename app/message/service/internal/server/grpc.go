package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v12 "microservice/api/message/service/v1"
	conf2 "microservice/app/message/service/internal/conf"
	service2 "microservice/app/message/service/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf2.Server, emailer *service2.EmailService, mobiler *service2.MobileService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v12.RegisterEmailServer(srv, emailer)
	v12.RegisterMobileServer(srv, mobiler)
	logger.Log(log.LevelInfo, "rpc in service")
	return srv
}