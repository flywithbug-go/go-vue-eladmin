package mail

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"trace_go/config"

	"gopkg.in/gomail.v2"
)

var (
	Mail *gomail.Dialer
)

func sendMail(to, title, subject, body, from string) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", from, title)
	m.SetHeader("Subject", subject)
	m.SetHeader("To", to)
	m.SetBody("text/html", body)

	if Mail == nil {
		return errors.New("mail is nil")
	}
	return Mail.DialAndSend(m)
}

func SendVerifyMail(title, verifyStr, mail string) error {
	return sendMail(mail, title, "邮箱验证", verifyStr, config.Conf().MailConfig.Username)
}

var routerRe = regexp.MustCompile(`^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$`)

func MailStringVerify(mail string) bool {
	match := routerRe.FindString(mail)
	return strings.EqualFold(match, mail)
}

func SendVerifyCode(title, code, mail string) error {
	verifyStr := fmt.Sprintf("您的验证码是： <b>%s</b> ", code)
	return sendMail(mail, title, "邮箱验证", verifyStr, config.Conf().MailConfig.Username)
}
