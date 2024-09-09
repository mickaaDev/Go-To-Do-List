// HTTP handlers for user-related routes
package users

import "github.com/gin-gonic/gin"

func UserRoutes(router *gin.Engine) {

	router.GET("/user/me", getUser)
	router.GET("/users", getUsers)
	// router.PATCH("/user/edit", updateUser)
	// router.DELETE()
}
