package tokenutil

import (
	"fmt"
	"go-backend-pos/domain"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))

	claims := &domain.JwtCustomClaims{
		Name:      user.Name,
		UserID:    user.ID.Hex(),
		SessionID: uuid.New().String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	claimsRefresh := &domain.JwtCustomRefreshClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err
}

func ExtractClaimsFromToken(requestToken string, secret string) (*domain.JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(requestToken, &domain.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*domain.JwtCustomClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

//func ExtractClaimsFromToken(requestToken string, secret string) (string, error) {
//	claims, err := ExtractDataFromToken(requestToken, secret)
//	if err != nil {
//		return "", err
//	}
//
//	return claims.ID, nil
//}
//
//func ExtractSessionIDFromToken(requestToken string, secret string) (string, error) {
//	claims, err := ExtractDataFromToken(requestToken, secret)
//	if err != nil {
//		return "", err
//	}
//
//	return claims.SessionID, nil
//}
