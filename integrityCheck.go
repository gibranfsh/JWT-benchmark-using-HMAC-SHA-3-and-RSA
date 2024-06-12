// File: integrityCheck.go
// Berikut adalah program untuk melakukan integrity check pada token JWT dengan algoritma RSA dan HMAC SHA-3.

package main

import (
	"crypto/rsa"
	"fmt"
)

// Function for integrity check with RSA
func integrityCheckRSA(publicKey *rsa.PublicKey) {
	privateKey, _, err := generateRSAKeys()
	if err != nil {
		panic(err)
	}

	// Create a valid token
	tokenString, err := createRSAToken(privateKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated RSA JWT for Integrity Check:", tokenString)

	// Tamper the token
	tamperedToken := tokenString + "tampered"

	// Verify the original token
	_, err = verifyRSAToken(tokenString, publicKey)
	if err != nil {
		fmt.Println("Original token verification failed:", err)
	} else {
		fmt.Println("Original token verification succeeded")
	}

	// Verify the tampered token
	_, err = verifyRSAToken(tamperedToken, publicKey)
	if err != nil {
		fmt.Println("Tampered token verification failed as expected:", err)
	} else {
		fmt.Println("Tampered token verification succeeded unexpectedly")
	}
}

// Function for integrity check with HMAC
func integrityCheckHMAC() {
	// Create a valid token
	secret := []byte("secret")
	tokenString, err := createHMACSHA3Token(secret)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated HMAC JWT for Integrity Check:", tokenString)

	// Tamper the token
	tamperedToken := tokenString + "tampered"

	// Verify the original token
	_, err = verifyHMACSHA3Token(tokenString, secret)
	if err != nil {
		fmt.Println("Original token verification failed:", err)
	} else {
		fmt.Println("Original token verification succeeded")
	}

	// Verify the tampered token
	_, err = verifyHMACSHA3Token(tamperedToken, secret)
	if err != nil {
		fmt.Println("Tampered token verification failed as expected:", err)
	} else {
		fmt.Println("Tampered token verification succeeded unexpectedly")
	}
}

