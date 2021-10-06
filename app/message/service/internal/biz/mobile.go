package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Mobile struct {
	MobileNumber string
	Message      string
}

type MobileRepo interface {
	SendMessage(context.Context, *Mobile) error
}

type MobileUseCase struct {
	log   *log.Helper
}

func NewMobileUseCase(logger log.Logger) *MobileUseCase {
	return &MobileUseCase{
		log:   log.NewHelper(logger),
	}
}

func (uc *MobileUseCase) SendMessage(ctx context.Context, m *Mobile) error {
	uc.log.Infof("send mobile message, phone number: %s, message: %s\n", m.MobileNumber, m.Message)
	return nil
}
