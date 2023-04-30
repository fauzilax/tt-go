package services

import (
	"errors"
	"strings"
	"tt-go/features/member"
)

type memberUseCase struct {
	qry member.MemberData
}

func New(md member.MemberData) member.MemberService {
	return &memberUseCase{
		qry: md,
	}
}

// Delete implements member.MemberService
func (muc *memberUseCase) Delete(memberID uint) error {
	err := muc.qry.Delete(memberID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "username already registered"
		} else {
			msg = "server error"
		}
		return errors.New(msg)
	}

	return nil
}

// Insert implements member.MemberService
func (muc *memberUseCase) Insert(newData member.Core) (member.Core, error) {
	res, err := muc.qry.Insert(newData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "username already registered"
		} else {
			msg = "server error"
		}
		return member.Core{}, errors.New(msg)
	}

	return res, nil
}

// Update implements member.MemberService
func (muc *memberUseCase) Update(id uint, updData member.Core) (member.Core, error) {
	res, err := muc.qry.Update(id, updData)
	if err != nil {
		return member.Core{}, errors.New("server error")
	}

	return res, nil
}

// SelectAllMember implements member.MemberService
func (muc *memberUseCase) SelectAllMember() ([]member.Core, error) {
	res, err := muc.qry.SelectAllMember()
	if err != nil {
		return []member.Core{}, errors.New("server error")
	}

	return res, nil
}
