package course

import "schoolsystem/learninghub-api/models"

type CourseUsecase interface {
	GetCourse(string) (*models.Course, error)
	GetCourses() (*models.Courses, error)
}
