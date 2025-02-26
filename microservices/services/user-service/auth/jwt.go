package auth

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("your-secret-key")

func GenerateToken(userID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    return token.SignedString(jwtKey)
}
