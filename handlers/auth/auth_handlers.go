package auth_handlers

import (
	"fmt"
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/modals"
	"github.com/MalshanPerera/go-expense-tracker/services"
	"github.com/MalshanPerera/go-expense-tracker/utils"
	"github.com/go-playground/validator/v10"
)

type AuthControllerInterface services.AuthServiceInterface

func Init(authController AuthControllerInterface) http.Handler {
	authHandlers := http.NewServeMux()

	authHandlers.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		loginHandler(authController, w, r)
	})
	authHandlers.HandleFunc("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		registerHandler(authController, w, r)
	})

	return authHandlers
}

func loginHandler(authController AuthControllerInterface, w http.ResponseWriter, r *http.Request) {

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

	loggedInUser, err := authController.Login(r.Context(), user)
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

func registerHandler(authController AuthControllerInterface, w http.ResponseWriter, r *http.Request) {

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

	registeredUser, err := authController.Register(r.Context(), user)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, registeredUser, http.StatusOK)

}
