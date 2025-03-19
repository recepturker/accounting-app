package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("PostgreSQL bağlantısı başarısız: ", err)
	}

	log.Println("PostgreSQL bağlantısı başarılı!")
}
