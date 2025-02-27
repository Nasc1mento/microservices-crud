package handlers

import (
	"encoding/json"
	"fmt"
	"microservices-crud/user-service/db/user"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandler struct {
	q *user.Queries
	r *chi.Mux
}

func NewUserHandler(q *user.Queries) UserHandler {
	a := UserHandler{
		q: q,
	}

	r := chi.NewRouter()

	a.r = r

	return a
}

func (h UserHandler) Routes() chi.Router {
	h.r.Post("/", h.Create)
	h.r.Delete("/{id}", h.Delete)
	h.r.Get("/email/{email}", h.GetUserByEmail)
	h.r.Get("/{id}", h.GetUserById)

	return h.r
}

func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.q.CreateUser(r.Context(), user.CreateUserParams{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})

	fmt.Println(created)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.q.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	u, err := h.q.GetUserByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func (h UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.q.GetUserById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(u)
}
