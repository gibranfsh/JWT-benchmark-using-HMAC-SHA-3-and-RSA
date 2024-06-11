// File: benchmark.go
// Berikut adalah program untuk benchmarking pembuatan dan verifikasi token JWT dengan algoritma RSA dan HMAC SHA-3.

package main

import (
	"crypto/rsa"
	"fmt"
	"time"
)

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

	fmt.Println("JWT hasil generasi menggunakan RSA: ", tokenString+"\n")

	start = time.Now()
	for i := 0; i < 1000; i++ {
		verifyRSAToken(tokenString, publicKey)
	}
	verificationTime := time.Since(start)

	fmt.Println("1000x JWT Creation using RSA Time:", creationTime)
	fmt.Println("1000x JWT Verification using RSA Time:", verificationTime)
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

	fmt.Println("\nJWT hasil generasi menggunakan HMAC SHA-3: ", tokenString+"\n")

	start = time.Now()
	for i := 0; i < 1000; i++ {
		verifyHMACSHA3Token(tokenString, hmacSecret)
	}
	verificationTime := time.Since(start)

	fmt.Println("1000x JWT Creation using HMAC SHA-3 Time:", creationTime)
	fmt.Println("1000x JWT Verification using HMAC SHA-3 Time:", verificationTime)
}