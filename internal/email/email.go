package email

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/sirupsen/logrus"
)

func Send(toEmail []string, subj string, body string) error {
	fromEmail := os.Getenv("EMAIL_SENDER")
	from := mail.Address{Address: fromEmail}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["Subject"] = subj

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	serverName := os.Getenv("SMTP_SERVER")

	host, _, err := net.SplitHostPort(serverName)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	emailAppKey := os.Getenv("EMAIL_APP_KEY")
	auth := smtp.PlainAuth("", from.Address, emailAppKey, host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", serverName, tlsconfig)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	// Auth
	if err = client.Auth(auth); err != nil {
		logrus.Errorln(err)
		return err
	}

	// To && From
	if err = client.Mail(from.Address); err != nil {
		logrus.Errorln(err)
		return err
	}
	for _, addr := range toEmail {
		if err = client.Rcpt(addr); err != nil {
			logrus.Errorln(err)
			return err
		}
	}
	// Data
	w, err := client.Data()
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	err = w.Close()
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	err = client.Quit()
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}
