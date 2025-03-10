package functions

import (
	"github.com/go-gomail/gomail"
	"sync"
)

// WaitGroup is a struct that waits for a collection of goroutines to finish
var wg sync.WaitGroup

var dialer *gomail.Dialer

// InitMail initializes the mailer
func InitMail(SMTPServerConfigFile string) {
	// Load the SMTP server configuration
	smtpHost, smtpPort, smtpUser, smtpPass := loadSMTPServerConfig(SMTPServerConfigFile)
	dialer = gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
}

// LoadSMTPServerConfig loads the SMTP server configuration from a json file.
// The file should contain the following fields:
// - smtpHost: the SMTP server host
// - smtpPort: the SMTP server port
// - smtpUser: the SMTP server username
// - smtpPass: the SMTP server password
func loadSMTPServerConfig(SMTPServerConfigFile string) (string, int, string, string) {
	return "", 0, "", ""
}

// SendMail sends an email to the specified address
func sendMail(to string, from string, subject string, wg *sync.WaitGroup) {
	defer wg.Done()
}
