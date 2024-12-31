package helpers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prathamesh/login_jwt_project/connections"
	"github.com/redis/go-redis/v9"
)

func VerifyOTP(c *gin.Context) {
	type VerifyRequest struct {
		Email string `json:"email" binding:"required"`
		OTP   string `json:"otp" binding:"required"`
	}

	var request VerifyRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve OTP from Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	storedOTP, err := connections.RedisClient.Get(ctx, request.Email).Result()
	if err == redis.Nil || storedOTP != request.OTP {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	// Insert user into database
	collection := connections.GetCollection("your_database_name", "users")
	_, err = collection.InsertOne(ctx, map[string]interface{}{
		"username": request.Email,
		"verified": true,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User verified and registered successfully"})
}
