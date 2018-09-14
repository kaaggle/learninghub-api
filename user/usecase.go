package user

import "schoolsystem/learninghub-api/models"

type UserUsecase interface {
	Signup(user *models.User) (*models.User, error)
}
