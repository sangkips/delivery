package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-delivery/pkg/controllers"
)

func AuthenticationRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.GET("/logout", controllers.Logout)
}
