package algorithms

import (
    "crypto/hmac"
    "encoding/base64"
    "errors"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/sha3"
)

// HMACSHA3SigningMethod represents the HMAC SHA-3 signing method
type HMACSHA3SigningMethod struct {
    Name string
}

// Alg returns the name of the signing method
func (m *HMACSHA3SigningMethod) Alg() string {
    return m.Name
}

// Sign signs the token using the HMAC SHA-3 algorithm
func (m *HMACSHA3SigningMethod) Sign(signingString string, key interface{}) (string, error) {
    var keyBytes []byte

    switch k := key.(type) {
    case []byte:
        keyBytes = k
    case string:
        keyBytes = []byte(k)
    default:
        return "", errors.New("invalid key type")
    }

    mac := hmac.New(sha3.New256, keyBytes)
    mac.Write([]byte(signingString))
    return base64.RawURLEncoding.EncodeToString(mac.Sum(nil)), nil
}

// Verify verifies the token using the HMAC SHA-3 algorithm
func (m *HMACSHA3SigningMethod) Verify(signingString, signature string, key interface{}) error {
    var keyBytes []byte

    switch k := key.(type) {
    case []byte:
        keyBytes = k
    case string:
        keyBytes = []byte(k)
    default:
        return errors.New("invalid key type")
    }

    sig, err := base64.RawURLEncoding.DecodeString(signature)
    if err != nil {
        return err
    }

    mac := hmac.New(sha3.New256, keyBytes)
    mac.Write([]byte(signingString))
    if !hmac.Equal(sig, mac.Sum(nil)) {
        return errors.New("signature is invalid")
    }

    return nil
}

// Initialize the custom HMAC SHA-3 signing method
func init() {
    jwt.RegisterSigningMethod("HS3", func() jwt.SigningMethod {
        return &HMACSHA3SigningMethod{Name: "HS3"}
    })
}

// createHMACSHA3Token creates a JWT signed with HMAC SHA-3
func createHMACSHA3Token(secret []byte) (string, error) {
    token := jwt.New(jwt.GetSigningMethod("HS3"))
    claims := jwt.MapClaims{
        "foo": "bar",
    }
    token.Claims = claims
    return token.SignedString(secret)
}

// verifyHMACSHA3Token verifies a JWT signed with HMAC SHA-3
func verifyHMACSHA3Token(tokenString string, secret []byte) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if token.Method.Alg() != "HS3" {
            return nil, jwt.ErrSignatureInvalid
        }
        return secret, nil
    })
}
