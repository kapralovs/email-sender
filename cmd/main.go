package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kapralovs/email-sender/internal/examples"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//fromEmail := os.Args[1]
	err = examples.TestSendMail(os.Getenv("EMAIL_SUBJECT"), "test message body")
	if err != nil {
		log.Fatal(err)
	}
}
