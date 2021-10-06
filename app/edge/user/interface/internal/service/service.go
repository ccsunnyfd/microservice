package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "microservice/api/edge/user/interface/v1"
	"microservice/app/edge/user/interface/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserEdgeInterfaceService)

type UserInterface struct {
	v1.UnimplementedUserEdgeInterfaceServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserEdgeInterfaceService(uc *biz.UserUseCase, logger log.Logger) *UserInterface {
	return &UserInterface{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/edge/user")),
	}
}
