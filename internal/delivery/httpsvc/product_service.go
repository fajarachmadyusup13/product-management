package httpsvc

import (
	"encoding/json"
	"net/http"

	"github.com/fajarachmadyusup13/product-management/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *HTTPService) CreateProductHandler(c echo.Context) error {

	product := new(model.Product)

	err := c.Bind(product)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	ctx := c.Request().Context()

	err = h.productUsecase.CreateProduct(ctx, product)
	if err != nil {
		logrus.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, "created")
}

func (h *HTTPService) SearchAllProductsHandler(c echo.Context) error {
	sortType, orderType := model.ParseSorterToModel(c.QueryParam("sortType"), c.QueryParam("orderType"))

	ctx := c.Request().Context()

	res, err := h.productUsecase.GetAllProduct(ctx, sortType, orderType)
	if err != nil {
		logrus.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	j, err := json.Marshal(res)
	if err != nil {
		logrus.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, string(j))
}
