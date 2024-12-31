package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prathamesh/login_jwt_project/connections"
	"github.com/prathamesh/login_jwt_project/helpers"
	"github.com/prathamesh/login_jwt_project/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// SignUpUser handles the signup functionality
func SignUpUser(c *gin.Context) {
	var user models.SignUpUser

	// Validate payload
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if username or email already exists
	collection := connections.GetCollection("jwt_signup_login", "user_data")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.SignUpUser
	err := collection.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"username": user.Username},
			{"email": user.Email},
		},
	}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
		return
	}

	// Validate and hash the password
	if err := helpers.ValidatePassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	// Generate and store OTP in Redis
	otp := helpers.GenerateOTP()
	err = connections.RedisClient.Set(ctx, user.Email, otp, 5*time.Minute).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store OTP"})
		return
	}

	// Log or print the OTP for manual entry
	// In production, you should not log sensitive data.
	c.JSON(http.StatusOK, gin.H{
		"message": "OTP generated and stored. Please enter it manually for verification.",
		"otp":     otp, // Remove this line in production to prevent exposing the OTP.
	})

}
