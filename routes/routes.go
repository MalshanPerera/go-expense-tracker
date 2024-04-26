package routes

import (
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/database"
	"github.com/MalshanPerera/go-expense-tracker/database/sqlc"
	"github.com/MalshanPerera/go-expense-tracker/handlers"
	"github.com/MalshanPerera/go-expense-tracker/middlewares"
	"github.com/MalshanPerera/go-expense-tracker/services"
)

type Route struct {
	AuthHandler    *handlers.AuthHandler
	ExpenseHandler *handlers.ExpenseHandler
}

func NewRoute() http.Handler {
	db := database.GetDB()
	queries := sqlc.New(db)

	authService := services.NewAuthService(db, queries)
	expenseService := services.NewExpenseService(db, queries)

	authHandler := handlers.NewAuthHandler(authService)
	expenseHandler := handlers.NewExpenseHandler(expenseService)

	return NewRouter(&Route{
		AuthHandler:    authHandler,
		ExpenseHandler: expenseHandler,
	})
}

func NewRouter(r *Route) http.Handler {
	router := http.NewServeMux()

	apiV1 := r.registerRoutes()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1))

	return router
}

func (r *Route) registerRoutes() *http.ServeMux {
	apiV1 := http.NewServeMux()

	authRoute := NewAuthRoute(r.AuthHandler)
	expenseRoute := NewExpenseRoute(r.ExpenseHandler)

	initHandler(apiV1, "/auth/", authRoute.RegisterAuthRoutes())
	initHandler(apiV1, "/expense", r.registerProtectedRoutes(expenseRoute.RegisterExpenseRoutes()))

	return apiV1
}

func (r *Route) registerProtectedRoutes(handler http.Handler) http.Handler {
	return middlewares.CreateStack(middlewares.IsAuthenticated)(handler)
}

func initHandler(mux *http.ServeMux, pattern string, handler http.Handler) {
	mux.Handle(pattern, handler)
}
