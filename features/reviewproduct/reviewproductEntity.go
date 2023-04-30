package reviewproduct

import "github.com/labstack/echo/v4"

type Core struct {
	IDReview   uint
	IDProduct  uint
	IDMember   uint
	DescReview string
}

type ReviewProductHandler interface {
	SelectProductID() echo.HandlerFunc
}

type ReviewProductService interface {
	SelectProductID(id_product uint) (interface{}, error)
}

type ReviewProductData interface {
	SelectProductID(id_product uint) (interface{}, error)
}
