package usecase

import (
	"context"
	"testing"

	"github.com/fajarachmadyusup13/product-management/internal/model"
	"github.com/fajarachmadyusup13/product-management/internal/model/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProductUsecase_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	product := &model.Product{
		ID:          123,
		Name:        "Name",
		Price:       10000,
		Description: "Desc",
		Quantity:    12,
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo := mock.NewMockProductRepository(ctrl)
		mockProductRepo.EXPECT().Create(ctx, product).
			Times(1).
			Return(nil)

		productUC := &productUsecase{
			productRepo: mockProductRepo,
		}

		err := productUC.CreateProduct(ctx, product)
		assert.NoError(t, err)
	})
}
