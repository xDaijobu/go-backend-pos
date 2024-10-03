package usecase

import (
	"context"
	"time"

	"go-backend-pos/domain"
	"go-backend-pos/internal/tokenutil"
)

type loginUsecase struct {
	userRepository  domain.UserRepository
	tokenRepository domain.TokenRepository
	contextTimeout  time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, tokenRepository domain.TokenRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		contextTimeout:  timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(c context.Context, user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return lu.tokenRepository.CreateToken(c, user, secret, expiry, false)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
