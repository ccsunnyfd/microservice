package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "microservice/api/edge/course/interface/v1"
	"microservice/app/edge/course/interface/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCourseEdgeInterfaceService)

type CourseInterface struct {
	v1.UnimplementedCourseEdgeInterfaceServer
	uc  *biz.CourseUseCase
	log *log.Helper
}

func NewCourseEdgeInterfaceService(uc *biz.CourseUseCase, logger log.Logger) *CourseInterface {
	return &CourseInterface{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/edge/course")),
	}
}