package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"microservice/app/user/service/internal/biz"
	"microservice/app/user/service/internal/data/ent/user"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) GetUserByID(ctx context.Context, ID int64) (*biz.User, error) {
	po, err := r.data.db.User.Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:        po.ID,
		Username:  po.Name,
		Password:  po.Password,
		RealName:  po.RealName,
		Mobile:    po.Mobile,
		Email:     po.Email,
		IsTeacher: po.IsTeacher,
		Stars:     po.Stars,
	}, nil
}

func (r *userRepo) GetUserByName(ctx context.Context, name string) (*biz.User, error) {
	po, err := r.data.db.User.Query().Where(user.NameEQ(name)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:        po.ID,
		Username:  po.Name,
		Password:  po.Password,
		RealName:  po.RealName,
		Mobile:    po.Mobile,
		Email:     po.Email,
		IsTeacher: po.IsTeacher,
		Stars:     po.Stars,
	}, nil
}

func (r *userRepo) RegisterUser(ctx context.Context, u *biz.User) (bool, error) {
	_, err := r.data.db.User.
		Create().
		SetName(u.Username).
		SetRealName(u.RealName).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetMobile(u.Mobile).
		SetIsTeacher(u.IsTeacher).
		SetStars(u.Stars).
		Save(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}
