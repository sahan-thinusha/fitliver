package healthrecord

import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/healthrecord"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetHealthRecordByPatient(c echo.Context) ([]*model.HealthRecord, error) {
	id:= c.Param("id")
	patientId,_ := strconv.ParseInt(id, 10, 64)
	result,err := op.GetHealthRecordByPatient(patientId)
	return result, err
}