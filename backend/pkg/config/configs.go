package config

import (
	"fmt"

	"github.com/go-delivery/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func InitDB(cfg Config) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Address{},
		&models.Driver{},
		&models.Payment{},
		&models.Order{}); err != nil {
		panic(err)
	}

	fmt.Println("Database migrated successfully")

	DB = db
}
