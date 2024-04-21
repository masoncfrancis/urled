package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

type URLrecord struct {
	gorm.Model
	ShortURL string
	LongURL  string
}

func main() {
	// Check for sqlite db
	db, err := gorm.Open(sqlite.Open("urlr.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&URLrecord{})

	// Parse flags
	addFlag := flag.String("add", "", "Add a new URL")
	listFlag := flag.Bool("list", false, "List all URLs")
	removeFlag := flag.String("remove", "", "Remove a URL")
	flag.Parse()

}
