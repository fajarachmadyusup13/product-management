package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fajarachmadyusup13/product-management/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	initializeTest()
	mockDB, mocker := initializeCockroachMockConn()

	product := &model.Product{
		ID:          123,
		Name:        "Name",
		Price:       10000,
		Description: "desc",
		Quantity:    2,
	}

	t.Run("success", func(t *testing.T) {
		mocker.ExpectBegin()
		mocker.ExpectQuery("INSERT INTO \"products\"").
			WithArgs(product.Name, product.Price, product.Description, product.Quantity, sqlmock.AnyArg(), sqlmock.AnyArg(), product.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(product.ID))
		mocker.ExpectCommit()

		repo := NewProductRepository(mockDB)
		err := repo.Create(context.TODO(), product)
		assert.NoError(t, err)
	})

	t.Run("error commit", func(t *testing.T) {
		mocker.ExpectBegin()
		mocker.ExpectQuery("INSERT INTO \"products\"").
			WithArgs(product.Name, product.Price, product.Description, product.Quantity, sqlmock.AnyArg(), sqlmock.AnyArg(), product.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(product.ID))
		mocker.ExpectCommit().WillReturnError(errors.New("error commit"))

		repo := NewProductRepository(mockDB)
		err := repo.Create(context.TODO(), product)
		assert.Error(t, err, errors.New("error commit"))
	})
}

func TestGetAll(t *testing.T) {
	initializeTest()
	mockDB, mocker := initializeCockroachMockConn()

	product := &model.Product{
		ID:          123,
		Name:        "Name",
		Price:       10000,
		Description: "desc",
		Quantity:    2,
	}

	t.Run("success", func(t *testing.T) {
		queryRes := sqlmock.NewRows([]string{"id", "name", "price", "description", "quantity"}).
			AddRow(product.ID, product.Name, product.Price, product.Description, product.Quantity)

		mocker.ExpectQuery("SELECT .+ FROM \"products\"").
			WillReturnRows(queryRes)

		productRepo := NewProductRepository(mockDB)
		res, err := productRepo.GetAll(context.TODO(), model.SortProductByCreatedAt, model.OrderDesc)
		assert.NoError(t, err)
		assert.Equal(t, []*model.Product{product}, res)

	})
}
