package usecase

import (
	"context"
	"time"

	"go-backend-pos/domain"
)

type logoutUsecase struct {
	tokenRepository domain.TokenRepository
	contextTimeout  time.Duration
}

func NewLogoutUsecase(tokenRepository domain.TokenRepository, timeout time.Duration) domain.LogoutUsecase {
	return &logoutUsecase{
		tokenRepository: tokenRepository,
		contextTimeout:  timeout,
	}
}

func (tu *logoutUsecase) DeleteAccessToken(c context.Context, accessToken string, userID string) (isSuccess bool, err error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	err = tu.tokenRepository.InvalidateToken(ctx, accessToken, userID)
	if err != nil {
		return false, err
	}

	return true, nil
}
