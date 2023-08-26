package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-delivery/pkg/config"
	"github.com/go-delivery/pkg/models"
)

// Add new product
func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.Header("Content-Type", "application/json")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// create product
	item := models.Product{Name: product.Name, Price: product.Price, Description: product.Description, Brand: product.Brand}
	config.DB.Create(&item)

	c.JSON(200, gin.H{"item": item})

}

// Get product by ID
func GetProductByID(c *gin.Context) {
	var product models.Product

	if err := config.DB.Where("id = ?", c.Param("prod_id")).First(&product).Error; err != nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	var product models.Product

	if err := config.DB.Where("id = ?", c.Param("prod_id")).First(&product).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// validate the item first

	var item models.Product

	if err := c.ShouldBindJSON(&item); err != nil {
		c.Header("Content-Type", "application/json")
		c.JSON(400, gin.H{"error": err.Error()})
	}
	// update product to item

	config.DB.Model(&product).Updates(&item)

	c.JSON(200, gin.H{"success": item})
}

// Delete product
func DeleteProduct(c *gin.Context) {
	var product models.Product

	var existingProduct models.Product
	config.DB.Where("id = ?", c.Param("prod_id")).First(&existingProduct)

	// check if product exist
	if existingProduct.ID == 0 {
		c.JSON(400, gin.H{"error": "product doesn't exist"})
		return
	}

	if err := config.DB.Where("id = ?", c.Param("prod_id")).Find(&product).Error; err != nil {
		c.JSON(400, gin.H{"error": "error deleting the item"})
		return
	}

	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"success": "successfully deleted item"})
}
