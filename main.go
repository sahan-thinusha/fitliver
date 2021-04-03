package main

import (
	logger "fitliver/pkg/logger"
	model "fitliver/pkg/model"
	gorm "github.com/jinzhu/gorm"
)

import (
	"fitliver/pkg/env"
)
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
import (
	"github.com/dgrijalva/jwt-go"
)


func main() {

	database0, err := gorm.Open("mysql", env.DBuser+":"+env.DBpwd+"@tcp("+env.DBhost+":"+env.DBport+")/"+env.DBdb+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logger.Log.Error(err)
	}
	env.RDB = database0

	model.InitModels(database0)

	RunProxy()
}

type JwtCustomClaims struct {
	Sub  string
	Auth []string
	jwt.StandardClaims
}

func RunProxy() {
	flag.Parse()
	defer glog.Flush()
	run()
}

var (
	endpoint = flag.String("endpoint", "localhost:50051", "Your Description")
)

func run() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	r := e.Group("api")
	jwtConfig := middleware.JWTConfig{

		Claims:     &JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(jwtConfig))
	controller_echo.APIControllerServer(r)
	controller_echo.APIControllerDoctor(r)
	controller_echo.APIControllerBlog(r)
	controller_echo.LoginController(r)
	controller_echo.APIControllerCalculation(r)
	e.Logger.Fatal(e.Start(":" + env.RestPort))
}


