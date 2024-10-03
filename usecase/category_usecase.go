package usecase

import (
	"context"
	"go-backend-pos/domain"
	"time"
)

type categoryUsecase struct {
	categoryRepository domain.CategoryRepository
	contextTimeout     time.Duration
}

func NewCategoryUsecase(categoryRepository domain.CategoryRepository, timeout time.Duration) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryRepository: categoryRepository,
		contextTimeout:     timeout,
	}
}

func (cu categoryUsecase) Create(c context.Context, category *domain.Category) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepository.Create(ctx, category)
}

func (cu categoryUsecase) FetchAll(c context.Context) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepository.FetchAll(ctx)
}

func (cu categoryUsecase) FetchByName(c context.Context, name string) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()

	return cu.categoryRepository.FetchByName(ctx, name)
}

func (cu categoryUsecase) Update(c context.Context, category *domain.Category) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()

	return cu.categoryRepository.Update(ctx, category)
}

func (cu categoryUsecase) Delete(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()

	return cu.categoryRepository.Delete(ctx, id)
}
