package user

import "schoolsystem/learninghub-api/models"

type UserRepository interface {
	Signup(user *models.User) (*models.User, error)
}
