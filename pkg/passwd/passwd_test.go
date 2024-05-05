package passwd

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// TestHashPassword tests the HashPassword function
func TestHashPassword(t *testing.T) {
	testPasswords := []string{"password", "password123", "password123!@#"}

	for _, password := range testPasswords {
		hashedPassword, err := HashPassword(password)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	}

	// Test too long password (max 72 characters)
	_, err := HashPassword("123456789012345678901234567890123456789012345678901234567890123456789012345678901")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// TestComparePassword tests the ComparePassword function
func TestComparePassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if !ComparePassword(password, hashedPassword) {
		t.Errorf("Expected true, got false")
	}
	if ComparePassword("wrongpassword", hashedPassword) {
		t.Errorf("Expected false, got true")
	}
}
