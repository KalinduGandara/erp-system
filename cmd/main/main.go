package main

import (
	"path/filepath"

	"github.com/KalinduGandara/erp-system/internal/app/services"
	"github.com/KalinduGandara/erp-system/internal/domain/repositories"
	"github.com/KalinduGandara/erp-system/internal/infrastructure/database"
	"github.com/KalinduGandara/erp-system/internal/presentation/views"
	"github.com/KalinduGandara/erp-system/pkg/common/config"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}

	// Create absolute path for database
	dbPath := filepath.Join(".", cfg.Database.Path)

	// Initialize database
	db, err := database.NewDatabase(dbPath)
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}

	// Initialize JWT service
	// jwtService, err := auth.NewJWTService(cfg.Auth.JWTSecret, cfg.Auth.TokenExpiry)
	// if err != nil {
	// 	logger.Fatal("Failed to initialize JWT service", zap.Error(err))
	// }

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Create and start main window
	mainWindow := views.NewMainWindow(userService)
	mainWindow.Show()
}
