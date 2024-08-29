// Entry point for your application

package main

import (
	"Album/internal/db"
	"Album/internal/tasks"
	"log"
)

func main() {
	router := tasks.TaskRoutes()
	loadDatabase()
	if err := router.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loadDatabase() {
	db.ConnectDatabase()
	db.Database.AutoMigrate(&tasks.Task{})
}
