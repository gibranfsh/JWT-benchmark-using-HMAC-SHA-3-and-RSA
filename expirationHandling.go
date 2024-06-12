// File: expirationHandling.go
// Berikut adalah program untuk mengetes token JWT yang telah expired dengan algoritma RSA.

package main

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Function for expiration handling
func expirationHandling(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) {
	// Create a token with expiration
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	claims := jwt.MapClaims{
		"foo": "bar",
		"exp": time.Now().Add(5 * time.Second).Unix(),
	}
	token.Claims = claims
	tokenString, _ := token.SignedString(privateKey)

	fmt.Println("Generated RSA JWT for Expiration Handling:", tokenString)

	// Wait for the token to expire
	fmt.Println("Waiting for token to expire...")
	time.Sleep(6 * time.Second)

	// Verify the expired token
	_, err := verifyRSAToken(tokenString, publicKey)
	if err != nil {
		fmt.Println("Expired token verification failed as expected:", err)
	} else {
		fmt.Println("Expired token verification succeeded unexpectedly")
	}
}