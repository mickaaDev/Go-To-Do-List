// HTTP handlers for task-related routes
package tasks

import (
	"Album/internal/utils"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {

	router.GET("/todos", getTasks)
	router.POST("/todos", utils.CheckAuth, createTask)
	router.PUT("/todos/:id", updateTask)
	router.GET("/todos/:id", getTaskId)
	router.DELETE("/todos/:id", deleteTask)
}
