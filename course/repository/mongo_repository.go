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
		CourseCollection: session.DB("school-system").C("lh-courses"),
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

func (m *mongoCourseRepo) AddCourse(c *models.Course) (*models.Course, error) {
	err := m.CourseCollection.Insert(&c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (m *mongoCourseRepo) DeleteCourse(id string) (error) {
	err := m.CourseCollection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil {
		return err
	}

	return nil
}


