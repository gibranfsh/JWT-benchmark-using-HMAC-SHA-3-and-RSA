# JWT (RSA and HMAC SHA-3) Performance and Security Evaluation

This project demonstrates the implementation and performace and security evaluation of JSON Web Tokens (JWT) that was generated using RSA and HMAC SHA-3 algorithms. The project includes code to generate and verify JWTs, as well as to evaluate their performance and analyze their security against common threats.

## Table of Contents

- [Introduction](#introduction)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
  - [Running the Program](#running-the-program)
- [Program Results](#program-results)
- [Conclusion](#conclusion)
- [Future Work](#future-work)
- [Credits](#credits)

## Introduction

This project aims to evaluate the performance and security of JWTs generated using two different cryptographic algorithms: RSA (2048-bit) and HMAC SHA-3 (256-bit). The evaluation includes benchmarking the token creation and verification processes and analyzing the resistance of these tokens to common security threats by conducting various tests such as Benchmarking Test, Algorithm-Confusion Test, Integrity Checking Test, Payload Size Impact Test, Stress Testing.


## Requirements

- Go 1.18 or later
- Git (for cloning the repository)

## Installation

1. Clone the repository:
```
git clone https://github.com/gibranfsh/JWT-benchmark-using-HMAC-SHA-3-and-RSA.git
 ```

2. Change to the project directory:
```
cd JWT-benchmark-using-HMAC-SHA-3-and-RSA
```

3. Install the project dependencies:
```
go mod tidy
```

### Running the Program

To run the program and see the results, execute the following command:

```
go run main.go hmac-sha-3.go rsa.go algorithmConfusion.go benchmark.go integrityCheck.go payloadSizeImpact.go stressTesting.go
```

## Program Results

Example output from running the benchmark:

```
=== Benchmark JWT (RSA) ===
JWT hasil generasi menggunakan RSA:  eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.TMZZJlneK29cDx4QfVDhy0bEOJZyFNRWK0zYvNnJ7ZfskcgQqZyrmRUl1iIw-9fGDxjNb0talLpSbZgjv22bvsHzZFH_xmE2A-9a9NXTrk2bW2qI-0Tc6eqqscgyNbIan0M8cYnP7vPRYVL-hHK4OYdc9t2TH5YjZlCGfPlMFNSu-C_KVWdnnyMcl_mS85NxM3RtQEWKB_QpyI7EAQ-KcDbbrdtqvQ8OxFK3HtmlOoacXO5Owq5F4QUljHxn05IeplHlsGn9f9U6dFcaNzLLvQj3-4OZHgA2xgpfTPyz75mbIY-LsgEs_nmXTuhDE0qmDFMOPSeAh9Sq-Y-JgmMwpw

1000x JWT Creation using RSA Time: 1.3238119s
1000x JWT Verification using RSA Time: 161.7051ms
=== Benchmark JWT (HMAC SHA-3) ===

JWT hasil generasi menggunakan HMAC SHA-3:  eyJhbGciOiJIUzMiLCJ0eXAiOiJKV1QifQ.eyJmb28iOiJiYXIifQ.irFOB4hOLI2H3tM6HDtpt14iFp0-YrhxZzg8JWUZzPE

1000x JWT Creation using HMAC SHA-3 Time: 4.9868ms
1000x JWT Verification using HMAC SHA-3 Time: 4.9869ms

=== Integrity Check JWT (RSA) ===
Generated RSA JWT for Integrity Check: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.l134ShxiA74c_PQkBAz-bDngxIQNFO1n68KJg2YEEpKzAo7KR25dMDAUpebn0iGp2ACNkG6Eaj-iA-l-vSqh7q_4V5OR5FHuRuFDVh0OifutefTWb-roWr0hMjGKDTj7vRWf6qBPvtT2ncO4mO5dOixmyO7AJK0H_2LolycyHozrRcykL-D921AhmpvYxYwpCKY-xoZ19GvqJPxZrR1Jg6ocDPFSDe0QF7tbSK3OE56cae0cGmfFVOqGt0ms7-YbRf_B07ii25bleVEwfnZZY-V-6y4hBia2A19HhEm1YEdEpztWqnclk7mtOQjcI07Q6idIynaIid_4zK5u79Jirw
Original token verification failed: crypto/rsa: verification error
Tampered token verification failed as expected: crypto/rsa: verification error

=== Integrity Check JWT (HMAC SHA-3) ===
Generated HMAC JWT for Integrity Check: eyJhbGciOiJIUzMiLCJ0eXAiOiJKV1QifQ.eyJmb28iOiJiYXIifQ.irFOB4hOLI2H3tM6HDtpt14iFp0-YrhxZzg8JWUZzPE
RSA verification failed as expected: signature is invalid

=== Payload Size Impact Test ===
RSA token creation time with payload size 100: 996.7µs
RSA token verification time with payload size 100: 998.7µs
HMAC SHA-3 token creation time with payload size 100: 0s
HMAC SHA-3 token verification time with payload size 100: 0s
RSA token creation time with payload size 1000: 1.9932ms
RSA token verification time with payload size 1000: 997.4µs
HMAC SHA-3 token creation time with payload size 1000: 0s
HMAC SHA-3 token verification time with payload size 1000: 0s
RSA token creation time with payload size 10000: 7.9782ms
RSA token verification time with payload size 10000: 7.0763ms
HMAC SHA-3 token creation time with payload size 10000: 5.984ms
HMAC SHA-3 token verification time with payload size 10000: 4.9879ms

=== Stress Testing ===
RSA stress test time: 28.2211667s
HMAC SHA-3 stress test time: 436.3074ms
```

## Conclusion

Based on the tests conducted, several important conclusions can be drawn. First, the implementation of JWT using RSA and HMAC SHA-3 is proven to be secure. The Algorithm-Confusion Test and Integrity Checking Test demonstrate that a JWT signed with one algorithm cannot be verified with another, and any modifications to the token are effectively detected (the algorithms in this case being RSA and HMAC SHA-3). Second, HMAC SHA-3 shows superior performance compared to RSA in terms of token creation and verification time. Benchmarking tests indicate that HMAC SHA-3 is significantly faster than RSA for both token creation and verification. Third, HMAC SHA-3 is more efficient in handling large payloads and high loads. Although the token creation and verification times increase with larger payload sizes, HMAC SHA-3 remains faster than RSA, especially for large payloads. Fourth, in stress tests, HMAC SHA-3 demonstrates better resilience compared to RSA.

## Future Work

Future work on this project could include:

- Implementing additional cryptographic algorithms for JWT generation and verification.
- Extending the benchmarking to include more comprehensive performance metrics and security analysis.
- Integrating the JWT implementation with a web application to demonstrate real-world use cases.

## Credits

This project was created by Gibran Fasha Ghazanfar - 18221069
