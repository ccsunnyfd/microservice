//+build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	biz "microservice/app/edge/user/interface/internal/biz"
	conf "microservice/app/edge/user/interface/internal/conf"
	data "microservice/app/edge/user/interface/internal/data"
	server "microservice/app/edge/user/interface/internal/server"
	service "microservice/app/edge/user/interface/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.External_User, *conf.External_Message, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
