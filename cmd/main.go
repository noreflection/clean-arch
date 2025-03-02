package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-cqrs/internal/infrastructure/container"
	"go-cqrs/internal/infrastructure/logger"
)

func main() {
	// Create application container
	app, err := container.NewContainer()
	if err != nil {
		fmt.Printf("Failed to initialize application: %v\n", err)
		os.Exit(1)
	}
	defer app.Close()

	// Create server
	srv := &http.Server{
		Addr:         app.Config.ServerAddress(),
		Handler:      app.Router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		app.Logger.Info("Server is running", logger.String("address", app.Config.ServerAddress()))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("Failed to start server", logger.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.Logger.Info("Shutting down server...")

	// Create context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server
	if err := srv.Shutdown(ctx); err != nil {
		app.Logger.Fatal("Server forced to shutdown", logger.Error(err))
	}

	app.Logger.Info("Server shutdown complete")
	app.Logger.Sync()
}
