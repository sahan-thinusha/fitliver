package patient

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/patient"
	"github.com/labstack/echo/v4"
	"io"
	"os"
	"strconv"
	"time"
)

func CreatePatient(c echo.Context) (*model.Patient, error) {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	dob := c.FormValue("dob")
	gender := c.FormValue("gender")

	user := model.User{}
	user.Role = env.PATIENT
	user.Email = email
	user.Password = password
	user.Name = name

	patient := model.Patient{}
	patient.User = user
	patient.Name = name
	patient.DateOfBirth = dob
	patient.Gender = gender

	file, err := c.FormFile("profile_pic")
	if err != nil {
		return nil,err
	}
	src, err := file.Open()
	if err != nil {
		return nil,err
	}
	defer src.Close()

	folderPath := "patient_images"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		_ = os.Mkdir(folderPath, os.ModePerm)
	}

	fileSuffix := strconv.FormatInt(time.Now().UnixNano(), 10)
	dst, err := os.Create(folderPath + "/" + fileSuffix + "_" + file.Filename)
	if err != nil {
		return nil,err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil,err
	}

	patient.ProfilePic = dst.Name()


	result,err := op.PatientRegister(&patient)
	return result, err
}

