package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-delivery/pkg/config"
	"github.com/go-delivery/pkg/models"
)

func AddDriver(c *gin.Context) {
	var driver models.Driver

	if err := c.ShouldBindJSON(&driver); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	drivers := models.Driver{Name: driver.Name, Registration: driver.Registration, IdentityCard: driver.IdentityCard}
	config.DB.Create(&drivers)

	c.JSON(200, gin.H{"success": drivers})
}

func GetAllDrivers(c *gin.Context) {
	var drivers []models.Driver

	config.DB.Find(&drivers)
	c.JSON(200, gin.H{"success": drivers})
}

func GetDriverByID(c *gin.Context) {
	var driver models.Driver

	if err := config.DB.Where("id = ?", c.Param("id")).First(&driver).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"success": driver})
}

func UpdateDriver(c *gin.Context) {}

func DeleteDriver(c *gin.Context) {}
