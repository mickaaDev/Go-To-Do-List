// Entry point for your application

package main

import (
	"Album/internal/db"
	"Album/internal/tasks"
	"Album/internal/users"
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
	err := db.Database.AutoMigrate(&tasks.Task{})
	if err != nil {
		log.Fatalf("Failed to migrate task mode: %v", err)
	}

	err = db.Database.AutoMigrate(&users.User{})
	if err != nil {
		log.Fatalf("Failed to migrate Task model: %v", err)
	}
}
