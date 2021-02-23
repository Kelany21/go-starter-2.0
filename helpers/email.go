package helpers

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
)

/**
* function to send email with subject and content
 */
func SendMail(email string, subject string, content string) {
	go func() {
		m := gomail.NewMessage()
		m.SetHeader("From", os.Getenv("STMP_EMAIL_SENDER"))
		m.SetHeader("To", email)
		m.SetHeader("Subject", os.Getenv("APP_NAME")+" "+subject)
		m.SetBody("text/html", content)

		// Send the email to Bob
		d := gomail.NewPlainDialer(os.Getenv("STMP_EMAIL_HOST"), 587, os.Getenv("STMP_EMAIL_ADDRESS"), os.Getenv("STMP_EMAIL_PASSWORD"))
		fmt.Println("++++++++++++")
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
		fmt.Println("++++++++++++")
	}()
	return
}
