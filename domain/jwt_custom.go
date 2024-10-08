package domain

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	Name      string `json:"name"`
	UserID    string `json:"userID"`
	SessionID string `json:"sessionID"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
