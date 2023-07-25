package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-delivery/pkg/controllers"
)

func AuthenticationRoutes(r *gin.Engine) {
	// Authentication
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.GET("/logout", controllers.Logout)

	// Users
	r.GET("/users", controllers.GetUsers)
	r.GET("/user/:user_id", controllers.GetUSerByID)

	// Products
	r.POST("/add", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)
	r.GET("/product/:prod_id", controllers.GetProductByID)
	r.PUT("/update/:prod_id", controllers.UpdateProduct)
	r.DELETE("/product/:prod_id", controllers.DeleteProduct)
}
