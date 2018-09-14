package user

import (
	"encoding/json"
	"net/http"
	"schoolsystem/learninghub-api/models"
	"time"

	"github.com/go-chi/chi"
)

type HttpUserHandler struct {
	UserUsecase UserUsecase
}

func (h *HttpUserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.CreatedAt = time.Now()
	u.Role = "LEARNINGHUB_USER"

	err := u.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	c, err := h.UserUsecase.Signup(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "User registered successfully.",
		Data:    &c,
	})
}

func NewUserHttpHandler(r *chi.Mux, su UserUsecase) {
	handler := HttpUserHandler{
		UserUsecase: su,
	}

	r.Post("/p/signup", handler.Signup)
}
