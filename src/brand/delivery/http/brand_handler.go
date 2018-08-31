package http

import (
	"context"
	"net/http"

	"github.com/hardyantz/go-clean-arch/src/brand"
	"github.com/hardyantz/go-clean-arch/model"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// ResponseError brand http response error
type ResponseError struct {
	Message string `json:"message"`
}

// BrandHandler struct for http brand handling
type BrandHandler struct {
	BrandUseCase brand.UseCase
}

// NewBrandHTTPHandler route handling for brand
func NewBrandHTTPHandler(e *echo.Echo, b brand.UseCase) {
	handler := &BrandHandler{
		BrandUseCase: b,
	}
	e.GET("/brand/:id", handler.GetByID)
}

// GetByID rest handler for brand get by ID
func (b *BrandHandler) GetByID(c echo.Context) error {

	idBrand := c.Param("id")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := b.BrandUseCase.GetByID(ctx, idBrand)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case model.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case model.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case model.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
