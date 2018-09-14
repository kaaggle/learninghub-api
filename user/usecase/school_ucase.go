package usecase

import (
	"schoolsystem/learninghub-api/models"
	"schoolsystem/learninghub-api/user"
)

type userUsecase struct {
	userRepo user.UserRepository
}

func NewUserUsecase(userRepo user.UserRepository) user.UserRepository {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (a *userUsecase) Signup(u *models.User) (*models.User, error) {
	c, err := a.userRepo.Signup(u)

	if err != nil {
		return nil, err
	}

	return c, nil
}
