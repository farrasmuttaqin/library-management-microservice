package models

import "github.com/golang-jwt/jwt/v4"

type JWTTokenClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
