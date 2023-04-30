package data

import (
	"time"
	"tt-go/features/likereview"

	"gorm.io/gorm"
)

type LikeReview struct {
	IDReview  uint
	IDMember  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func DataToCore(data LikeReview) likereview.Core {
	return likereview.Core{
		IDReview: data.IDReview,
		IDMember: data.IDMember,
	}
}

func CoreToData(core likereview.Core) LikeReview {
	return LikeReview{
		IDReview: core.IDReview,
		IDMember: core.IDMember,
	}
}
