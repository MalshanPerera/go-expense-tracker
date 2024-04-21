package auth_handlers

import (
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/controllers"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthControllerInterface controllers.AuthControllerInterface

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

	// var user LoginUser
	// if err := utils.ParseJSON(r, &user); err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err)
	// 	return
	// }

	// err := authController.Login(user.Email, user.Password)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// user = LoginUser{
	// 	Email:    loggedInUser.Email,
	// 	Password: loggedInUser.Password,
	// }

	// userJson, err := json.Marshal(user)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(userJson)

	w.Write([]byte("Login Handler"))
}

func registerHandler(authController AuthControllerInterface, w http.ResponseWriter, r *http.Request) {

	// var user schema.User
	// if err := utils.ParseJSON(r, &user); err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err)
	// 	return
	// }

	// registeredUser, err := authController.Register(user)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// userJson, err := json.Marshal(registeredUser)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(userJson)

	w.Write([]byte("Register Handler"))
}
