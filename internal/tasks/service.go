// Business logic for tasks
package tasks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func getTasks(c *gin.Context) {
	getAllTasks(c)
}

var validate = validator.New()

func createTask(c *gin.Context) {
	var todo Task
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if err := validate.Struct(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := createTaskInDB(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func updateTask(c *gin.Context) {
	var todo Task
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
	}
	if err := updateTaskById(id, todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error with updating"})
		return
	}
	c.JSON(http.StatusOK, todo)

}

func deleteTask(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	if err := deleteTaskbyId(id); err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})

}
