package course

import (
	"encoding/json"
	"net/http"
	"schoolsystem/learninghub-api/models"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/chi"
)

type HttpCourseHandler struct {
	CourseUsecase CourseUsecase
}

func (h *HttpCourseHandler) AddCourse(w http.ResponseWriter, r *http.Request) {
		c := models.Course{}

		json.NewDecoder(r.Body).Decode(&c)
		c.CreatedAt = time.Now()
		err := c.Validate()

		if err != nil {
			json.NewEncoder(w).Encode(models.Response{
				Error:   true,
				Message: err.Error(),
				Data:    err,
			})
			return
		}

		attendance, err := h.CourseUsecase.AddCourse(&c)
		if err != nil {
			json.NewEncoder(w).Encode(models.Response{
				Error:   true,
				Message: err.Error(),
				Data:    err,
			})
			return
		}

		json.NewEncoder(w).Encode(models.Response{
			Error:   false,
			Message: "Added course.",
			Data:    &attendance,
		})
}

func (h *HttpCourseHandler) GetCourses(w http.ResponseWriter, r *http.Request) {

	c, err := h.CourseUsecase.GetCourses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "Retrieved all courses.",
		Data:    &c,
	})
}

func (h *HttpCourseHandler) GetCourse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" || !bson.IsObjectIdHex(id) {
		http.Error(w, "Please provide a course id which is a valid mongo id.", http.StatusForbidden)
		return
	}

	c, err := h.CourseUsecase.GetCourse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "Retrieved course.",
		Data:    &c,
	})
}

func (h *HttpCourseHandler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" || !bson.IsObjectIdHex(id) {
		http.Error(w, "Please provide a course id which is a valid mongo id.", http.StatusForbidden)
		return
	}

	err := h.CourseUsecase.DeleteCourse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "Deleted course.",
		Data:    nil,
	})
}


func NewCourseHttpHandler(r *chi.Mux, su CourseUsecase) {
	handler := HttpCourseHandler{
		CourseUsecase: su,
	}

	r.Get("/pa/courses", handler.GetCourses)
	r.Get("/pa/courses/{id}", handler.GetCourse)
	r.Delete("/pa/courses/{id}", handler.DeleteCourse)
	r.Post("/pa/courses", handler.AddCourse)
}
