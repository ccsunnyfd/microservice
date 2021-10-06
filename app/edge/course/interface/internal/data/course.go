package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	courseV1 "microservice/api/course/service/v1"
	userV1 "microservice/api/user/service/v1"
	"microservice/app/edge/course/interface/internal/biz"
)

var _ biz.CourseRepo = (*courseRepo)(nil)

type courseRepo struct {
	data *Data
	log  *log.Helper
}

// ListCourses rpc call .
func (cr *courseRepo) ListCourses(ctx context.Context) ([]*biz.CourseInDB, error) {
	rv, err := cr.data.cc.ListCourse(ctx, &courseV1.ListCourseRequest{})
	if err != nil {
		return []*biz.CourseInDB{}, err
	}
	rvs := rv.Courses
	result := make([]*biz.CourseInDB, 0, len(rvs))
	for _, r := range rvs {
		result = append(result, &biz.CourseInDB{
			ID:          r.ID,
			Title:       r.Title,
			Description: r.Description,
			TeacherID:   r.TeacherId,
		})
	}
	return result, nil
}

// GetTeacher rpc call .
func (cr *courseRepo) GetTeacher(ctx context.Context, teacherID int64) (*biz.Teacher, error) {
	rv, err := cr.data.uc.GetUserByID(ctx, &userV1.GetUserByIDRequest{
		ID: teacherID,
	})
	if err != nil {
		return &biz.Teacher{}, err
	}
	return &biz.Teacher{
		ID:       rv.UserInfo.ID,
		Username: rv.UserInfo.Username,
		Password: rv.UserInfo.Password,
		RealName: rv.UserInfo.RealName,
		Mobile:   rv.UserInfo.Mobile,
		Email:    rv.UserInfo.Email,
		Stars:    rv.UserInfo.Stars,
	}, nil
}

func NewCourseRepo(data *Data, logger log.Logger) biz.CourseRepo {
	return &courseRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/edge/course")),
	}
}
