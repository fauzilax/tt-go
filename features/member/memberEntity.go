package member

import "github.com/labstack/echo/v4"

type Core struct {
	IDMember  uint   `json:"id_member"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	SkinType  string `json:"skin_type"`
	SkinColor string `json:"skin_color"`
}

type MemberHandler interface {
	Insert() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	SelectAllMember() echo.HandlerFunc
}

type MemberService interface {
	Insert(newData Core) (Core, error)
	Update(id uint, updData Core) (Core, error)
	Delete(memberID uint) error
	SelectAllMember() ([]Core, error)
}

type MemberData interface {
	Insert(newData Core) (Core, error)
	Update(id uint, updData Core) (Core, error)
	Delete(memberID uint) error
	SelectAllMember() ([]Core, error)
}
