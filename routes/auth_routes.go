package routes

import (
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/handlers"
)

type AuthRoute struct {
	Handler *handlers.AuthHandler
}

func NewAuthRoute(authHandler *handlers.AuthHandler) *AuthRoute {
	return &AuthRoute{
		Handler: authHandler,
	}
}

func (route *AuthRoute) RegisterAuthRoutes() http.Handler {
	authHandlers := http.NewServeMux()

	handlersMap := map[string]func(w http.ResponseWriter, r *http.Request){
		"/login":    route.Handler.LoginHandler,
		"/register": route.Handler.RegisterHandler,
	}

	for pattern, handlerFunc := range handlersMap {
		handler := handlers.HandleFunc(handlerFunc)
		authHandlers.HandleFunc(pattern, handler)
	}

	authHandlers.Handle("/auth/", http.StripPrefix("/auth", authHandlers))

	return authHandlers
}
