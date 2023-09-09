package httpsvc

import (
	"net/http"

	"github.com/fajarachmadyusup13/product-management/internal/model"
	"github.com/labstack/echo/v4"
)

type HTTPService struct {
	productUsecase model.ProductUsecase
}

func NewHTTPService() *HTTPService {
	return new(HTTPService)
}

func (h *HTTPService) SetProductUsecase(u model.ProductUsecase) {
	h.productUsecase = u
}

func (h *HTTPService) InitRoutes(route *echo.Echo) {
	productGroup := route.Group("/product")
	productGroup.GET("/ping/", h.PingHandler)

	productGroup.POST("/create/", h.CreateProductHandler)
	productGroup.GET("/search/", h.SearchAllProductsHandler)

}

func (h *HTTPService) PingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "PONG")
}
