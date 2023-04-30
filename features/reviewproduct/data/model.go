package data

import (
	"time"

	"gorm.io/gorm"
)

type ReviewProduct struct {
	IDReview   uint `gorm:"primaryKey"`
	IDProduct  uint
	IDMember   uint
	DescReview string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type Product struct {
	IDProduct   uint `gorm:"primaryKey"`
	NameProduct string
	Price       float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Member struct {
	IDMember  uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Gender    string
	SkinType  string
	SkinColor string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type LikeReview struct {
	IDReview  uint
	IDMember  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Output struct {
	IDReview         uint
	Username         string
	Gender           string
	SkinType         string
	SkinColor        string
	DescReview       string
	JumlahLikeReview int
}
type OutputRes struct {
	Username         string `json:"username"`
	Gender           string `json:"gender"`
	SkinType         string `json:"skin_type"`
	SkinColor        string `json:"skin_color"`
	DescReview       string `json:"desc_review"`
	JumlahLikeReview int    `json:"jumlah_like_review"`
}

func ResponseRes(data Output) OutputRes {
	return OutputRes{
		Username:         data.Username,
		Gender:           data.Gender,
		SkinType:         data.SkinType,
		SkinColor:        data.SkinColor,
		DescReview:       data.DescReview,
		JumlahLikeReview: data.JumlahLikeReview,
	}
}
