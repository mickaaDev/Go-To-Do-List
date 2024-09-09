// Database interactions (CRUD operations)
package tasks

import (
	"Album/internal/db"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllTasks(c *gin.Context) {
	var tasks []Task

	if err := db.Database.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func createTaskInDB(task *Task) error {
	if err := db.Database.Create(task).Error; err != nil {
		log.Printf("Error creating task: %v", err)
		return err
	}
	return nil
}

func updateTaskById(id int, todo Task) error {

	var existingTask Task

	if err := db.Database.First(&existingTask, id).Error; err != nil {
		return err
	}
	if todo.Title != "" {
		existingTask.Title = todo.Title
	}
	if todo.Description != "" {
		existingTask.Description = todo.Description
	}
	existingTask.Completed = todo.Completed

	if err := db.Database.Save(&existingTask).Error; err != nil {
		return err
	}

	return nil
}

func getTaskId(c *gin.Context) {
	var task Task

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := db.Database.Where("id = ?", id).First(&task).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
		}
		return
	}

	c.JSON(http.StatusOK, task)
}

func deleteTaskbyId(id int) error {
	var task Task

	if err := db.Database.Where("id = ?", id).First(&task).Error; err != nil {
		return err
	}

	result := db.Database.Delete(task)

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("no task record was deleted")
	}
	return nil

}
