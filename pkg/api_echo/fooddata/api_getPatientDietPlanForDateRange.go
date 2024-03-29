package fooddata


import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/diet_plan"
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)
const (Time = "02-01-2006")
func GetPatientDietPlanForDateRange(c echo.Context) ([]*model.DietPlan, error) {
	id:= c.Param("id")
	patientId,_ := strconv.ParseInt(id, 10, 64)
	fromDate := c.QueryParam("fromdate")
	toDate := c.QueryParam("todate")

	from, _ := time.Parse(Time, fromDate)
fmt.Println(from)
	to, _ := time.Parse(Time, toDate)

	result,err := op.GetPatientDietPlansByDate(patientId,from,to)
	return result, err
}