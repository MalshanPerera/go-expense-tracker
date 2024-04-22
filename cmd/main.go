package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MalshanPerera/go-expense-tracker/database"
	"github.com/MalshanPerera/go-expense-tracker/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.Connect()
	defer database.Close()

	server := server.NewServer()

	// Create a channel to listen for interrupt or termination signals from the OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server is running at http://localhost%s\n", server.Addr)

		err := server.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server has shut down.")
		} else if err != nil {
			log.Fatalf("Server failed to start: %v\n", err)
			os.Exit(1)
		}
	}()

	// Block until we receive a shutdown signal
	<-shutdown

	// Create a context to potentially cancel the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server failed to shut down gracefully: %v\n", err)
		os.Exit(1)
	}
}
