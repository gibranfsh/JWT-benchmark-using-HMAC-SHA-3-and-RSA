package main

import (
	"crypto/rsa"
	"fmt"
	"time"
)

func main() {
	// Generate RSA keys
	privateKey, publicKey, err := generateRSAKeys()
	if err != nil {
		panic(err)
	}

	// Generate HMAC SHA-3 secret
	hmacSecret := []byte("secret")

	// Benchmarking RSA
	benchmarkRSA(privateKey, publicKey)

	// Benchmarking HMAC SHA-3
	benchmarkHMACSHA3(hmacSecret)
}

// Function to benchmark RSA token creation and verification
func benchmarkRSA(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) {
	// Benchmark token creation
	start := time.Now()
	for i := 0; i < 1000; i++ {
		createRSAToken(privateKey)
	}
	creationTime := time.Since(start)

	// Benchmark token verification
	tokenString, _ := createRSAToken(privateKey)
	start = time.Now()
	for i := 0; i < 1000; i++ {
		verifyRSAToken(tokenString, publicKey)
	}
	verificationTime := time.Since(start)

	fmt.Println("RSA Token Creation Time:", creationTime)
	fmt.Println("RSA Token Verification Time:", verificationTime)
}

// Function to benchmark HMAC SHA-3 token creation and verification
func benchmarkHMACSHA3(hmacSecret []byte) {
	// Benchmark token creation
	start := time.Now()
	for i := 0; i < 1000; i++ {
		createHMACSHA3Token(hmacSecret)
	}
	creationTime := time.Since(start)

	// Benchmark token verification
	tokenString, _ := createHMACSHA3Token(hmacSecret)
	start = time.Now()
	for i := 0; i < 1000; i++ {
		verifyHMACSHA3Token(tokenString, hmacSecret)
	}
	verificationTime := time.Since(start)

	fmt.Println("HMAC SHA-3 Token Creation Time:", creationTime)
	fmt.Println("HMAC SHA-3 Token Verification Time:", verificationTime)
}
