package notice

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

var email = Email{
	Host:     viper.GetString("email.host"),
	Port:     viper.GetInt("email.port"),
	Username: viper.GetString("email.username"),
	Password: viper.GetString("email.password"),
	From:     viper.GetString("email.from"),
	SendName: viper.GetString("email.sendname"),
	To:       viper.GetString("email.to"),
}

type Email struct {
	Host     string
	Port     int
	SendName string
	Username string
	Password string
	From     string
	To       string
}

func (e *Email) send(replyTo string, subject string, msg string) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", e.From, e.SendName)
	m.SetHeader("To", e.To)
	m.SetHeader("Subject", subject)
	m.SetHeader("Reply-To", replyTo)
	m.SetBody("text/html", msg)

	dialer := gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)
	if err := dialer.DialAndSend(m); err != nil {
		fmt.Println(err.Error())
	}

}
