package controller_echo



import (
	"fitliver/pkg/api_echo/booking"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateBookingService(c echo.Context) error {
	result, err := booking.CreateBookingService(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}



func GetBookingRequests(c echo.Context) error {
	result, err := booking.GetAllNewBookingRequestsForDoctor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}




func GetAllBookingServicesForDoctor(c echo.Context) error {
	result, err := booking.GetAllBookingsForDoctorByDate(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func GetAllBookingServicesForDoctorById(c echo.Context) error {
	result, err := booking.GetAllBookingsForDoctorById(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func GetAllBookingServicesForPatient(c echo.Context) error {
	result, err := booking.GetAllBookingsForPatientByDate(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}
func BookingServicesApprove(c echo.Context) error {
	result, err := booking.ApproveBookingRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}



func APIControllerBookingService(g *echo.Group) {
	g.POST("api/bookingservice", CreateBookingService)
	g.GET("api/bookingservice/request", GetBookingRequests)
	g.GET("api/bookingservice/doctor", GetAllBookingServicesForDoctor)
	g.GET("api/bookingservice/doctor/:id", GetAllBookingServicesForDoctorById)
	g.GET("api/bookingservice/patient", GetAllBookingServicesForPatient)
	g.POST("api/bookingservice/request/approve/:id", BookingServicesApprove)

}
