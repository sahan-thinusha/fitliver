package workout



import (
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/workout"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetPatientWorkoutPlans(c echo.Context) ([]*model.WorkoutPlan, error) {
	id:= c.Param("id")
	patientId,_ := strconv.ParseInt(id, 10, 64)
	result,err := op.GetWorkoutPlansForPatient(patientId)
	return result, err
}
