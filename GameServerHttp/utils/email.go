package utils

import (
	"net/mail"
	"regexp"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

var emailDialer *gomail.Dialer

// 邮箱校验
func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// 手机校验
func IsPhone(phone string) bool {
	regex := `^1[3-9]\d{9}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(phone)
}

// 初始化邮箱
func SetupEmail() {
	emailDialer = gomail.NewDialer(Conf.Email.Addr, Conf.Email.Port, Conf.Email.UserName, Conf.Email.Password)
	if emailDialer == nil {
		logrus.Fatal("InitEmail is failed")
	}
}

// 发送邮件
func SendEmail(from string, to []string, subject, body string) error {
	if !IsEmail(from) || to == nil || len(to) <= 0 ||
		len(subject) <= 0 || len(body) <= 0 || emailDialer == nil {
		return ErrParameter
	}
	for _, v := range to {
		if !IsEmail(v) {
			return ErrParameter
		}
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if err := emailDialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
