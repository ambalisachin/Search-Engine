// Package config provides configuration settings for the search engine application.
package config

import (
	"fmt"
	"log"
	"os"
	"search-engine/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// ConnectDB function establishes a connection to the database and returns a pointer to the database object.
func ConnectDB() (*gorm.DB, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return nil, err
	}

	// Get database connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Create a connection string using the database connection details
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open a connection to the database using the connection string
	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Panicf("failed to connect database: %v", err)
		return nil, err
	}

	// Automigrate the schema to create database tables
	err = db.AutoMigrate(&models.Book{}).Error
	if err != nil {
		log.Panicf("failed to migrate schema: %v", err)
		return nil, err
	}

	// Return a pointer to the database object
	return db, nil
}
