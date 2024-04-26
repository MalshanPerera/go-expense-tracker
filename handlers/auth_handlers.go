package handlers

import (
	"fmt"
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/modals"
	"github.com/MalshanPerera/go-expense-tracker/services"
	"github.com/MalshanPerera/go-expense-tracker/utils"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	AuthService services.AuthServiceInterface
}

func NewAuthHandler(authService services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user modals.LoginUserParams
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, fmt.Errorf("invalid payload: %v", errors), http.StatusBadRequest)
		return
	}

	loggedInUser, err := h.AuthService.Login(r.Context(), user)
	if loggedInUser == nil {
		utils.WriteError(w, err, http.StatusFound)
		return
	}

	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, loggedInUser, http.StatusOK)
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var user modals.CreateUserParams
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, fmt.Errorf("invalid payload: %v", errors), http.StatusBadRequest)
		return
	}

	registeredUser, err := h.AuthService.Register(r.Context(), user)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, registeredUser, http.StatusOK)

}
