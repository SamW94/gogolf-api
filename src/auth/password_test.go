package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword_Success(t *testing.T) {
	password := "MySecurePassword123!"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(hashed) == 0 {
		t.Fatal("expected non-empty hash string")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		t.Errorf("hash does not match original password: %v", err)
	}
}

func TestCheckPasswordHash_Valid(t *testing.T) {
	password := "CorrectHorseBatteryStaple"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}

	err = CheckPasswordHash(hash, password)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestCheckPasswordHash_Invalid(t *testing.T) {
	password := "RightPassword123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}

	wrongPassword := "WrongPassword123"

	err = CheckPasswordHash(hash, wrongPassword)
	if err == nil {
		t.Error("expected error for invalid password, got nil")
	}
}

func TestCheckPasswordHash_BadHash(t *testing.T) {
	badHash := "not-a-valid-bcrypt-hash"
	password := "somepassword"

	err := CheckPasswordHash(badHash, password)
	if err == nil {
		t.Error("expected error for invalid hash input, got nil")
	}
}
