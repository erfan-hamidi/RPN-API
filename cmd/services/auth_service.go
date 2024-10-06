package services

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("l#x0l&a=9r=2w87uz(*42)jxu_+l%bdce^%92qa=+@$$zfa3$5")

// GenerateJWT generates a new JWT token
func GenerateJWT(username string) (string, error) {
    // Create the claims
    claims := jwt.MapClaims{
        "authorized": true,
        "username":    username,
        "exp":        time.Now().Add(time.Hour * 24).Unix(), // Token expiration
    }

    // Create the token using the claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token with the secret key
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
