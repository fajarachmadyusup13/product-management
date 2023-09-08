package repository

import (
	"context"

	"github.com/fajarachmadyusup13/product-management.git/internal/model"
)

type productRepository struct {
}

func NewProductRepository() model.ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Create(ctx context.Context, product *model.Product) error {
	return nil
}

func (r *productRepository) GetAll(ctx context.Context) ([]*model.Product, error) {
	return nil, nil
}
