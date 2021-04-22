package consultation_service


import (
	"fitliver/pkg/env"
	_ "fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/consultation_service"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func ApproveConsultationRequest(c echo.Context) (*model.ConsultationRequest, error) {
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

	result,err := op.ApproveConsultationRequest(int64(id),status,claims.Sub)
	return result, err
}