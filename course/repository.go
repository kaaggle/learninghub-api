package course

import "schoolsystem/learninghub-api/models"

type CourseRepository interface {
	GetCourse(string) (*models.Course, error)
	GetCourses() (*models.Courses, error)
}
