package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MalshanPerera/go-expense-tracker/config"
	"github.com/MalshanPerera/go-expense-tracker/database"
	"github.com/MalshanPerera/go-expense-tracker/routes"
	server "github.com/MalshanPerera/go-expense-tracker/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := config.NewConfig()
	defer database.Close()

	app := server.NewServer(cfg)

	// Register routes
	routes.NewRoute(app).RegisterRoutes()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		log.Printf("Server is running at http://localhost%s\n", cfg.HTTP.Port)

		err := app.Start()

		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server has shut down.")
		} else if err != nil {
			log.Fatalf("Server failed to start: %v\n", err)
			os.Exit(1)
		}

	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Echo.Shutdown(ctx); err != nil {
		app.Echo.Logger.Fatal(err)
	}
}
