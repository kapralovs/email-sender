package examples

import (
	"fmt"
	"os"
	"strings"

	"github.com/kapralovs/email-sender/internal/email"
)

func TestSendMail(mailSubject, errMsg string) error {
	messageBody := "Some test message:\n\n"
	messageBody += fmt.Sprintf("Тело ответа: \n%s\n", errMsg)
	toEmails := os.Getenv("EMAIL_RECIPIENTS")
	recipients := strings.Split(toEmails, ",")
	if messageBody != "" {
		if err := email.Send(recipients, mailSubject, messageBody); err != nil {
			return err
		}
	}
	return nil
}
