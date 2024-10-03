package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCategory = "categories"
)

type Category struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" binding:"required" form:"name"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by"`
}

type CategoryRepository interface {
	Create(c context.Context, category *Category) error
	FetchAll(c context.Context) ([]Category, error)
	FetchByName(c context.Context, name string) (Category, error)
	Update(c context.Context, category *Category) error
	Delete(c context.Context, id string) error
}

type CategoryUsecase interface {
	Create(c context.Context, category *Category) error
	FetchAll(c context.Context) ([]Category, error)
	FetchByName(c context.Context, name string) (Category, error)
	Update(c context.Context, category *Category) error
	Delete(c context.Context, id string) error
}
