package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	messageV1 "microservice/api/message/service/v1"
	userV1 "microservice/api/user/service/v1"
	"microservice/app/edge/user/interface/internal/conf"
	"strconv"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewUserRepo,
	NewData,
	NewEmailServiceClient,
	NewMobileServiceClient,
	NewUserServiceClient,
	NewRedisClient,
)

// Data .
type Data struct {
	log *log.Helper
	rdb *redis.Client
	ec  messageV1.EmailClient
	mc  messageV1.MobileClient
	uc  userV1.UserClient
}

// NewData .
func NewData(logger log.Logger, rdb *redis.Client, ec messageV1.EmailClient, mc messageV1.MobileClient, uc userV1.UserClient) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "edge-interface/user/data"))

	d := &Data{log: l, rdb: rdb, ec: ec, mc: mc, uc: uc}
	return d, func() {
		l.Info("message", "closing the redis resources")
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

func NewRedisClient(conf *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	return rdb
}

func NewEmailServiceClient(conf *conf.External_Message) messageV1.EmailClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(conf.GetAddr() + ":" + strconv.Itoa(int(conf.GetPort()))),
		//grpc.WithEndpoint("discovery:///beer.cart.service"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return messageV1.NewEmailClient(conn)
}

func NewMobileServiceClient(conf *conf.External_Message) messageV1.MobileClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(conf.GetAddr() + ":" + strconv.Itoa(int(conf.GetPort()))),
		//grpc.WithEndpoint("discovery:///beer.cart.service"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return messageV1.NewMobileClient(conn)
}

func NewUserServiceClient(conf *conf.External_User) userV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(conf.GetAddr() + ":" + strconv.Itoa(int(conf.GetPort()))),
		//grpc.WithEndpoint("discovery:///beer.cart.service"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return userV1.NewUserClient(conn)
}
