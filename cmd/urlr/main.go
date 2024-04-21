package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"net/url"
	"os"
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

	// Check for sqlite db
	db, dbErr := gorm.Open(sqlite.Open("urlr.db"), &gorm.Config{})
	if dbErr != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&URLrecord{})

	// Parse flags
	addFlag := flag.String("add", "", "Add a new URL")
	//listFlag := flag.Bool("list", false, "List all URLs")
	//removeFlag := flag.String("remove", "", "Remove a URL")
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

		fmt.Println("Short URL: " + os.Getenv("BASE_URL") + "/" + shortURL)
		return // exit
	}

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
	var url URLrecord
	db.Where("short_url = ?", shortURL).First(&url)
	if url.ShortURL != "" {
		shortURL = generateShortURL(db)
	}

	return shortURL
}
