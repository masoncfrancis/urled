package main

import (
	"flag"
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type URLrecord struct {
	gorm.Model
	ShortURL string
	LongURL  string
}

func main() {
	// Load .env file
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Error loading .env file. Please check to make sure it exists and is formatted correctly.")
		os.Exit(1)
	}

	// Setup logging
	myLogger := setupLogging()

	// Connect to the database
	db, dbErr := gorm.Open(sqlite.Open("urledfiles/urled.db"), &gorm.Config{
		Logger: myLogger,
	})
	if dbErr != nil {
		fmt.Println("Error connecting to database")
		os.Exit(1)
	}

	// Migrate the schema
	db.AutoMigrate(&URLrecord{})

	// Parse flags
	addFlag := flag.String("add", "", "Add a new URL")
	listFlag := flag.Bool("list", false, "List all URLs")
	removeByShortURLFlag := flag.String("remove", "", "Remove a URL using the short URL suffix")
	removeByLongURLFlag := flag.String("remove-long", "", "Remove a URL using the long URL")
	flag.Parse()

	// Add a new URL, generating short code
	if *addFlag != "" {
		// Validate url
		if !validateURL(*addFlag) {
			return // exit
		}
		// trim trailing slash
		*addFlag = strings.TrimSuffix(*addFlag, "/")

		shortURL := generateShortURL(db)
		db.Create(&URLrecord{ShortURL: shortURL, LongURL: *addFlag})

		fmt.Println(*addFlag + " added successfully")
		fmt.Println("Short URL: " + os.Getenv("BASE_URL") + "/" + shortURL)
		return // exit
	}

	// List all URLs
	if *listFlag {
		var urls []URLrecord
		db.Find(&urls)
		if len(urls) == 0 {
			fmt.Println("No URLs found")
			return // exit
		}
		for _, urlItem := range urls {
			fmt.Println(urlItem.LongURL + " -> " + os.Getenv("BASE_URL") + "/" + urlItem.ShortURL)
		}
		return // exit
	}

	// Remove a URL by short URL
	if *removeByShortURLFlag != "" {
		var urlRecord URLrecord
		result := db.Where("short_url = ?", *removeByShortURLFlag).First(&urlRecord)
		if result.RowsAffected == 0 {
			fmt.Println("Error: Short URL not found")
			return // exit
		}
		db.Delete(&urlRecord)
		fmt.Println("URL(s) removed successfully")
		return // exit
	}

	// Remove a URL by long URL
	if *removeByLongURLFlag != "" {
		fmt.Println("Note: This will remove all URLs with the same long URL")
		var urlRecord URLrecord
		result := db.Where("long_url = ?", *removeByLongURLFlag).Delete(&urlRecord)
		if result.RowsAffected == 0 {
			fmt.Println("Error: Long URL not found")
			return // exit
		}
		fmt.Println("URL removed successfully")
		return // exit
	}

}

func setupLogging() logger.Interface {
	// Define the log file path
	logFilePath := "urledfiles/logs/gorm.log"

	// Create the directory if it doesn't exist
	dir := filepath.Dir(logFilePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Open a file for logging
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Error, // Log level
			Colorful: false,        // Disable color
		},
	)

	return newLogger
}

func validateURL(inputURL string) bool {
	_, err := url.ParseRequestURI(inputURL)
	if err != nil {
		fmt.Println("Error: Invalid URL")
		return false
	}
	return true
}

func generateShortURL(db *gorm.DB) string {
	shortURL := ""
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < 6; i++ {
		shortURL += string(charSet[rand.Intn(len(charSet))])
	}

	// Check db if shortURL already exists, generate a new one if it does
	var short URLrecord
	result := db.Where("short_url = ?", shortURL).First(&short)
	if result.RowsAffected > 0 {
		return generateShortURL(db)
	}

	return shortURL
}
