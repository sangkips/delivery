package controllers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-delivery/pkg/config"
	"github.com/go-delivery/pkg/models"
	"github.com/go-delivery/pkg/utils"
)

// handle login requests from the user
func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	config.DB.Where("email = ?", user.Email).First(&existingUser)

	// check if user exist
	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user doesn't exist"})
		return
	}

	// compare given password with what is in the db
	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"StandardClaims": jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(500, gin.H{"error": "Can not generate token"})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "successfully logged in"})
}

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//check if user already exists
	var existingUser models.User

	config.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	//generate hashed password
	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "Error generating hashed password"})
		return
	}

	config.DB.Create(&user)
	c.JSON(200, gin.H{"data": user})
}

// clears the token once user logout
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged out"})
}

// Read users

func GetUsers(c *gin.Context) {
	var users []models.User

	config.DB.Find(&users)

	c.JSON(200, gin.H{"data": users})

}

// Get user by ID
func GetUSerByID(c *gin.Context) {
	var user models.User

	if err := config.DB.Where("id = ?", c.Param("user_id")).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "user not find"})
		return
	}

	c.JSON(200, gin.H{"data": user})
}
