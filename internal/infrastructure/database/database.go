package database

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/KalinduGandara/erp-system/internal/domain/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase(dbPath string) (*gorm.DB, error) {
	// Create database directory if it doesn't exist
	dbDir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	// Open database connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(
		&entities.User{},
		&entities.Customer{},
		&entities.Product{},
		&entities.Invoice{},
		&entities.InvoiceItem{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	// Create default admin user
	if err := createDefaultAdmin(db); err != nil {
		return nil, fmt.Errorf("failed to create default admin: %w", err)
	}

	return db, nil
}

func createDefaultAdmin(db *gorm.DB) error {
	var count int64
	db.Model(&entities.User{}).Count(&count)
	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		admin := &entities.User{
			Username: "admin",
			Password: string(hashedPassword),
			Role:     "admin",
		}
		return db.Create(admin).Error
	}
	return nil
}
