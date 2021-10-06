package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"microservice/app/message/service/internal/biz"

	pb "microservice/api/message/service/v1"
)

type EmailService struct {
	pb.UnimplementedEmailServer
	ec  *biz.EmailUseCase
	log *log.Helper
}

func NewEmailService(ec *biz.EmailUseCase, logger log.Logger) *EmailService {
	return &EmailService{
		ec:  ec,
		log: log.NewHelper(log.With(logger, "module", "service/email")),
	}
}

func (s *EmailService) SendEmail(ctx context.Context, req *pb.EmailRequest) (*pb.EmailReply, error) {
	s.log.Infof("rpc sendEmail received a request. Receiver: %s, Message: %s\n", req.Email, req.Message)
	err := s.ec.SendMessage(ctx, &biz.Email{
		To: req.Email,
		Message: req.Message,
	})
	if err != nil {
		return nil, err
	}
	return &pb.EmailReply{
		Success: true,
	}, nil
}
