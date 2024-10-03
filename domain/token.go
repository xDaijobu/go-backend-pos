package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionToken = "tokens"
)

type Token struct {
	ID        primitive.ObjectID `bson:"_id" json:"-"`
	UserID    primitive.ObjectID `bson:"userID" binding:"required" json:"-"`
	Token     string             `bson:"token" form:"token" binding:"required" json:"title"`
	Blacklist bool               `bson:"blacklist" form:"blacklist" json:"blacklist"`
	Expiry    primitive.DateTime `bson:"expiry" form:"expiry" json:"expiry"`
	CreatedAt primitive.DateTime `bson:"createdAt" binding:"required" json:"-"`
}

type TokenRepository interface {
	CreateToken(c context.Context, user *User, accessTokenSecret string, expiry int, blacklist bool) (string, error)
	InvalidateToken(c context.Context, accessToken string, userId string) error
	FetchByUserID(c context.Context, userID string) ([]Token, error)
	FetchByToken(c context.Context, accessToken string) (Token, error)
	FetchByBlacklist(c context.Context, blacklist bool) ([]Token, error)
	//FetchByBlacklistAndUserID(c context.Context, blacklist bool, userID string) ([]Token, error)
}
