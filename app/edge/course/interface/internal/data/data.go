package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	courseV1 "microservice/api/course/service/v1"
	userEdgeV1 "microservice/api/edge/user/interface/v1"
	userV1 "microservice/api/user/service/v1"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewCourseRepo,
	NewCourseServiceClient,
	NewUserEdgeServiceClient,
	NewUserServiceClient,
	NewAuthRepo,
)

// Data .
type Data struct {
	log   *log.Helper
	cc    courseV1.CourseClient
	uc    userV1.UserClient
	uec   userEdgeV1.UserEdgeInterfaceHTTPClient
}

// NewData .
func NewData(logger log.Logger, cc courseV1.CourseClient, uec userEdgeV1.UserEdgeInterfaceHTTPClient, uc userV1.UserClient) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "edge-interface/course/data"))

	d := &Data{log: l, cc: cc, uec: uec, uc: uc}
	//d := &Data{log: l, cc: cc, uec: uec, uc: uc, cache: cache}
	return d, nil
}

func NewCourseServiceClient() (courseClient courseV1.CourseClient, cleanup func()) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("course-service:10903"),
		//grpc.WithEndpoint("discovery:///beer.cart.service"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return courseV1.NewCourseClient(conn), func() {
		if err := conn.Close(); err != nil {
			_ = log.DefaultLogger.Log(log.LevelError, "close course service grpc conn error")
		}
	}
}

func NewUserEdgeServiceClient() (userEdgeClient userEdgeV1.UserEdgeInterfaceHTTPClient, cleanup func()) {
	conn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint("user-edge-service:8000"),
	)
	if err != nil {
		panic(err)
	}
	return userEdgeV1.NewUserEdgeInterfaceHTTPClient(conn), func() {
		if err := conn.Close(); err != nil {
			_ = log.DefaultLogger.Log(log.LevelError, "close userEdge service http conn error")
		}
	}
}

func NewUserServiceClient() (userClient userV1.UserClient, cleanup func()) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("user-service:10901"),
		//grpc.WithEndpoint("discovery:///beer.cart.service"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return userV1.NewUserClient(conn), func() {
		if err := conn.Close(); err != nil {
			_ = log.DefaultLogger.Log(log.LevelError, "close user service grpc conn error")
		}
	}
}
