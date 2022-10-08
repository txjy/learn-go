package define

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Id       uint64
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"

var MailPassword = os.Getenv("MailPassword")

//验证码长度
var CodeLenght = 6

//验证码过期时间
var CodeExpire = 300
