package usecase

import (
	"schoolsystem/learninghub-api/course"
	"schoolsystem/learninghub-api/models"
)

type courseUsecase struct {
	courseRepo course.CourseRepository
}

func NewCourseUsecase(courseRepo course.CourseRepository) course.CourseRepository {
	return &courseUsecase{
		courseRepo: courseRepo,
	}
}

func (a *courseUsecase) AddCourse(course *models.Course) (*models.Course, error) {
	c, err := a.courseRepo.AddCourse(course)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (a *courseUsecase) DeleteCourse(id string) (error) {
	err := a.courseRepo.DeleteCourse(id)

	if err != nil {
		return err
	}

	return nil
}


func (a *courseUsecase) GetCourses() (*models.Courses, error) {
	c, err := a.courseRepo.GetCourses()

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (a *courseUsecase) GetCourse(id string) (*models.Course, error) {
	c, err := a.courseRepo.GetCourse(id)

	if err != nil {
		return nil, err
	}

	return c, nil
}
