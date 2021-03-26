package main

import (
	model "fitliver/pkg/model"
	logger "fitliver/pkg/logger"
	gorm "github.com/jinzhu/gorm"
	"os"
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

func readEnvs() {

	if val := os.Getenv(env.REST_PORT); val != "" {
		env.RestPort = val
	}

	if val := os.Getenv(env.E3_URL); val != "" {
		env.E3url = val
	}
	if val := os.Getenv(env.E3_DIALET); val != "" {
		env.E3DIALET = val
	}
}

func main() {

	readEnvs()
	env.LoadEnvs()

	database0, err := gorm.Open("mysql", env.E3user+":"+env.E3pwd+"@tcp("+env.E3host+":"+env.E3port+")/"+env.E3db+"?charset=utf8mb4&parseTime=True&loc=Local")
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
	r := e.Group("")
	jwtConfig := middleware.JWTConfig{

		Claims:     &JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(jwtConfig))
	controller_echo.APIControllerEcho(r)
	e.Logger.Fatal(e.Start(":" + env.RestPort))
}


