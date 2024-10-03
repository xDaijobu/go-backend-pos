package usecase

import (
	"context"
	"time"

	"go-backend-pos/domain"
	"go-backend-pos/internal/tokenutil"
)

type signupUsecase struct {
	userRepository  domain.UserRepository
	tokenRepository domain.TokenRepository
	contextTimeout  time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, tokenRepository domain.TokenRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		contextTimeout:  timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(c context.Context, user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return su.tokenRepository.CreateToken(c, user, secret, expiry, false)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
