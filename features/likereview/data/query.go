package data

import (
	"errors"
	"log"
	"tt-go/features/likereview"

	"gorm.io/gorm"
)

type likereviewQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) likereview.LikeReviewData {
	return &likereviewQuery{
		db: db,
	}
}

// Insert implements likereview.LikeReviewData
func (lq *likereviewQuery) Insert(idReview uint, idMember uint) (likereview.Core, error) {
	data := LikeReview{}
	data.IDReview = idReview
	data.IDMember = idMember
	err := lq.db.Create(&data).Error
	if err != nil {
		log.Println("query error", err.Error())
		return likereview.Core{}, errors.New("server error")
	}
	res := DataToCore(data)
	return res, nil
}

// Delete implements likereview.LikeReviewData
func (lq *likereviewQuery) Delete(idReview uint, idMember uint) error {
	qryDelete := lq.db.Delete(&LikeReview{}, "id_review = ? && id_member = ?", idReview, idMember)
	affRow := qryDelete.RowsAffected
	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete user content, data not found")
	}
	return nil
}
