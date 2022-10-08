package helper

import (
	"cloud_disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/smtp"
	"net/textproto"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

func MD5(s string) string { return fmt.Sprintf("%x", md5.Sum([]byte(s))) }

func GenerateToken(id int, identity, name string) (string, error) {
	//id
	//identity
	//name
	uc := define.UserClaim{
		Id:       uint64(id),
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//MailSendCode
//邮箱验证发送
func MailSendCode(mail, code string) error {
	e := &email.Email{
		To:      []string{"mail"},
		From:    "get <leeclong@163.com>",
		Subject: "test",
		Text:    []byte("Test"),
		HTML:    []byte("你的验证码为：<h1>" + code + "</h1>"),
		Headers: textproto.MIMEHeader{},
	}
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "leeclong@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}

func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLenght; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func UUID() string {
	return uuid.NewV4().String()
}
