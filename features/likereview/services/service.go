package services

import (
	"errors"
	"strings"
	"tt-go/features/likereview"
)

type likereviewUseCase struct {
	qry likereview.LikeReviewService
}

func New(md likereview.LikeReviewData) likereview.LikeReviewService {
	return &likereviewUseCase{
		qry: md,
	}
}

// Insert implements likereview.LikeReviewService
func (luc *likereviewUseCase) Insert(idReview uint, idMember uint) (likereview.Core, error) {
	res, err := luc.qry.Insert(idReview, idMember)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "username already registered"
		} else {
			msg = "server error"
		}
		return likereview.Core{}, errors.New(msg)
	}

	return res, nil
}

// Delete implements likereview.LikeReviewService
func (luc *likereviewUseCase) Delete(idReview uint, idMember uint) error {
	err := luc.qry.Delete(idReview, idMember)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "username already registered"
		} else {
			msg = "server error"
		}
		return errors.New(msg)
	}

	return nil
}
