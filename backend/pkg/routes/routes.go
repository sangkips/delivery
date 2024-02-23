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

	// Shopping cart
	r.POST("/cart", controllers.AddToCart)

	// Address
	r.POST("/address", controllers.CreateAddress)

	// Payment
	r.POST("/payment", controllers.MakePayment)

	// driver
	r.POST("/add/driver", controllers.AddDriver)
	r.GET("/drivers", controllers.GetAllDrivers)
	r.PUT("/update/driver/:id", controllers.UpdateDriver)
	r.GET("/driver/:id", controllers.GetDriverByID)
	r.DELETE("/delete/driver/:id", controllers.DeleteDriver)
}
