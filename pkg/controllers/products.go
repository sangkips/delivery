package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-delivery/pkg/config"
	"github.com/go-delivery/pkg/models"
)

// Add new product
func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// create product
	item := models.Product{Name: product.Name, Price: product.Price, Description: product.Description}
	config.DB.Create(&item)

	c.JSON(200, gin.H{"item": item})

}

// Get product by ID
func GetProductByID(c *gin.Context) {
	var product models.Product

	if err := config.DB.Where("id = ?", c.Param("prod_id")).First(&product).Error; err != nil {
		c.JSON(400, gin.H{"error": "error retrieving item"})
		return
	}
	c.JSON(200, gin.H{"item": product})
}

// Get all products
func GetProducts(c *gin.Context) {
	var products []models.Product

	config.DB.Find(&products)
	c.JSON(200, gin.H{"items": products})
}

// Update product
func UpdateProduct(c *gin.Context) {
	return
}

// Delete product
func DeleteProduct(c *gin.Context) {
	return
}
