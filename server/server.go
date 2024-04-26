package core

import (
	"fmt"
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/config"
	"github.com/MalshanPerera/go-expense-tracker/database"
	"github.com/MalshanPerera/go-expense-tracker/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	Echo    *echo.Echo
	Config  *config.Config
	DB      *pgxpool.Pool
	Queries *sqlc.Queries
}

func NewServer(cfg *config.Config) *Server {

	db := database.Connect(cfg)

	return &Server{
		Echo:    echo.New(),
		Config:  cfg,
		DB:      db,
		Queries: sqlc.New(db),
	}
}

func (server *Server) Start() error {

	server.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${status} ${uri} ${latency_human}\n",
	}))
	server.Echo.Use(middleware.Recover())

	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{http.MethodPost, http.MethodGet, http.MethodDelete, http.MethodPatch},
		MaxAge:           86400,
		AllowCredentials: false,
	}))

	err := server.Echo.Start(fmt.Sprintf(":%s", server.Config.HTTP.Port))

	return err
}
