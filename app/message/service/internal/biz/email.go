package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sync"
)

type EmailSender struct {
	DialerPool *sync.Pool
	From       string
}

type Email struct {
	To      string
	Message string
}

type EmailRepo interface {
	SendMessage(context.Context, *Email) error
}

type EmailUseCase struct {
	repo EmailRepo
	log  *log.Helper
}

func NewEmailUseCase(repo EmailRepo, logger log.Logger) *EmailUseCase {
	return &EmailUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "useCase/email")),
	}
}

func (uc *EmailUseCase) SendMessage(ctx context.Context, m *Email) error {
	uc.log.Infof("send email message, email receiver: %s, message: %s\n", m.To, m.Message)
	return uc.repo.SendMessage(ctx, m)
}
