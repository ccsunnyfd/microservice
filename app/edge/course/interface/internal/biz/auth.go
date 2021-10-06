package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID       int64
	Username string
	Password string
	RealName string
	Mobile   string
	Email    string
	IsTeacher bool
	Stars    int32
}

type AuthRepo interface {
	VerifyToken(ctx context.Context, token string) (*User, error)
}

type AuthUseCase struct {
	repo AuthRepo
	log  *log.Helper
}

func (ac *AuthUseCase) Authenticate(ctx context.Context, token string) error {
	_, err := ac.repo.VerifyToken(ctx, token)
	return err
}

func NewAuthUseCase(repo AuthRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "useCase/edge/course-auth"))}
}