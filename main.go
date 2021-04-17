package main

import (
	"crypto/subtle"
	logger "fitliver/pkg/logger"
	model "fitliver/pkg/model"
	gorm "github.com/jinzhu/gorm"
	"os"
)

import (
	"fitliver/pkg/env"
)
import (
	_"github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"

)
import (
	"flag"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
import (
	"fitliver/pkg/controller_echo"
)

func main() {

	database, err := gorm.Open("mysql", env.DBuser+":"+env.DBpwd+"@tcp("+env.DBhost+":"+env.DBport+")/"+env.DBdb+"?charset=utf8mb4&parseTime=True&loc=Local")
	//database, err :=gorm.Open("sqlite3", env.DBdb)
	if err != nil {
		logger.Log.Error(err)
	}
	env.RDB = database

	model.InitModels(database)

	folderPath := "patient_images"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		_ = os.Mkdir(folderPath, os.ModePerm)
	}

	folderPath = "doctor_images"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		_ = os.Mkdir(folderPath, os.ModePerm)
	}

	folderPath = "blog_images"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		_ = os.Mkdir(folderPath, os.ModePerm)
	}
	CreateDefaultUser()
	RunProxy()
}



func RunProxy() {
	flag.Parse()
	defer glog.Flush()
	run()
}

func run() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Static("/patient_images","patient_images")
	e.Static("/blog_images","blog_images")
	e.Static("/doctor_images","doctor_images")

	r := e.Group("/")
	jwtConfig := middleware.JWTConfig{
		Claims:     &env.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}


	r.Use(middleware.JWTWithConfig(jwtConfig))
	controller_echo.APIControllerServer(r)
	controller_echo.APIControllerDoctor(r)
	controller_echo.APIControllerBlog(r)
	controller_echo.APIControllerHospital(r)
	controller_echo.APIControllerPatient(r)
	controller_echo.APIControllerForum(r)
	controller_echo.APIControllerPayment(r)
	controller_echo.APIControllerConsultationService(r)
	controller_echo.APIControllerDietPlanService(r)
	controller_echo.APIControllerHealthRecordService(r)
	controller_echo.APIControllerRating(r)
	u := e.Group("/")
	u.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("fitliver")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("fitliver@123")) == 1 {
			return true, nil
		}
		return false, nil
	}))
	controller_echo.LoginController(u)
	controller_echo.APIControllerHospitalBase(u)
	controller_echo.APIControllerCalculation(u)

	e.Logger.Fatal(e.Start(":" + env.RestPort))
}


func CreateDefaultUser(){
	db := env.RDB
	users := []*model.User{}
	db.Find(&users)
	var targetUser *model.User
	if len(users) == 0 {
		targetUser = &model.User{}
		targetUser.Role = env.ADMIN
		targetUser.Email = "admin"
		targetUser.Password = "admin"
		targetUser.Name = "Admin"
		db.Create(&targetUser)
	}
}
