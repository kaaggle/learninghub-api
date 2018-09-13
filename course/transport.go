package course

import (
	"encoding/json"
	"net/http"
	"schoolsystem/learninghub-api/models"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/chi"
)

type HttpCourseHandler struct {
	CourseUsecase CourseUsecase
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

func NewCourseHttpHandler(r *chi.Mux, su CourseUsecase) {
	handler := HttpCourseHandler{
		CourseUsecase: su,
	}

	r.Get("/pa/courses", handler.GetCourses)
	r.Get("/pa/courses/{id}", handler.GetCourse)
}
