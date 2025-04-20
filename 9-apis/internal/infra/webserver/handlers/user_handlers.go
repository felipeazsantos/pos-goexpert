package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/felipeazsantos/pos-goexpert/apis/internal/dto"
	"github.com/felipeazsantos/pos-goexpert/apis/internal/entity"
	"github.com/felipeazsantos/pos-goexpert/apis/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	db           database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{db: db, Jwt: jwt, JwtExpiresIn: jwtExpiresIn}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var input dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.db.FindByEmail(input.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if !user.ValidatePassword(input.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	_, token, err := h.Jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
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
