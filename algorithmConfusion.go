// File: algorithmConfusion.go
// Berikut adalah program untuk melakukan uji algorithm confusion pada JWT yang dibuat dengan algoritma HMAC SHA-256 dan di-parse dengan algoritma RSA SHA-256.
package main

import (
	"crypto/rsa"
	"fmt"
)

// Function for algorithm confusion test
func algorithmConfusion(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, hmacSecret []byte) {
	fmt.Println("=== Algorithm Confusion Test ===")

	// Create RSA token
	rsaToken, err := createRSAToken(privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Generated RSA Token:", rsaToken)

	// Try to verify RSA token with HMAC SHA-3
	if _, err := verifyHMACSHA3Token(rsaToken, hmacSecret); err != nil {
		fmt.Println("HMAC SHA-3 verification failed as expected:", err)
	} else {
		fmt.Println("HMAC SHA-3 verification should have failed but didn't")
	}

	// Create HMAC SHA-3 token
	hmacToken, err := createHMACSHA3Token(hmacSecret)
	if err != nil {
		panic(err)
	}
	fmt.Println("Generated HMAC SHA-3 Token:", hmacToken)

	// Try to verify HMAC SHA-3 token with RSA
	if _, err := verifyRSAToken(hmacToken, publicKey); err != nil {
		fmt.Println("RSA verification failed as expected:", err)
	} else {
		fmt.Println("RSA verification should have failed but didn't")
	}
}
