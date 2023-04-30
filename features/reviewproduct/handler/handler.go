package handler

import (
	"net/http"
	"strconv"
	"tt-go/features/reviewproduct"

	"github.com/labstack/echo/v4"
)

type reviewproductController struct {
	srv reviewproduct.ReviewProductService
}

func New(ms reviewproduct.ReviewProductService) reviewproduct.ReviewProductHandler {
	return &reviewproductController{
		srv: ms,
	}
}

// SelectProductID implements reviewproduct.ReviewProductHandler
func (rc *reviewproductController) SelectProductID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_productStr := c.Param("id_product")
		id_product, _ := strconv.Atoi(id_productStr)
		res, err := rc.srv.SelectProductID(uint(id_product))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "Success",
			"data":    res,
		})
	}
}
