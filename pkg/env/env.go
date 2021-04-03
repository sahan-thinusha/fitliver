package env

import (
	"github.com/dgrijalva/jwt-go"
	gorm "github.com/jinzhu/gorm"
)

var RDB *gorm.DB

var RestPort = "8080"

var DBDIALET = "mysql"

var DBdb = "fitliver"

var DBhost = "127.0.0.1"

var DBport = "3307"

var DBuser = "root"

var DBpwd = "sahan"



const (
	DOCTOR = "ROLE_DOCTOR"
	PATIENT = "ROLE_PATIENT"
	ADMIN = "ROLE_ADMIN"
)

type JwtCustomClaims struct {
	Sub  string
	Auth string
	jwt.StandardClaims
}