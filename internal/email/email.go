package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "vino@gmail.com")
	m.SetHeader("To", "bob@example.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer(
		viper.GetString("MAIL_HOST"),
		viper.GetInt("MAIL_PORT"),
		viper.GetString("MAIL_USERNAME"),
		viper.GetString("MAIL_PASSWORD"),
	)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
