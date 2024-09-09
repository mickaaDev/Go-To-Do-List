package auth

import (
	"Album/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	router.POST("/auth/signup", CreateUser)
	router.POST("/auth/login", Login)
	router.GET("/user/profile", utils.CheckAuth, GetUserProfile)
	// router.GET("/user/logout", utils.CheckAuth, LogOut)
}
