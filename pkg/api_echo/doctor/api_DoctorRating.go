package doctor


import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/doctor"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func AddRating(c echo.Context) (*model.Rating, error) {
	rating := model.Rating{}
	if err := c.Bind(&rating); err != nil {
		return nil, err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result,err := op.AddRating(&rating,claims.Sub)
	return result, err
}

func UpdateRating(c echo.Context) (*model.Rating, error) {
	rating := model.Rating{}
	if err := c.Bind(&rating); err != nil {
		return nil, err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result,err := op.RatingUpdate(&rating,claims.Sub)
	return result, err
}

func GetUserRatingForDoctor(c echo.Context) (*model.Rating, error) {
	id := c.Param("id")
	doctorId,_ := strconv.ParseInt(id, 10, 64)


	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result,err := op.GetUserRatingForADoctor(doctorId,claims.Sub)
	return result, err
}

func GetOverallRatingForDoctor(c echo.Context) (*model.Doctor, error) {
	id := c.Param("id")
	doctorId,_ := strconv.ParseInt(id, 10, 64)

	result,err := op.GetOverallRatingForDoctor(doctorId)
	return result, err
}