package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Course struct {
	ID          int64
	Title       string
	Description string
	TeacherID   int64
}

type CourseRepo interface {
	ListCourses(ctx context.Context) ([]*Course, error)
}

type CourseUseCase struct {
	repo CourseRepo
	log  *log.Helper
}

func (cc *CourseUseCase) ListCourses(ctx context.Context) ([]*Course, error) {
	return cc.repo.ListCourses(ctx)
}

func NewCourseUseCase(repo CourseRepo, logger log.Logger) *CourseUseCase {
	return &CourseUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "useCase/user")),
	}
}