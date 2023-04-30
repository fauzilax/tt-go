package data

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	IDProduct   uint `gorm:"primaryKey"`
	NameProduct string
	Price       float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
