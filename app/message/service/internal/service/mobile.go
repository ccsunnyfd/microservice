package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"microservice/app/message/service/internal/biz"

	pb "microservice/api/message/service/v1"
)

type MobileService struct {
	pb.UnimplementedMobileServer
	mc  *biz.MobileUseCase
	log *log.Helper
}

func NewMobileService(mc *biz.MobileUseCase, logger log.Logger) *MobileService {
	return &MobileService{
		mc:  mc,
		log: log.NewHelper(log.With(logger, "module", "service/mobile")),
	}
}

func (s *MobileService) SendMobile(ctx context.Context, req *pb.MobileRequest) (*pb.MobileReply, error) {
	err := s.mc.SendMessage(ctx, &biz.Mobile{
		MobileNumber: req.Mobile,
		Message:      req.Message,
	})
	if err != nil {
		return nil, err
	}
	return &pb.MobileReply{
		Success: true,
	}, nil
}
