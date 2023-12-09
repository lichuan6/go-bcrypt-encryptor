package main

import (
	"fmt"

	bcryptencryptor "github.com/lichuan6/go-bcrypt-encryptor"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Create a new BcryptEncryptor instance
	encryptor := bcryptencryptor.NewBcryptEncryptor(bcrypt.DefaultCost)

	// Encrypting Passwords
	password := "myPassword"
	encryptedPassword, err := encryptor.Encrypt(password)
	if err != nil {
		fmt.Println("Error encrypting password:", err)
		return
	}
	fmt.Println("Encrypted password:", encryptedPassword)

	// Comparing Passwords
	matchErr := encryptor.Compare(encryptedPassword, password)
	if matchErr != nil {
		fmt.Println("Passwords do not match")
	} else {
		fmt.Println("Passwords match")
	}

	// Hash and Password Comparison
	hashErr := encryptor.CompareHashAndPassword(encryptedPassword, password)
	if hashErr != nil {
		fmt.Println("Passwords do not match")
	} else {
		fmt.Println("Passwords match")
	}
}
