package https

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kiranraj27/sprint-planner/internal/domain"
)

type UserHandler struct {
    userUseCase domain.UserUseCase
}

func NewUserHandler(router *mux.Router, uc domain.UserUseCase) {
    handler := &UserHandler{uc}
    router.HandleFunc("/users", handler.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
    router.HandleFunc("/users", handler.ListUsers).Methods("GET")
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    user := &domain.User{}
    json.NewDecoder(r.Body).Decode(user)
    if err := h.userUseCase.Register(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    user, err := h.userUseCase.GetUser(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
    users, err := h.userUseCase.ListUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(users)
}
