# JWT Benchmarking with RSA and HMAC SHA-3

This project demonstrates the implementation and benchmarking of JSON Web Tokens (JWT) using RSA and HMAC SHA-3 algorithms. The project includes code to generate and verify JWTs, as well as to benchmark their performance and analyze their security against common threats.

## Table of Contents

- [Introduction](#introduction)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
  - [Running the Benchmark](#running-the-benchmark)
- [Benchmark Results](#benchmark-results)
- [Conclusion](#conclusion)
- [Future Work](#future-work)
- [Credits](#credits)

## Introduction

This project aims to evaluate the performance and security of JWTs generated using two different cryptographic algorithms: RSA (2048-bit) and HMAC SHA-3 (256-bit). The evaluation includes benchmarking the token creation and verification processes and analyzing the resistance of these tokens to common security threats such as brute force attacks, MITM attacks, and token forgery.

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

## Usage

The project includes two main files: `hmac_sha3.go` and `rsa.go`, which contain the implementation of JWTs using HMAC SHA-3 and RSA algorithms, respectively. To run the benchmark, execute the following command:

### Running the Benchmark

To run the benchmarking and see the results, execute the following command:

```
 go run main.go rsa.go hmac-sha-3.go
```

This command will generate and verify JWTs using both RSA and HMAC SHA-3 algorithms, and output the performance results and generated tokens to the console.

## Benchmark Results

Example output from running the benchmark:

```
JWT hasil generasi menggunakan RSA:  eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.JqeG863DiysrcAbset8V1bPk9pAFV9g_LXrB-fNbjI52xrifexcARWCu5hp1h5XfDuQJv_hASLveEMny777GZ0EixAx5DmoNfTfjOQApICLkdVcXJWhI9-eylq3fe0cxrtU9VkdeBGwo484LuPZldnTW2BTwveqVTCgbaf0TVEFeeXYwsh0zkvkeFMJYi1MmsM3DErTroKAmbqdmrPYj6Sj9IqAJ7Zo4Zj1htizauMjM3yump_DbT4to3eFT5aU4rdQo2clAEIbQefTBTWfaw4aBkMY_piYVAxqxw2CdQJv7oV92GB4JMkFzBAhkRdHkx2A6rCUO4qgeHi5Yg9J4Gw

1000x JWT Creation using RSA Time: 1.3292608s
1000x JWT Verification using RSA Time: 155.9339ms

JWT hasil generasi menggunakan HMAC SHA-3:  eyJhbGciOiJIUzMiLCJ0eXAiOiJKV1QifQ.eyJmb28iOiJiYXIifQ.irFOB4hOLI2H3tM6HDtpt14iFp0-YrhxZzg8JWUZzPE

1000x JWT Creation using HMAC SHA-3 Time: 5.0936ms
1000x JWT Verification using HMAC SHA-3 Time: 5.1205ms
```

The benchmark results show the time taken to generate and verify 1000 JWTs using RSA and HMAC SHA-3 algorithms. The results demonstrate the performance difference between the two algorithms, with HMAC SHA-3 being significantly faster than RSA for both token generation and verification.

## Conclusion

The benchmarking results indicate that HMAC SHA-3 is significantly faster than RSA in both token creation and verification. Both algorithms provide robust security, but RSA offers advantages in public key distribution scenarios. The choice between RSA and HMAC SHA-3 should consider specific application requirements, including performance and key management practices.

## Future Work

Future work on this project could include:

- Implementing additional cryptographic algorithms for JWT generation and verification.
- Extending the benchmarking to include more comprehensive performance metrics and security analysis.
- Integrating the JWT implementation with a web application to demonstrate real-world use cases.

## Credits

This project was created by Gibran Fasha Ghazanfar - 18221069
