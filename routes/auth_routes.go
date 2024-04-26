package routes

import (
	"github.com/MalshanPerera/go-expense-tracker/handlers"
	"github.com/labstack/echo"
)

type AuthRoute struct {
	V1      *echo.Group
	Handler *handlers.AuthHandler
}

func NewAuthRoute(v1 *echo.Group, authHandler *handlers.AuthHandler) *AuthRoute {
	return &AuthRoute{
		V1:      v1,
		Handler: authHandler,
	}
}

func (route *AuthRoute) RegisterAuthRoutes() {
	route.V1.POST("/auth/login", route.Handler.LoginHandler)
	route.V1.POST("/auth/register", route.Handler.RegisterHandler)
}
