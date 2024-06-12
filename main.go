// File: main.go
// Berikut adalah program utama yang menjalankan semua program uji yang telah diimplementasikan.

package main

import (
	"fmt"
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
	fmt.Println("=== Benchmark JWT (RSA) ===")
	benchmarkRSA(privateKey, publicKey)

	// Benchmarking HMAC SHA-3
	fmt.Println("=== Benchmark JWT (HMAC SHA-3) ===")
	benchmarkHMACSHA3(hmacSecret)

	fmt.Println("")

	// Integrity check for RSA
	fmt.Println("=== Integrity Check JWT (RSA) ===")
	integrityCheckRSA(publicKey)

	fmt.Println("")

	// Integrity check for HMAC SHA-3
	fmt.Println("=== Integrity Check JWT (HMAC SHA-3) ===")
	integrityCheckHMAC()

	fmt.Println("")

	// Test algorithm confusion
	fmt.Println("=== Algorithm Confusion Test ===")
	algorithmConfusion(privateKey, publicKey, hmacSecret)

	fmt.Println("")

	// Test payload size impact
	fmt.Println("=== Payload Size Impact Test ===")
	payloadSizeImpact(privateKey, publicKey, hmacSecret)

	fmt.Println("")

	// Stress testing
	fmt.Println("=== Stress Testing ===")
	stressTesting(privateKey, publicKey, hmacSecret)
}
