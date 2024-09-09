// Entry point for your application

package main

import (
	"Album/internal/auth"
	"Album/internal/db"
	"Album/internal/tasks"
	"Album/internal/users"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("my_session", store))
	tasks.TaskRoutes(router)
	users.UserRoutes(router)
	auth.AuthRoutes(router)

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
