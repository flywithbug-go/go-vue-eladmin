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
	if !MailVerify(to) {
		return fmt.Errorf("mail not right")
	}
	m := gomail.NewMessage()
	m.SetAddressHeader("From", from, title)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if Mail == nil {
		return errors.New("mail is nil")
	}
	return Mail.DialAndSend(m)
}

func SendVerifyMail(title, verifyStr, mail string) error {
	if !MailVerify(mail) {
		return fmt.Errorf("mail not right")
	}
	return sendMail(mail, title, "邮箱验证", verifyStr, config.Conf().MailConfig.Username)
}

var routerRe = regexp.MustCompile(`^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$`)

func MailVerify(mail string) bool {
	match := routerRe.FindString(mail)
	return strings.EqualFold(match, mail)
}

func SendVerifyCode(title, code, mail string) error {
	if !MailVerify(mail) {
		return fmt.Errorf("mail not right")
	}
	body := fmt.Sprintf("您的验证码是： %s ", code)
	return sendMail(mail, title, "邮箱验证", body, config.Conf().MailConfig.Username)
}
