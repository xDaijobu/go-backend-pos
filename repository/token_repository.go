package repository

import (
	"context"
	"errors"
	"fmt"

	"go-backend-pos/domain"
	"go-backend-pos/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoDB "go.mongodb.org/mongo-driver/mongo"
	"time"
)

type tokenRepository struct {
	database   mongo.Database
	collection string
}

func NewTokenRepository(db mongo.Database, collection string) domain.TokenRepository {
	return &tokenRepository{
		database:   db,
		collection: collection,
	}
}

func (t tokenRepository) CreateToken(c context.Context, accessToken string, userId string, blacklist bool) error {
	collection := t.database.Collection(t.collection)

	userIdHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(c, domain.Token{
		ID:        primitive.NewObjectID(),
		Token:     accessToken,
		UserID:    userIdHex,
		Blacklist: blacklist,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	})

	return err
}

func (t tokenRepository) InvalidateToken(c context.Context, accessToken string, userId string) error {
	token, err := t.FetchByToken(c, accessToken)
	if err != nil {
		fmt.Print(err)
		if errors.Is(err, mongoDB.ErrNoDocuments) {
			err = t.CreateToken(c, accessToken, userId, true)
			if err != nil {
				return err
			}

			token, err = t.FetchByToken(c, accessToken)
			if err != nil {
				return err
			}
		}

		return err
	}

	token.Blacklist = true

	_, err = t.database.Collection(t.collection).UpdateOne(c, bson.M{"_id": token.ID}, bson.M{"$set": token})
	return err
}

func (t tokenRepository) FetchByUserID(c context.Context, userID string) ([]domain.Token, error) {
	collection := t.database.Collection(t.collection)

	userIdHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(c, bson.M{"userID": userIdHex})
	if err != nil {
		return nil, err
	}

	var tokens []domain.Token
	err = cursor.All(c, &tokens)
	if tokens == nil {
		return []domain.Token{}, err
	}

	return tokens, err
}

func (t tokenRepository) FetchByToken(c context.Context, accessToken string) (domain.Token, error) {
	collection := t.database.Collection(t.collection)

	var token domain.Token
	err := collection.FindOne(c, bson.M{"token": accessToken}).Decode(&token)

	return token, err
}

func (t tokenRepository) FetchByBlacklist(c context.Context, blacklist bool) ([]domain.Token, error) {
	collection := t.database.Collection(t.collection)

	cursor, err := collection.Find(c, bson.M{"blacklist": blacklist})
	if err != nil {
		return nil, err
	}

	var tokens []domain.Token
	err = cursor.All(c, &tokens)
	if tokens == nil {
		return []domain.Token{}, err
	}

	return tokens, err
}
