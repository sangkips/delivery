package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-delivery/pkg/config"
	"github.com/go-delivery/pkg/models"
)

func AddToCart(c *gin.Context) {
	var cart models.ShoppingCart

	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// add to cart
	items := models.ShoppingCart{Product: cart.Product, Amount: cart.Amount, ProductID: cart.ProductID}
	config.DB.Create(&items)

	c.JSON(200, gin.H{"items": items})
}
