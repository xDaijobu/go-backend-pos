package usecase

import (
	"context"
	"time"

	"go-backend-pos/domain"
	"go-backend-pos/internal/tokenutil"
)

type refreshTokenUsecase struct {
	userRepository  domain.UserRepository
	tokenRepository domain.TokenRepository
	contextTimeout  time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, tokenRepository domain.TokenRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		contextTimeout:  timeout,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, rtu.contextTimeout)
	defer cancel()
	return rtu.userRepository.GetByID(ctx, email)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(c context.Context, user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return rtu.tokenRepository.CreateToken(c, user, secret, expiry, false)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	claims, err := tokenutil.ExtractClaimsFromToken(requestToken, secret)
	if err != nil {
		return "", err
	}
	return claims.UserID, nil
}
