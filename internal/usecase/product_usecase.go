package usecase

import (
	"context"

	"github.com/fajarachmadyusup13/product-management/internal/model"
	"github.com/sirupsen/logrus"
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
	err := u.productRepo.Create(ctx, product)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":     ctx,
			"product": product,
		}).Error(err)
		return err
	}
	return nil
}

func (u *productUsecase) GetAllProduct(ctx context.Context, sortType model.SearchProductSortType, orderType model.SearchProductOrderType) ([]*model.Product, error) {
	res, err := u.productRepo.GetAll(ctx, sortType, orderType)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":       ctx,
			"sortType":  sortType,
			"orderType": orderType,
		}).Error(err)
		return nil, err
	}

	return res, nil
}
