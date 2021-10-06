package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"microservice/app/edge/user/interface/internal/pkg/code"
	"time"
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
	GetUserByName(ctx context.Context, username string) (*User, error)
	SetToken(ctx context.Context, token string, user *User, expire time.Duration) error
	GetToken(ctx context.Context, token string) (*User, error)

	CreateUser(ctx context.Context, u *User) error
	SendMobile(ctx context.Context, mobile string, message string) error
	SendEmail(ctx context.Context, email string, message string) error
	GetVerifyCode(ctx context.Context, codeKey string) (string, error)
	SetVerifyCode(ctx context.Context, codeKey string, codeValue string, expire time.Duration) error
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func (uc *UserUseCase) Login(ctx context.Context, user *User) (string, error) {
	// get user info from db
	dbUser, err := uc.repo.GetUserByName(ctx, user.Username)
	if err != nil {
		return "", err
	}
	// compare send pass with db pass
	if code.GenMD5([]byte(user.Password)) != dbUser.Password {
		return "", errors.New(403, "login fail", "username or password error")
	}
	// generate token
	token := code.GenToken()
	// save token to cache
	err = uc.repo.SetToken(ctx, token, dbUser, 3600*time.Second)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uc *UserUseCase) GenerateVerifyCode(ctx context.Context, mobile, email string) error {
	if mobile == "" && email == "" {
		return errors.New(403, "register fail", "must provide mail or mobile")
	}
	verifyCode := string(code.RandCode("0123456789abcdfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 6))
	message := "Verify code is: " + verifyCode
	if mobile != "" {
		err := uc.repo.SendMobile(ctx, mobile, message)
		if err != nil {
			return err
		}
		err = uc.repo.SetVerifyCode(ctx, mobile, verifyCode, 300*time.Second)
		if err != nil {
			return err
		}
		return nil
	}
	err := uc.repo.SendEmail(ctx, email, message)
	if err != nil {
		return err
	}
	err = uc.repo.SetVerifyCode(ctx, email, verifyCode, 300*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) Register(ctx context.Context, u *User, sendCode string) error {
	// check verifyCode
	mobile, email := u.Mobile, u.Email
	if mobile == "" && email == "" {
		return errors.New(403, "register fail", "must provide mail or mobile")
	}
	var codeKey string
	if mobile != "" {
		codeKey = mobile
	} else {
		codeKey = email
	}
	dbCode, err := uc.repo.GetVerifyCode(ctx, codeKey)
	if err != nil {
		return err
	}
	if dbCode != sendCode {
		return errors.New(403, "register fail", "verify code incorrect")
	}
	// pass, create user
	err = uc.repo.CreateUser(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) Auth(ctx context.Context, token string) (*User, error) {
	return uc.repo.GetToken(ctx, token)
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "useCase/edge/user"))}
}
