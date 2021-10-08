package service

import (
	"context"
	"microservice/api/edge/user/interface/v1"
	"microservice/app/edge/user/interface/internal/biz"
)

func (s *UserInterface) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	err := s.uc.Register(ctx, &biz.User{
		ID:        0,
		Username:  req.User.Username,
		Password:  req.User.Password,
		RealName:  req.User.RealName,
		Mobile:    req.User.GetMobile(),
		Email:     req.User.GetEmail(),
		IsTeacher: req.User.IsTeacher,
		Stars:     req.User.Stars,
	}, req.VerifyCode)
	if err != nil {
		s.log.Errorf("register fail: %v", err)
		return &v1.RegisterReply{
			Success: false,
		}, v1.ErrorRegisterFail("%s register fail", req.User.Username)
	}
	return &v1.RegisterReply{
		Success: true,
	}, nil
}
func (s *UserInterface) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	token, err := s.uc.Login(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		s.log.Errorf("login fail: %v", err)
		return &v1.LoginReply{
		}, v1.ErrorLoginFail("%s login failed", req.Username)
	}
	return &v1.LoginReply{
		Token: token,
	}, nil
}
func (s *UserInterface) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	return &v1.LogoutReply{}, nil
}
func (s *UserInterface) Auth(ctx context.Context, req *v1.AuthReq) (*v1.AuthReply, error) {
	bizUser, err := s.uc.Auth(ctx, req.Token)
	if err != nil {
		s.log.Errorf("auth fail: %v", err)
		return &v1.AuthReply{
			Success: false,
		}, v1.ErrorAuthFail("%s auth fail", req.Token)
	}
	return &v1.AuthReply{
		Success: true,
		User: &v1.DBUser{
			Username:  bizUser.Username,
			Password:  bizUser.Password,
			Email:     bizUser.Email,
			Mobile:    bizUser.Mobile,
			RealName:  bizUser.RealName,
			IsTeacher: bizUser.IsTeacher,
			Stars:     bizUser.Stars,
		},
	}, nil
}

func (s *UserInterface) GenVerifyCode(ctx context.Context, req *v1.GenVerifyCodeReq) (*v1.GenVerifyCodeReply, error) {
	return &v1.GenVerifyCodeReply{}, s.uc.GenerateVerifyCode(ctx, req.GetMobile(), req.GetEmail())
}
