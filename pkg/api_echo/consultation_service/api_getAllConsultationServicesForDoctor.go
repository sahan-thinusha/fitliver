package consultation_service




import (
	"fitliver/pkg/env"
	_ "fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/consultation_service"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GetAllConsultationServicesForDoctor(c echo.Context) ([]*model.ConsultationService, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)
	result,err := op.ConsultationServiceForDoctor(claims.Sub)
	return result, err
}
