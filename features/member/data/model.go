package data

import (
	"time"
	"tt-go/features/member"

	"gorm.io/gorm"
)

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

func DataToCore(data Member) member.Core {
	return member.Core{
		IDMember:  data.IDMember,
		Username:  data.Username,
		Gender:    data.Gender,
		SkinType:  data.SkinType,
		SkinColor: data.SkinColor,
	}
}

func CoreToData(core member.Core) Member {
	return Member{
		IDMember:  core.IDMember,
		Username:  core.Username,
		Gender:    core.Gender,
		SkinType:  core.SkinType,
		SkinColor: core.SkinColor,
	}
}
