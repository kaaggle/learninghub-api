package repository

import (
	"schoolsystem/learninghub-api/course"
	"schoolsystem/learninghub-api/models"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type mongoCourseRepo struct {
	Session          *mgo.Session
	CourseCollection *mgo.Collection
}

func NewMongoCourseRepository(session *mgo.Session) course.CourseRepository {

	return &mongoCourseRepo{
		Session:          session,
		CourseCollection: session.DB("school-system").C("courses"),
	}
}

func (m *mongoCourseRepo) GetCourses() (*models.Courses, error) {
	c := models.Courses{}
	err := m.CourseCollection.Find(bson.M{}).All(&c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (m *mongoCourseRepo) GetCourse(id string) (*models.Course, error) {
	c := models.Course{}
	err := m.CourseCollection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
