package mail

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendMail(receiver string, subject string, body string) (err error) {
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", viper.GetString("SENDER_NAME"))
	mailer.SetHeader("To", receiver)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		viper.GetString("SMTP_HOST"),
		viper.GetInt("SMTP_PORT"),
		viper.GetString("MAIL_USERNAME"),
		viper.GetString("MAIL_PASSWORD"),
	)

	if err = dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}
