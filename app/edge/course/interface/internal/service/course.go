package service

import (
	"context"
	"microservice/api/edge/course/interface/v1"
)

func (s *CourseInterface) ListCourse(ctx context.Context, req *v1.ListCourseRequest) (*v1.ListCourseReply, error) {
	courses, err := s.uc.List(ctx)
	if err != nil {
		return &v1.ListCourseReply{}, v1.ErrorListCoursesFail("list courses fail")
	}

	result := make([]*v1.ListCourseReply_Course, 0, len(courses))
	for _, c := range courses {
		v1C := &v1.ListCourseReply_Course{
			ID:          c.ID,
			Title:       c.Title,
			Description: c.Description,
		}
		if c.Teacher != nil {
			v1C.Teacher = &v1.ListCourseReply_Course_Teacher{
				ID:       c.Teacher.ID,
				Username: c.Teacher.Username,
				RealName: c.Teacher.RealName,
				Mobile:   c.Teacher.Mobile,
				Email:    c.Teacher.Email,
				Stars:    c.Teacher.Stars,
			}
		}
		result = append(result, v1C)
	}
	return &v1.ListCourseReply{
		Courses: result,
	}, nil
}