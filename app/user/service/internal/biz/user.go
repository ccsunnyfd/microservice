package biz

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID        int64
	Username  string
	Password  string
	RealName  string
	Mobile    string
	Email     string
	IsTeacher bool
	Stars     int32
}

type UserRepo interface {
	GetUserByID(ctx context.Context, ID int64) (*User, error)
	GetUserByName(ctx context.Context, name string) (*User, error)
	RegisterUser(ctx context.Context, u *User) (bool, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, ID int64) (*User, error) {
	return uc.repo.GetUserByID(ctx, ID)
}

func (uc *UserUseCase) GetUserByName(ctx context.Context, name string) (*User, error) {
	return uc.repo.GetUserByName(ctx, name)
}

func (uc *UserUseCase) RegisterUser(ctx context.Context, u *User) (bool, error) {
	newUser := &User{
		Username: u.Username,
		Password: genMD5([]byte(u.Password)),
		RealName: u.RealName,
		Mobile:   u.Mobile,
		Email:    u.Email,
		IsTeacher: u.IsTeacher,
		Stars: u.Stars,
	}
	return uc.repo.RegisterUser(ctx, newUser)
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "useCase/user"))}
}

func genMD5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}
