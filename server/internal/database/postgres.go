package database

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"server/internal/config"
)

var db *gorm.DB

func getPostgresConnectionString() string {
	dbConfig := config.GetDbConfig()
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
}

func InitPostgres() {
	conn, err := getPostgresConnection()
	if err != nil {
		log.Fatal("Error while connecting to postgres:", err)
	}
	db = conn
	log.Println("Connected to postgres")
}

func getPostgresConnection() (*gorm.DB, error) {
	gormConfig := &gorm.Config{}
	connStr := getPostgresConnectionString()

	db, err := gorm.Open(postgres.Open(connStr), gormConfig)
	if err != nil {
		return nil, err
	}

	err = migrateDb(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetPostgresConnection() *gorm.DB {
	return db
}
