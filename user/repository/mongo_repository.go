package repository

import (
	"schoolsystem/learninghub-api/models"
	"schoolsystem/learninghub-api/user"

	"github.com/globalsign/mgo"
)

type mongoUserRepo struct {
	Session        *mgo.Session
	UserCollection *mgo.Collection
}

func NewMongoUserRepository(session *mgo.Session) user.UserRepository {

	return &mongoUserRepo{
		Session:        session,
		UserCollection: session.DB("school-system").C("users"),
	}
}

func (m *mongoUserRepo) Signup(u *models.User) (*models.User, error) {
	err := m.UserCollection.Insert(&u)

	if err != nil {
		return nil, err
	}

	return u, nil
}
