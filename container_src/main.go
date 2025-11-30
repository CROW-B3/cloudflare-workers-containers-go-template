package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"server/config"
	"server/internal/handlers"
	"server/internal/repositories"
	"server/internal/routes"
	"server/internal/services"
	"server/pkg/database"
	"server/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	if err := logger.InitLogger(cfg.App.Environment); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Info("Starting application")

	// Initialize database
	if err := database.InitDatabase(cfg); err != nil {
		logger.Fatal("Failed to initialize database")
	}
	defer database.CloseDatabase()

	// Auto-migrate database models
	// Uncomment when ready to use database
	// if err := database.DB.AutoMigrate(&models.User{}); err != nil {
	// 	logger.Fatal("Failed to migrate database")
	// }

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Initialize router
	router := gin.New()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(database.DB)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler(cfg)
	userHandler := handlers.NewUserHandler(userService)

	// Setup routes
	routes.SetupRoutes(router, healthHandler, userHandler)

	// Create server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start server in goroutine
	go func() {
		logger.Info(fmt.Sprintf("Server starting on %s", addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit

	logger.Info(fmt.Sprintf("Received signal (%s), shutting down server...", sig))

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Server.ShutdownTimeout)*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown")
	}

	logger.Info("Server shutdown successfully")
}
