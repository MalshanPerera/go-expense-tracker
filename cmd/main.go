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
	"github.com/MalshanPerera/go-expense-tracker/routes"
	server "github.com/MalshanPerera/go-expense-tracker/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := config.NewConfig()

	app := server.NewServer(cfg)
	defer app.Close()

	routes := routes.NewRoute(app)
	routes.RegisterRoutes()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		log.Printf("Server is running at http://localhost%s\n", cfg.HTTP.Port)

		if err := app.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v\n", err)
		} else {
			log.Println("Server has shut down.")
		}

	}()

	<-ctx.Done()

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	shutdownError := make(chan error, 1)
	go func() {
		shutdownError <- app.Echo.Shutdown(ctxShutdown)
	}()

	select {
	case err := <-shutdownError:
		if err != nil {
			app.Echo.Logger.Fatal(err)
		}
	case <-ctxShutdown.Done():
		app.Echo.Logger.Fatal("shutdown timeout")
	}
}
