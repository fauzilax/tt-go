package data

import (
	"errors"
	"log"
	"tt-go/features/member"

	"gorm.io/gorm"
)

type memberQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) member.MemberData {
	return &memberQuery{
		db: db,
	}
}

// Insert implements member.MemberData
func (mq *memberQuery) Insert(newData member.Core) (member.Core, error) {
	cnv := CoreToData(newData)
	err := mq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return member.Core{}, errors.New("server error")
	}
	res := DataToCore(cnv)
	return res, nil
}

// Update implements member.MemberData
func (mq *memberQuery) Update(id uint, updData member.Core) (member.Core, error) {
	cnv := CoreToData(updData)
	// log.Println(cnv)
	qry := mq.db.Where("id_member = ?", id).Updates(&cnv)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return member.Core{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return member.Core{}, errors.New("member not found")
	}

	res := DataToCore(cnv)
	res.IDMember = id
	return res, nil
}

// SelectAllMember implements member.MemberData
func (mq *memberQuery) SelectAllMember() ([]member.Core, error) {
	cnv := []Member{}
	err := mq.db.Find(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []member.Core{}, errors.New("server error")
	}

	res := []member.Core{}
	for _, val := range cnv {
		res = append(res, DataToCore(val))
	}
	return res, nil
}

// Delete implements member.MemberData
func (mq *memberQuery) Delete(memberID uint) error {
	qryDelete := mq.db.Delete(&Member{}, memberID)
	affRow := qryDelete.RowsAffected
	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete user content, data not found")
	}

	return nil
}
