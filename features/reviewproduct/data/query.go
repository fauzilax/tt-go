package data

import (
	"errors"
	"log"
	"tt-go/features/reviewproduct"

	"gorm.io/gorm"
)

type reviewproductQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) reviewproduct.ReviewProductData {
	return &reviewproductQuery{
		db: db,
	}
}

// SelectProductID implements reviewproduct.ReviewProductData
func (rq *reviewproductQuery) SelectProductID(id_product uint) (interface{}, error) {
	data := []Output{}
	err := rq.db.Raw("select review_products.id_review, username, gender, skin_type, skin_color, desc_review,count(like_reviews.id_member) as JUMLAH_LIKE_REVIEW from review_products left join members on members.id_member = review_products.id_member left join like_reviews on like_reviews.id_review = review_products.id_review where id_product= ? group by review_products.id_review", id_product).Scan(&data).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []reviewproduct.Core{}, errors.New("server error")
	}
	res := []Output{}
	for i, val := range data {
		res = append(res, val)
		temp := []LikeReview{}
		err := rq.db.Where("id_review = ?", val.IDReview).Find(&temp).Error
		if err != nil {
			log.Println("query error", err.Error())
			return []reviewproduct.Core{}, errors.New("server error")
		}
		res[i].JumlahLikeReview = len(temp)
	}
	outputRes := []OutputRes{}
	for _, val := range res {
		outputRes = append(outputRes, ResponseRes(val))
	}

	return outputRes, nil
}
