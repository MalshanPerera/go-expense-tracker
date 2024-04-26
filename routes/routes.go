package routes

import (
	"github.com/MalshanPerera/go-expense-tracker/handlers"
	m "github.com/MalshanPerera/go-expense-tracker/middlewares"
	"github.com/MalshanPerera/go-expense-tracker/services"

	s "github.com/MalshanPerera/go-expense-tracker/server"
)

type Route struct {
	Server         *s.Server
	AuthHandler    *handlers.AuthHandler
	ExpenseHandler *handlers.ExpenseHandler
}

func NewRoute(server *s.Server) *Route {

	db := server.DB
	queries := server.Queries

	authService := services.NewAuthService(db, queries)
	expenseService := services.NewExpenseService(db, queries)

	authHandler := handlers.NewAuthHandler(authService)
	expenseHandler := handlers.NewExpenseHandler(expenseService)

	return &Route{
		Server:         server,
		AuthHandler:    authHandler,
		ExpenseHandler: expenseHandler,
	}
}

func (r *Route) RegisterRoutes() {

	v1 := r.Server.Echo.Group("/api/v1")

	authRoute := NewAuthRoute(v1, r.AuthHandler)
	authRoute.RegisterAuthRoutes()

	protectedRoutes := v1.Group("")
	protectedRoutes.Use(m.IsAuthenticated)

	expenseRoute := NewExpenseRoute(protectedRoutes, r.ExpenseHandler)
	expenseRoute.RegisterExpenseRoutes()

}
