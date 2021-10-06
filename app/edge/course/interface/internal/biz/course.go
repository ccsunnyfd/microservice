package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Teacher struct {
	ID       int64
	Username string
	Password string
	RealName string
	Mobile   string
	Email    string
	Stars    int32
}

type CourseInDB struct {
	ID          int64
	Title       string
	Description string
	TeacherID   int64
}

type Course struct {
	ID          int64
	Title       string
	Description string
	Teacher     *Teacher
}

type CourseRepo interface {
	ListCourses(ctx context.Context) ([]*CourseInDB, error)
	GetTeacher(ctx context.Context, teacherID int64) (*Teacher, error)
}

type CourseUseCase struct {
	repo CourseRepo
	log  *log.Helper
}

func (cc *CourseUseCase) List(ctx context.Context) ([]*Course, error) {
	// get course list from db
	dbCourses, err := cc.repo.ListCourses(ctx)
	if err != nil {
		return []*Course{}, err
	}
	result := make([]*Course, 0, len(dbCourses))
	// get each teacher from user rpc call
	for _, course := range dbCourses {
		c := &Course{
			ID:          course.ID,
			Title:       course.Title,
			Description: course.Description,
		}
		teacherID := course.TeacherID
		t, err := cc.repo.GetTeacher(ctx, teacherID)
		if err != nil {
			cc.log.Warnf("teacher info of course %s is missing", course.Title)
		} else {
			c.Teacher = t
		}
		result = append(result, c)
	}
	// return result
	return result, nil
}

func NewCourseUseCase(repo CourseRepo, logger log.Logger) *CourseUseCase {
	return &CourseUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "useCase/edge/course"))}
}
