// File: stressTesting.go
// Berikut adalah program untuk melakukan stress testing dengan membuat dan memverifikasi token secara parallel.

package main

import (
	"crypto/rsa"
	"fmt"
	"time"
)

// Function to test stress testing by creating and verifying tokens in parallel
func stressTesting(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, hmacSecret []byte) {
	numGoroutines := 100
	done := make(chan bool, numGoroutines)

	// Stress test RSA
	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				tokenString, _ := createRSAToken(privateKey)
				verifyRSAToken(tokenString, publicKey)
			}
			done <- true
		}()
	}
	for i := 0; i < numGoroutines; i++ {
		<-done
	}
	fmt.Println("RSA stress test time:", time.Since(start))

	// Stress test HMAC SHA-3
	start = time.Now()
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				tokenString, _ := createHMACSHA3Token(hmacSecret)
				verifyHMACSHA3Token(tokenString, hmacSecret)
			}
			done <- true
		}()
	}
	for i := 0; i < numGoroutines; i++ {
		<-done
	}
	fmt.Println("HMAC SHA-3 stress test time:", time.Since(start))
}