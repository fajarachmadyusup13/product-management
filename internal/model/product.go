package model

import (
	"context"
	"time"
)

type (
	SearchProductSortType  int
	SearchProductOrderType int
)

type Product struct {
	ID          int64
	Name        string
	Price       float64
	Description string
	Quantity    int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const (
	SortProductByCreatedAt = SearchProductSortType(0)
	SortProductByPrice     = SearchProductSortType(1)
	SortProductByName      = SearchProductSortType(2)

	OrderAsc  = SearchProductOrderType(1)
	OrderDesc = SearchProductOrderType(2)
)

type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	GetAll(ctx context.Context, sortType SearchProductSortType, orderType SearchProductOrderType) ([]*Product, error)
}

type ProductUsecase interface {
	CreateProduct(ctx context.Context, product *Product) error
	GetAllProduct(ctx context.Context, sortType SearchProductSortType, orderType SearchProductOrderType) ([]*Product, error)
}

func ParseSorterToModel(sortTypeStr, orderTypeStr string) (sortType SearchProductSortType, orderType SearchProductOrderType) {
	switch orderTypeStr {
	case "asc":
		orderType = OrderAsc
	case "desc":
		orderType = OrderDesc
	}

	switch sortTypeStr {
	case "createdAt":
		sortType = SortProductByCreatedAt
	case "name":
		sortType = SortProductByName
	case "price":
		sortType = SortProductByPrice
	}

	return sortType, orderType
}

func ParseSorterToString(sortType SearchProductSortType, orderType SearchProductOrderType) (sortTypeStr, orderTypeStr string) {
	switch orderType {
	case OrderAsc:
		orderTypeStr = "asc"
	case OrderDesc:
		orderTypeStr = "desc"
	}

	switch sortType {
	case SortProductByCreatedAt:
		sortTypeStr = "createdAt"
	case SortProductByName:
		sortTypeStr = "name"
	case SortProductByPrice:
		sortTypeStr = "price"
	}

	return sortTypeStr, orderTypeStr
}
