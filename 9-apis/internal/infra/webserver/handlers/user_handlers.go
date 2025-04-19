package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/felipeazsantos/pos-goexpert/apis/internal/dto"
	"github.com/felipeazsantos/pos-goexpert/apis/internal/entity"
	"github.com/felipeazsantos/pos-goexpert/apis/internal/infra/database"
)

type UserHandler struct {
	db database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.db.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
