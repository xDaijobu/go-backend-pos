package domain

import (
	"context"
)

type LogoutUsecase interface {
	DeleteAccessToken(c context.Context, accessToken string, userID string) (isSuccess bool, err error)
}
