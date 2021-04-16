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

var DBport = "3306"

var DBuser = "root"

var DBpwd = "12345"

var StripeAPIKey = "pk_test_51IdOGzIc2dfDJNEgxjsDGTWzPiq2SjJyRkKeqRXiYQoytDe1Ye8QrTJSyoHtLMyaFKKbxeAIqxjxxR3n8alBjfGM00xiDA7cDq"

var StripeSecretKey = "sk_test_51IdOGzIc2dfDJNEgcotpLXZQcC9WtMhubRf3KO4q6oQ8gzY3hrDwJizcL8tWsTRjuBhmpbiLIHNYCiqU9wLLGBtH008jb3bIbc"

var EApiKey = "819a3f4cae29c3e36d30a078f4e440e6"
var EAppId = "f3669381"

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

var FoodAPI = "https://api.edamam.com/api/food-database/v2/parser"