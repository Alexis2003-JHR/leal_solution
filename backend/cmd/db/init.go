package db

import (
	"fmt"
	"os"

	ormModels "leal/internal/core/domain/db"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	server := os.Getenv("PG_SERVER")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_DB_NAME")
	port := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", server, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&ormModels.User{},
		&ormModels.Business{},
		&ormModels.Branch{},
		&ormModels.ConversionFactor{},
		&ormModels.Transaction{},
		&ormModels.Campaign{},
		&ormModels.Earnings{},
		&ormModels.Reward{},
		&ormModels.Redemption{},
	)

	return db, nil
}
