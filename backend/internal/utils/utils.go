package utils

import (
	"crypto/rand"
	"html"
)

func SanitizeInput(input string) string {
	return html.EscapeString(input)
}

func GenerateCodes() (uint32, error) {
	n := make([]byte, 8)
	if _, err := rand.Read(n); err != nil {
		return 0, err
	}
	// Convert to uint32 and get last 5 digits
	code := uint32(n[0])<<24 | uint32(n[1])<<16 | uint32(n[2])<<8 | uint32(n[3])
	code = code%90000 + 10000 // Ensures number is between 10000 and 99999
	return code, nil
}
