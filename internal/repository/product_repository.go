package repository

import (
	"context"
	"fmt"

	"github.com/fajarachmadyusup13/product-management/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) model.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Create(ctx context.Context, product *model.Product) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"product": product,
	})

	tx := r.db.WithContext(ctx).Begin()

	err := r.db.Create(&product).Error
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (r *productRepository) GetAll(ctx context.Context, sortType model.SearchProductSortType, orderType model.SearchProductOrderType) (products []*model.Product, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"sortType":  sortType,
		"orderType": orderType,
	})
	sortTypeStr, orderTypeStr := model.ParseSorterToString(sortType, orderType)

	err = r.db.WithContext(ctx).Order(fmt.Sprint(sortTypeStr, " ", orderTypeStr)).Find(products).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return products, nil
}
