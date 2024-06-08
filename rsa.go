package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

func generateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return nil, nil, err
    }
    return privateKey, &privateKey.PublicKey, nil
}

type RSASigningMethod struct {
    Name string
}

func (m *RSASigningMethod) Alg() string {
    return m.Name
}

func (m *RSASigningMethod) Sign(signingString string, key interface{}) (string, error) {
    var rsaPrivateKey *rsa.PrivateKey

    switch k := key.(type) {
    case *rsa.PrivateKey:
        rsaPrivateKey = k
    default:
        return "", errors.New("invalid key type")
    }

    hash := sha256.New()
    hash.Write([]byte(signingString))
    hashed := hash.Sum(nil)

    signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashed)
    if err != nil {
        return "", err
    }
    return base64.RawURLEncoding.EncodeToString(signature), nil
}

func (m *RSASigningMethod) Verify(signingString, signature string, key interface{}) error {
    var rsaPublicKey *rsa.PublicKey

    switch k := key.(type) {
    case *rsa.PublicKey:
        rsaPublicKey = k
    default:
        return errors.New("invalid key type")
    }

    sig, err := base64.RawURLEncoding.DecodeString(signature)
    if err != nil {
        return err
    }

    hash := sha256.New()
    hash.Write([]byte(signingString))
    hashed := hash.Sum(nil)

    return rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, hashed, sig)
}

func init() {
    jwt.RegisterSigningMethod("RS256", func() jwt.SigningMethod {
        return &RSASigningMethod{Name: "RS256"}
    })
}

func createRSAToken(privateKey *rsa.PrivateKey) (string, error) {
    token := jwt.New(jwt.GetSigningMethod("RS256"))
    claims := jwt.MapClaims{
        "foo": "bar",
    }
    token.Claims = claims
    return token.SignedString(privateKey)
}

func verifyRSAToken(tokenString string, publicKey *rsa.PublicKey) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if token.Method.Alg() != "RS256" {
            return nil, jwt.ErrSignatureInvalid
        }
        return publicKey, nil
    })
}

