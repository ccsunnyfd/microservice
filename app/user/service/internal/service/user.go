package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"microservice/app/user/service/internal/biz"

	v1 "microservice/api/user/service/v1"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/user")),
	}
}

func (s *UserService) GetUserByID(ctx context.Context, req *v1.GetUserByIDRequest) (*v1.GetUserByIDReply, error) {
	user, err := s.uc.GetUserByID(ctx, req.GetID())
	if err != nil {
		return nil, err
	}
	return &v1.GetUserByIDReply{
		UserInfo: &v1.UserInfo{
			ID:        user.ID,
			Username:  user.Username,
			Password:  user.Password,
			RealName:  user.RealName,
			Mobile:    user.Mobile,
			Email:     user.Email,
			IsTeacher: user.IsTeacher,
			Stars:     user.Stars,
		},
	}, nil
}
func (s *UserService) GetUserByName(ctx context.Context, req *v1.GetUserByNameRequest) (*v1.GetUserByNameReply, error) {
	user, err := s.uc.GetUserByName(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	return &v1.GetUserByNameReply{
		UserInfo: &v1.UserInfo{
			ID:        user.ID,
			Username:  user.Username,
			Password:  user.Password,
			RealName:  user.RealName,
			Mobile:    user.Mobile,
			Email:     user.Email,
			IsTeacher: user.IsTeacher,
			Stars:     user.Stars,
		},
	}, nil
}
func (s *UserService) RegisterUser(ctx context.Context, req *v1.RegisterUserRequest) (*v1.RegisterUserReply, error) {
	b, err := s.uc.RegisterUser(ctx, &biz.User{
		ID:       req.UserInfo.ID,
		Username: req.UserInfo.Username,
		Password: req.UserInfo.Password,
		RealName: req.UserInfo.RealName,
		Mobile:   req.UserInfo.Mobile,
		Email:    req.UserInfo.Email,
		IsTeacher: req.UserInfo.IsTeacher,
		Stars: req.UserInfo.Stars,
	})
	return &v1.RegisterUserReply{
		Success: b,
	}, err
}
