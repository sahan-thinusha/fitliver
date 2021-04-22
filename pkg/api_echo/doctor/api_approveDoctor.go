package doctor

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/doctor"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func ApproveDoctor(c echo.Context) (*model.Doctor, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	approve,_ := strconv.ParseBool(c.QueryParam("approve"))
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	status := ""
	if approve{
		status = env.STATUS_APPROVED
	}else{
		status = env.STATUS_REJECTED
	}

	result,err := op.ApproveDoctor(id, status,claims.Auth)
	return result, err

}