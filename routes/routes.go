package routes

import (
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/controllers"
	"github.com/MalshanPerera/go-expense-tracker/database"
	auth_handlers "github.com/MalshanPerera/go-expense-tracker/handlers/auth"
	expense_handlers "github.com/MalshanPerera/go-expense-tracker/handlers/expense"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func RegisterRoutes() http.Handler {
	db = database.GetDB()

	router := http.NewServeMux()
	apiV1 := http.NewServeMux()

	authController := controllers.NewAuthController(controllers.AuthControllerParams{DB: db})

	apiV1.Handle("/auth/", auth_handlers.Init(authController))
	apiV1.Handle("/expense/", expense_handlers.Init())

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1))

	return router
}
