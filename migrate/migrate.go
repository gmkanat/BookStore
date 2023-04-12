package main

import (
	"BookStore/initializers"
	"BookStore/models"
	"fmt"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.Book{})
	fmt.Println("? Migration complete")
}
