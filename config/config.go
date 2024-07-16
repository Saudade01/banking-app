package config

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func LoadConfig() {
	// Tam yol belirterek .env dosyasını yükleyin
	absPath, err := filepath.Abs(".env")
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}

	// Dosyanın var olup olmadığını kontrol et
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		log.Fatalf(".env file does not exist at path: %s", absPath)
	}

	log.Printf("Loading .env file from: %s", absPath)
	err = godotenv.Load(absPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("Environment variables loaded successfully")
}

func InitDB() {
	var err error
	// Environment değişkenlerinden veritabanı bağlantı bilgilerini al
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Debug: Environment değişkenlerini yazdır
	log.Printf("DB_USER: %s, DB_PASSWORD: %s, DB_HOST: %s, DB_NAME: %s", dbUser, dbPassword, dbHost, dbName)

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	log.Println("Database connected successfully")
}
