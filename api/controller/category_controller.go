package controller

import (
	"github.com/gin-gonic/gin"
	"go-backend-pos/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type CategoryController struct {
	CategoryUsecase domain.CategoryUsecase
}

func (tc *CategoryController) Create(c *gin.Context) {
	var category domain.Category
	err := c.ShouldBind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = tc.CategoryUsecase.FetchByName(c, category.Name)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "Category already exists with the given name"})
		return
	}

	category.ID = primitive.NewObjectID()
	category.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	err = tc.CategoryUsecase.Create(c, &category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Category created successfully",
	})
}

func (tc *CategoryController) Fetch(c *gin.Context) {
	categories, err := tc.CategoryUsecase.FetchAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
