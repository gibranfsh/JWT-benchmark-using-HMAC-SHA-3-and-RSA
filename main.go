// File: main.go

package main

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

	// Test integrity check
	integrityCheckRSA(publicKey)
	integrityCheckHMAC()

	// Test expiration handling
	expirationHandling(privateKey, publicKey)
	
	// Test algorithm confusion
	algorithmConfusion(privateKey, publicKey, hmacSecret)

	// Test payload size impact
	payloadSizeImpact(privateKey, publicKey, hmacSecret)

	// Stress testing
	stressTesting(privateKey, publicKey, hmacSecret)
}
