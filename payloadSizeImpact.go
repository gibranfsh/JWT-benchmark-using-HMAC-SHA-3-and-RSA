package main

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func payloadSizeImpact(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, hmacSecret []byte) {
	sizes := []int{100, 1000, 10000}
	for _, size := range sizes {
		payload := generatePayload(size)
		// RSA
		start := time.Now()
		createCustomRSAToken(privateKey, payload)
		fmt.Printf("RSA token creation time with payload size %d: %s\n", size, time.Since(start))

		tokenString, _ := createCustomRSAToken(privateKey, payload)
		start = time.Now()
		verifyRSAToken(tokenString, publicKey)
		fmt.Printf("RSA token verification time with payload size %d: %s\n", size, time.Since(start))

		// HMAC SHA-3
		start = time.Now()
		createCustomHMACSHA3Token(hmacSecret, payload)
		fmt.Printf("HMAC SHA-3 token creation time with payload size %d: %s\n", size, time.Since(start))

		tokenString, _ = createCustomHMACSHA3Token(hmacSecret, payload)
		start = time.Now()
		verifyHMACSHA3Token(tokenString, hmacSecret)
		fmt.Printf("HMAC SHA-3 token verification time with payload size %d: %s\n", size, time.Since(start))
	}
}

// Function to generate a payload of specified size
func generatePayload(size int) map[string]interface{} {
	payload := make(map[string]interface{})
	for i := 0; i < size; i++ {
		payload[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}
	return payload
}

// Function to create RSA token with custom payload
func createCustomRSAToken(privateKey *rsa.PrivateKey, payload map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims(payload))
	tokenString, err := token.SignedString(privateKey)
	return tokenString, err
}

// Function to create HMAC SHA-3 token with custom payload
func createCustomHMACSHA3Token(secret []byte, payload map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}
