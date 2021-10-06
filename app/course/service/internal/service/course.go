package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"microservice/app/course/service/internal/biz"

	pb "microservice/api/course/service/v1"
)

type CourseService struct {
	pb.UnimplementedCourseServer
	uc  *biz.CourseUseCase
	log *log.Helper
}

func NewCourseService(useCase *biz.CourseUseCase, logger log.Logger) *CourseService {
	return &CourseService{
		uc:  useCase,
		log: log.NewHelper(log.With(logger, "module", "service/course")),
	}
}

func (s *CourseService) ListCourse(ctx context.Context, req *pb.ListCourseRequest) (*pb.ListCourseReply, error) {
	rv, err := s.uc.ListCourses(ctx)

	if err != nil {
		return &pb.ListCourseReply{}, err
	}

	rs := make([]*pb.ListCourseReply_Course, 0, len(rv))
	for _, x := range rv {
		rs = append(rs, &pb.ListCourseReply_Course{
			ID:          x.ID,
			Title:       x.Title,
			Description: x.Description,
			TeacherId:   x.TeacherID,
		})
	}
	return &pb.ListCourseReply{
		Courses: rs,
	}, nil
}
