package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	messageV1 "microservice/api/message/service/v1"
	"microservice/api/user/service/v1"
	"microservice/app/edge/user/interface/internal/biz"
	"time"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// CreateUser rpc .
func (ur *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	reply, err := ur.data.uc.RegisterUser(ctx, &v1.RegisterUserRequest{
		UserInfo: &v1.UserInfo{
			Username:  user.Username,
			Password:  user.Password,
			RealName:  user.RealName,
			Mobile:    user.Mobile,
			Email:     user.Email,
			IsTeacher: user.IsTeacher,
			Stars:     user.Stars,
		},
	})
	if err != nil {
		return err
	}
	if !reply.GetSuccess() {
		return errors.New(503, "rpc fail", "rpc create user fail")
	}
	return nil
}

// SendMobile rpc .
func (ur *userRepo) SendMobile(ctx context.Context, mobile string, message string) error {
	reply, err := ur.data.mc.SendMobile(ctx, &messageV1.MobileRequest{
		Mobile:  mobile,
		Message: message,
	})
	if err != nil {
		return err
	}
	if !reply.GetSuccess() {
		return errors.New(503, "rpc fail", "rpc send mobile fail")
	}
	return nil
}

// SendEmail rpc .
func (ur *userRepo) SendEmail(ctx context.Context, email string, message string) error {
	reply, err := ur.data.ec.SendEmail(ctx, &messageV1.EmailRequest{
		Email:   email,
		Message: message,
	})
	if err != nil {
		return err
	}
	if !reply.GetSuccess() {
		return errors.New(503, "rpc fail", "rpc send email fail")
	}
	return nil
}

// GetVerifyCode redis .
func (ur *userRepo) GetVerifyCode(ctx context.Context, codeKey string) (string, error) {
	cacheCode, err := ur.data.rdb.Get(ctx, codeKey).Result()
	if err == redis.Nil {
		return "", errors.New(403, "verify code fail", "codekey does not exist")
	} else if err != nil {
		return "", err
	}
	return cacheCode, nil
}

// SetVerifyCode redis .
func (ur *userRepo) SetVerifyCode(ctx context.Context, codeKey string, codeValue string, expire time.Duration) error {
	return ur.data.rdb.Set(ctx, codeKey, codeValue, expire).Err()
}

// GetUserByName rpc .
func (ur *userRepo) GetUserByName(ctx context.Context, username string) (*biz.User, error) {
	reply, err := ur.data.uc.GetUserByName(ctx, &v1.GetUserByNameRequest{
		Name: username,
	})
	if err != nil {
		return nil, err
	}
	userInfo := reply.GetUserInfo()
	return &biz.User{
		ID:        userInfo.ID,
		Username:  userInfo.Username,
		Password:  userInfo.Password,
		RealName:  userInfo.RealName,
		Mobile:    userInfo.Mobile,
		Email:     userInfo.Email,
		IsTeacher: userInfo.IsTeacher,
		Stars:     userInfo.Stars,
	}, nil
}

// SetToken redis .
func (ur *userRepo) SetToken(ctx context.Context, token string, user *biz.User, expire time.Duration) error {
	data, _ := json.Marshal(user)
	return ur.data.rdb.Set(ctx, token, data, expire).Err()
}

// GetToken redis .
func (ur *userRepo) GetToken(ctx context.Context, token string) (*biz.User, error) {
	get, err := ur.data.rdb.Get(ctx, token).Bytes()
	if err != nil {
		return nil, err
	}
	cacheUser := &biz.User{}
	err = json.Unmarshal(get, cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/edge/user")),
	}
}
