package helpers

import "fmt"

// Mock function to send OTP via email (implement using an email service)
func SendOTP(email, otp string) error {
	// Logic to send OTP (use an email service like SendGrid)
	fmt.Printf("Sending OTP %s to email %s\n", otp, email)
	return nil
}
