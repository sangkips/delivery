package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-delivery/pkg/models"
)

func CreateAddress(c *gin.Context) {
	var address models.Address

	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
