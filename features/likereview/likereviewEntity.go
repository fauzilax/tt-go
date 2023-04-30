package likereview

import "github.com/labstack/echo/v4"

type Core struct {
	// ID       uint
	IDReview uint `json:"id_review"`
	IDMember uint `json:"id_member"`
}

type LikeReviewHandler interface {
	Insert() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type LikeReviewService interface {
	Insert(idReview uint, idMember uint) (Core, error)
	Delete(idReview uint, idMember uint) error
}

type LikeReviewData interface {
	Insert(idReview uint, idMember uint) (Core, error)
	Delete(idReview uint, idMember uint) error
}
