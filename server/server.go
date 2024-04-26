package core

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/MalshanPerera/go-expense-tracker/middlewares"
	"github.com/MalshanPerera/go-expense-tracker/routes"
	"github.com/rs/cors"
)

func getPort() int {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000
	}

	return port
}

func NewServer() *http.Server {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodDelete,
			http.MethodPatch,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	stack := middlewares.CreateStack(
		middlewares.Logger,
	)

	handler := routes.NewRoute()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", getPort()),
		Handler:      stack(handler),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Wrap the server with the CORS middleware.
	server.Handler = c.Handler(server.Handler)

	return server
}
