package handler

import (
	"net/http"
	"strconv"
	"tt-go/features/member"

	"github.com/labstack/echo/v4"
)

type memberController struct {
	srv member.MemberService
}

func New(ms member.MemberService) member.MemberHandler {
	return &memberController{
		srv: ms,
	}
}

// Insert implements member.MemberHandler
func (mc *memberController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := InsertRequest{}
		err := c.Bind(&data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		res, err := mc.srv.Insert(*ReqToCore(data))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"status":  "Created",
			"message": "Success",
			"data":    res,
		})
	}
}

// Update implements member.MemberHandler
func (mc *memberController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_memberStr := c.Param("id_member")
		id_member, _ := strconv.Atoi(id_memberStr)
		data := UpdateRequest{}
		err := c.Bind(&data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		res, err := mc.srv.Update(uint(id_member), *ReqToCore(data))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "Success",
			"data":    res,
		})
	}
}

// SelectAllMember implements member.MemberHandler
func (mc *memberController) SelectAllMember() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := mc.srv.SelectAllMember()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "Success",
			"data":    res,
		})
	}
}

// Delete implements member.MemberHandler
func (mc *memberController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_memberStr := c.Param("id_member")
		id_member, _ := strconv.Atoi(id_memberStr)
		err := mc.srv.Delete(uint(id_member))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "Not Found",
				"message": "Member with ID " + id_memberStr + " Not Found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "Success Deleted",
		})
	}
}
