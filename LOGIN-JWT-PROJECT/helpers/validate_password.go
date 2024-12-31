package helpers

import (
	"fmt"
	"regexp"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	upper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	lower := regexp.MustCompile(`[a-z]`).MatchString(password)
	number := regexp.MustCompile(`[0-9]`).MatchString(password)
	special := regexp.MustCompile(`[!@#$%^&*()]`).MatchString(password)
	if !upper || !lower || !number || !special {
		return fmt.Errorf("password must contain uppercase, lowercase, number, and special character")
	}
	return nil
}
