package utils

import (
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

func GenerateCode() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(999999))
}

func SendVerificationCode(email, code string) error {
	from := "1366737405@qq.com"
	password := "qpnuswnvkfvxhaeg"
	to := []string{email}
	smtpHost := "smtp.qq.com"
	smtpPort := "465"
	auth := smtp.PlainAuth("", from, password, smtpHost)
	message := []byte("Subject: Verify your email account\n\nYour verification code is: " + code)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}
	return nil
}
