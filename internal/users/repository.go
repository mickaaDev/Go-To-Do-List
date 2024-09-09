// Database interactions (CRUD operations)
package users

import (
	"Album/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllUsers(c *gin.Context) {
	var users []User

	if err := db.Database.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve userse"})
		return
	}
	c.JSON(http.StatusOK, users)
}
