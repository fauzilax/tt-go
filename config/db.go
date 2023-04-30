package config

import (
	"fmt"
	"log"

	lrData "tt-go/features/likereview/data"
	mData "tt-go/features/member/data"
	pData "tt-go/features/product/data"

	rpData "tt-go/features/reviewproduct/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dc DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.DBUser, dc.DBPass, dc.DBHost, dc.DBPort, dc.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(mData.Member{})
	db.AutoMigrate(pData.Product{})
	db.AutoMigrate(lrData.LikeReview{})
	db.AutoMigrate(rpData.ReviewProduct{})
}
