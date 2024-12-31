package helpers

import (
	"fmt"
	"math/rand"
)

// Helper function to generate a random OTP
func GenerateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
