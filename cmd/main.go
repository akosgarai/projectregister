package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"

	"github.com/akosgarai/projectregister/pkg/application"
)

var (
	wait       time.Duration = 15 * time.Second
	dotEnvFile               = ".env"
)

func main() {
	dotenvConfig := make(map[string]string)
	dotenvConfig, err := godotenv.Read(dotEnvFile)
	if err != nil {
		log.Println("Error loading .env file")
	}
	// load .env file
	app := application.New(dotenvConfig)
	app.Initialize()

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := app.Run(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	app.Server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
