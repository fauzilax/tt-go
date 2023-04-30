package main

import (
	"log"
	"tt-go/config"
	likereviewData "tt-go/features/likereview/data"
	likereviewHdl "tt-go/features/likereview/handler"
	likereviewSrv "tt-go/features/likereview/services"
	memberData "tt-go/features/member/data"
	memberHdl "tt-go/features/member/handler"
	memberSrv "tt-go/features/member/services"
	reviewproductData "tt-go/features/reviewproduct/data"
	reviewproductHdl "tt-go/features/reviewproduct/handler"
	reviewproductSrv "tt-go/features/reviewproduct/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	config.Migrate(db)
	mData := memberData.New(db)
	mSrv := memberSrv.New(mData)
	mHdl := memberHdl.New(mSrv)
	lrData := likereviewData.New(db)
	lrSrv := likereviewSrv.New(lrData)
	lrHdl := likereviewHdl.New(lrSrv)
	rpData := reviewproductData.New(db)
	rpSrv := reviewproductSrv.New(rpData)
	rpHdl := reviewproductHdl.New(rpSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/insert", mHdl.Insert())
	e.PUT("/update/:id_member", mHdl.Update())
	e.DELETE("/delete/:id_member", mHdl.Delete())
	e.GET("/selectallmember", mHdl.SelectAllMember())

	e.GET("/selectidproduct/:id_product", rpHdl.SelectProductID())

	e.POST("/like", lrHdl.Insert())
	e.DELETE("/dislike", lrHdl.Delete())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
