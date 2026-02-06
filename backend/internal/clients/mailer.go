package clients

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// This directive tells Go to include the file in the binary
//
//go:embed expiry.html
var templateFS embed.FS

type PetEntry struct {
	ID      string
	PetName string
	Phone   string
	Date    string
}

type EmailData struct {
	Username   string
	ExpiryType string // "Microchip" or "Rabies"
	Items      []PetEntry
}

func sendGroupedEmail(targetEmail string, data EmailData) error {
	tmpl, err := template.ParseFS(templateFS, "expiry.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	log.Println("TargetEmail:", targetEmail)

	e := email.NewEmail()
	e.From = "Pet Registry <" + SmtpUser + ">"
	e.To = []string{targetEmail}
	e.Subject = fmt.Sprintf("⚠️ %s Expiration Alert", data.ExpiryType)
	e.HTML = body.Bytes()

	auth := smtp.PlainAuth("", SmtpUser, SmtpPass, SmtpServer)

	// Mailpit for testing
	return e.Send(SmtpServer+":"+SmtpPort, auth)
}
