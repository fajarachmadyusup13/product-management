package model

import "context"

type Product struct {
	ID          int64
	Name        string
	Price       float64
	Description string
	Quantity    int64
}

type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	GetAll(ctx context.Context) ([]*Product, error)
}

type ProductUsecase interface {
	CreateProduct(ctx context.Context, product *Product) error
	GetAllProduct(ctx context.Context) ([]*Product, error)
}
