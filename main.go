package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user_management/controllers"
	"user_management/models"
	"user_management/routes"
	"user_management/utils"
)

func main() {

	db, err := utils.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close() // Ensure the database connection is closed after the test

	// Initialize repository and controller
	userRepo := models.NewUserRepository(db)
	userController := &controllers.UserController{Repo: userRepo}

	e := echo.New()

	// Set up routes
	routes.SetupRoutes(e, userController)

	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: e,
	}

	// Run server in a goroutine
	go func() {
		log.Println("Starting server on port 8080")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
