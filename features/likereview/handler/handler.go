package handler

import (
	"net/http"
	"tt-go/features/likereview"

	"github.com/labstack/echo/v4"
)

type likereviewController struct {
	srv likereview.LikeReviewService
}

func New(ms likereview.LikeReviewService) likereview.LikeReviewHandler {
	return &likereviewController{
		srv: ms,
	}
}

// Insert implements likereview.LikeReviewHandler
func (lc *likereviewController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := InsertRequest{}
		err := c.Bind(&data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		res, err := lc.srv.Insert(data.IDReview, data.IDMember)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"status":  "Created",
			"message": "Success",
			"data":    res,
		})
	}
}

// Delete implements likereview.LikeReviewHandler
func (lc *likereviewController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := InsertRequest{}
		err := c.Bind(&data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		err = lc.srv.Delete(data.IDReview, data.IDMember)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "Not Found",
				"message": "Data Not Found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "Success Deleted",
		})
	}
}
