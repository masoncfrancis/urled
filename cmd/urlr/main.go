package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

type ShortenedURL struct {
	originalURL  string
	shortURLCode string
}

func main() {

	// Check for db

}
