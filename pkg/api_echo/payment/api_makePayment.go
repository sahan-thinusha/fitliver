package payment

import (
	"fitliver/pkg/env"
	_ "fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/payment"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func MakePayment(c echo.Context) (*model.Payment, error) {
	token := c.QueryParam("token")
	payment := model.Payment{}
	if error := c.Bind(&payment); error != nil {
		return nil, error
	}

	fmt.Println(payment)
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result,err := op.ProceedPayment(&payment,token,claims.Sub)
	return result, err
}


