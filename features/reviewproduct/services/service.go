package services

import (
	"errors"
	"tt-go/features/reviewproduct"
)

type reviewproductUseCase struct {
	qry reviewproduct.ReviewProductData
}

func New(md reviewproduct.ReviewProductData) reviewproduct.ReviewProductService {
	return &reviewproductUseCase{
		qry: md,
	}
}

// SelectProductID implements reviewproduct.ReviewProductService
func (ruc *reviewproductUseCase) SelectProductID(id_product uint) (interface{}, error) {
	res, err := ruc.qry.SelectProductID(id_product)
	if err != nil {
		return []reviewproduct.Core{}, errors.New("server error")
	}

	return res, nil
}
