package doctor

import (
	"fitliver/pkg/env"
	model "fitliver/pkg/model"
	op "fitliver/pkg/operation/doctor"
	"io"
	"os"
	"strconv"
	"time"
)

import (
	"github.com/labstack/echo/v4"
)

func CreateDoctor(c echo.Context) (*model.Doctor, error) {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	dob := c.FormValue("dob")
	gender := c.FormValue("gender")
	contactNo := c.FormValue("contactno")
	address := c.FormValue("address")
	specialization := c.FormValue("specialization")
	hospitals := c.FormValue("hospitals")

	user := model.User{}
	user.Role = env.DOCTOR
	user.Email = email
	user.Password = password
	user.Name = name

	doctor := model.Doctor{}
	doctor.User = user
	doctor.Name = name
	doctor.DateOfBirth = dob
	doctor.Gender = gender
	contact := model.ContactNo{}
	contact.ContactNo = contactNo
	doctor.ContactNo = append([]*model.ContactNo{},&contact)
	doctor.Address = address
	doctor.Specialization = specialization
	doctor.Status = env.STATUS_NEW



	file, err := c.FormFile("profile_pic")
	if err==nil{
		src, err := file.Open()
		if err != nil {
			return nil,err
		}
		defer src.Close()


		folderPath := "doctor_images"
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

		doctor.ProfilePic = dst.Name()

	}

	result,err := op.DoctorRegister(&doctor,hospitals)
	return result, err
}
