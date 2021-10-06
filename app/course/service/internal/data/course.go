package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"microservice/app/course/service/internal/biz"
)

var _ biz.CourseRepo = (*courseRepo)(nil)

type courseRepo struct {
	data *Data
	log  *log.Helper
}

func (r *courseRepo) ListCourses(ctx context.Context) ([]*biz.Course, error) {
	cs, err := r.data.db.Course.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Course, 0, len(cs))
	for _, c := range cs {
		rv = append(rv, &biz.Course{
			ID:          c.ID,
			Title:       c.Title,
			Description: c.Description,
			TeacherID:   c.TeacherID,
		})
	}
	return rv, nil
}

func NewCourseRepo(data *Data, logger log.Logger) biz.CourseRepo {
	return &courseRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/course")),
	}
}
