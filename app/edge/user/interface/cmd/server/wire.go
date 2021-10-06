//+build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	biz "microservice/app/edge/course/interface/internal/biz"
	conf "microservice/app/edge/course/interface/internal/conf"
	data "microservice/app/edge/course/interface/internal/data"
	server "microservice/app/edge/course/interface/internal/server"
	service "microservice/app/edge/course/interface/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
