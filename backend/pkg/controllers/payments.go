package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-delivery/pkg/models"
)

func MakePayment(c *gin.Context) {
	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
