// HTTP handlers for task-related routes
package tasks

import "github.com/gin-gonic/gin"

func TaskRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/todos", getTasks)
	router.POST("/todos", createTask)
	router.PUT("/todos/:id", updateTask)
	router.GET("/todos/:id", getTaskId)
	router.DELETE("/todos/:id", deleteTask)
	return router
}
