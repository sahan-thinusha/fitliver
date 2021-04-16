package healthrecord



import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/healthrecord"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func UpdateHealthRecord(c echo.Context) (*model.HealthRecord, error) {
	healthRecord := model.HealthRecord{}
	if err := c.Bind(&healthRecord); err != nil {
		return nil, err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result,err := op.UpdateHealthRecord(&healthRecord,claims.Sub)
	return result, err
}
