package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-delivery/pkg/models"
)

func AddToCart(c *gin.Context) {
	var cart models.Order

	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// add to cart
}
