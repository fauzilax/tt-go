package handler

import "tt-go/features/likereview"

type InsertRequest struct {
	IDReview uint `json:"id_review" form:"id_review"`
	IDMember uint `json:"id_member" form:"id_member"`
}

func ReqToCore(data interface{}) *likereview.Core {
	res := likereview.Core{}

	switch data.(type) {
	case InsertRequest:
		cnv := data.(InsertRequest)
		res.IDReview = cnv.IDReview
		res.IDMember = cnv.IDMember

	default:
		return nil
	}

	return &res
}
