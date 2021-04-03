package other

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/bmi"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

func CalculateBMI(c echo.Context) (*model.BMI, error) {
	w := c.QueryParam("weight")
	h := c.QueryParam("height")

	weight := 0.0
	if w != "" {
		weight, _ = strconv.ParseFloat(strings.TrimSpace(w), 64)
	}
	height := 0.0
	if h != "" {
		height, _ = strconv.ParseFloat(strings.TrimSpace(h), 64)
	}
	result,err := op.CalculateBMI(weight,height)
	return result, err

}

