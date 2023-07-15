package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-delivery/pkg/config"
	"github.com/go-delivery/pkg/models"
	"github.com/go-delivery/pkg/utils"
)

var jwtkey = []byte("nfdnf5534784nmnmdfj")

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

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"Role": existingUser.Role,
		"StandardClaims": jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	// expirationTime := time.Now().Add(5 * time.Minute)

	// claims := &models.Claims{

	// }

	tokenString, err := token.SignedString(jwtkey)

	if err != nil {
		c.JSON(500, gin.H{"error": "Can not generate token"})
		return
	}

	// c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": tokenString})
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
	c.JSON(200, gin.H{"success": "User successfully created"})
}

// Renders home page and user role logged in
func Home(c *gin.Context) {
	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse token used to allow stay in the home page
	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if claims.Role != "user" && claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(200, gin.H{"success": "home page", "role": claims.Role})
}

func Premium(c *gin.Context) {
	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(200, gin.H{"success": "premium page", "role": claims.Role})
}

// clears the token once user logout
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged out"})
}
