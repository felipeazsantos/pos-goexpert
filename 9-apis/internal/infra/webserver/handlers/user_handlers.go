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
	db database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{db: db}
}

// GetJWT godoc
// @Summary Get JWT token
// @Description Get JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param input body dto.GetJWTInput true "User input"
// @Success 200 {object} dto.GetJWTOutput
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var input dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errResponse := dto.ErrorResponse{Message: err.Error()}
		http.Error(w, errResponse.Message, http.StatusBadRequest)
		return
	}

	user, err := h.db.FindByEmail(input.Email)
	if err != nil {
		errResponse := dto.ErrorResponse{Message: err.Error()}
		http.Error(w, errResponse.Message, http.StatusNotFound)
		return
	}

	if !user.ValidatePassword(input.Password) {
		errResponse := dto.ErrorResponse{Message: "Invalid credentials"}
		http.Error(w, errResponse.Message, http.StatusUnauthorized)
		return
	}

	_, token, err := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	if err != nil {
		errResponse := dto.ErrorResponse{Message: err.Error()}
		http.Error(w, errResponse.Message, http.StatusInternalServerError)
		return
	}

	accessToken := dto.GetJWTOutput{
		AccessToken: token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param input body dto.CreateUserInput true "User input"
// @Success 201 {object} entity.User
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		errResponse := dto.ErrorResponse{Message: err.Error()}
		http.Error(w, errResponse.Message, http.StatusBadRequest)
		return
	}

	err = h.db.Create(user)
	if err != nil {
		errResponse := dto.ErrorResponse{Message: err.Error()}
		http.Error(w, errResponse.Message, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
