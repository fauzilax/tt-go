package handler

import "tt-go/features/member"

type InsertRequest struct {
	Username  string `json:"username" form:"username"`
	Gender    string `json:"gender" form:"gender"`
	SkinType  string `json:"skin_type" form:"skin_type"`
	SkinColor string `json:"skin_color" form:"skin_color"`
}

type UpdateRequest struct {
	Username  string `json:"username" form:"username"`
	Gender    string `json:"gender" form:"gender"`
	SkinType  string `json:"skin_type" form:"skin_type"`
	SkinColor string `json:"skin_color" form:"skin_color"`
}

func ReqToCore(data interface{}) *member.Core {
	res := member.Core{}

	switch data.(type) {
	case InsertRequest:
		cnv := data.(InsertRequest)
		res.Username = cnv.Username
		res.Gender = cnv.Gender
		res.SkinType = cnv.SkinType
		res.SkinColor = cnv.SkinColor
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Username = cnv.Username
		res.Gender = cnv.Gender
		res.SkinType = cnv.SkinType
		res.SkinColor = cnv.SkinColor
	default:
		return nil
	}

	return &res
}
