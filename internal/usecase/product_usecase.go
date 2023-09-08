package usecase

import (
	"context"

	"github.com/fajarachmadyusup13/product-management.git/internal/model"
)

type productUsecase struct {
	productRepo model.ProductRepository
}

func NewProductUsecase(productRepo model.ProductRepository) model.ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (u *productUsecase) CreateProduct(ctx context.Context, product *model.Product) error {
	return nil
}

func (u *productUsecase) GetAllProduct(ctx context.Context) ([]*model.Product, error) {
	return nil, nil
}
