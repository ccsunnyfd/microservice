package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	userEdgeV1 "microservice/api/edge/user/interface/v1"
	"microservice/app/edge/course/interface/internal/biz"
)

var _ biz.AuthRepo = (*authRepo)(nil)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "edge/course-auth")),
	}
}

func (r *authRepo) VerifyToken(ctx context.Context, token string) (*biz.User, error) {
	//_, err := r.data.cache.Get(token)
	rv, err := r.data.uec.Auth(context.Background(), &userEdgeV1.AuthReq{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	if !rv.Success {
		return nil, errors.New(403, "auth fail", "no valid token")
	}
	return &biz.User{
		ID:        rv.User.GetID(),
		Username:  rv.User.GetUsername(),
		Password:  rv.User.GetPassword(),
		RealName:  rv.User.GetRealName(),
		Mobile:    rv.User.GetMobile(),
		Email:     rv.User.GetEmail(),
		IsTeacher: rv.User.GetIsTeacher(),
		Stars:     rv.User.GetStars(),
	}, nil
}
