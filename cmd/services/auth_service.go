package services

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("l#x0l&a=9r=2w87uz(*42)jxu_+l%bdce^%92qa=+@$$zfa3$5")

type JwtCustomClaims struct {
	Username  string `json:"username"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token
func GenerateJWT(username string) (string, error) {
    // Create the claims
    claims := &JwtCustomClaims{
		username,
		false,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

    // Create the token using the claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token with the secret key
    tokenString, err := token.SignedString(JwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
