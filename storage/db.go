package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBUser     = "root"
	DBPassword = "root"
	DBName     = "root"
	DBHost     = "0.0.0.0"
	DBPort     = 5432
	DBType     = "postgres"
)

func GetPostgresConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBName, DBPassword)
}

func NewDB() *gorm.DB {
	DB, err := gorm.Open(postgres.Open(GetPostgresConnectionString()), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	return DB
}
