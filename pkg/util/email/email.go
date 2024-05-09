package email

import "net/smtp"

type EmailDispatcher struct {
}

func NewDispatcher() *EmailDispatcher {
	return &EmailDispatcher{}
}

func (e *EmailDispatcher) SendEmail(to []string) error {
	from := "tadilbek11kz@gmail.com"
	password := "77ae910pqsw71239pasx12asuws"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("This is a test email message.")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}
