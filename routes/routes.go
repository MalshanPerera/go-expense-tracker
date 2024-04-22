package routes

import (
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/controllers"
	"github.com/MalshanPerera/go-expense-tracker/database"
	"github.com/MalshanPerera/go-expense-tracker/database/sqlc"
	auth_handlers "github.com/MalshanPerera/go-expense-tracker/handlers/auth"
	expense_handlers "github.com/MalshanPerera/go-expense-tracker/handlers/expense"
)

func RegisterRoutes() http.Handler {
	db := database.GetDB()
	queries := sqlc.New(db)

	router := http.NewServeMux()
	apiV1 := http.NewServeMux()

	authController := controllers.NewAuthController(controllers.AuthControllerParams{DB: db, Queries: queries})

	apiV1.Handle("/auth/", auth_handlers.Init(authController))
	apiV1.Handle("/expense/", expense_handlers.Init())

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1))

	return router
}
