package functions

import (
	"encoding/json"
	"github.com/go-gomail/gomail"
	"log"
	"os"
	"sync"
)

type SMTPServerConfig struct {
	SMTPHost string `json:"smtpHost"`
	SMTPPort int    `json:"smtpPort"`
	SMTPUser string `json:"smtpUser"`
	SMTPPass string `json:"smtpPass"`
}

// WaitGroup is a struct that waits for a collection of goroutines to finish
var wg sync.WaitGroup

var dialer *gomail.Dialer

var initialized bool = false

// InitMail initializes the mailer
func InitMail(SMTPServerConfigFile string) {
	// Load the SMTP server configuration
	SMTPConfig, err := loadSMTPServerConfig(SMTPServerConfigFile)
	if err != nil {
		log.Printf("Error loading SMTP server configuration: %v\n", err)
		return
	}
	dialer = gomail.NewDialer(SMTPConfig.SMTPHost, SMTPConfig.SMTPPort, SMTPConfig.SMTPUser, SMTPConfig.SMTPPass)
	initialized = true
}

// LoadSMTPServerConfig loads the SMTP server configuration from a json file.
// The file should contain the following fields:
// - smtpHost: the SMTP server host (e.g. smtp.gmail.com) (string)
// - smtpPort: the SMTP server port (e.g. 587) (int)
// - smtpUser: the SMTP server username (e.g. examplemail@gmail.com) (string)
// - smtpPass: the SMTP server password (e.g. password123) (string)
func loadSMTPServerConfig(SMTPServerConfigFile string) (SMTPServerConfig, error) {
	content, err := os.ReadFile(SMTPServerConfigFile)
	if err != nil {
		return SMTPServerConfig{}, err
	}

	var config SMTPServerConfig
	err = json.Unmarshal(content, &config)
	if err != nil {
		return SMTPServerConfig{}, err
	}

	return config, nil
}

// SendMail sends an email to the specified address
func SendMail(to string, subject string, content string) {
	if !initialized {
		log.Println("Mail Service not initialized")
		return
	}

	wg.Add(1)
	go sendMail(to, subject, content, &wg)
	wg.Wait()
}

// sendMail sends an email to the specified address
// This function is used by SendMail to send the email in a goroutine
func sendMail(to string, subject string, content string, wg *sync.WaitGroup) {
	defer wg.Done()

	m := gomail.NewMessage()
	m.SetHeader("From", dialer.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content) // Envoi d'email en HTML

	if err := dialer.DialAndSend(m); err != nil {
		log.Printf("Error sending mail at %s: %v\n", to, err)
	} else {
		log.Printf("Mail sent to %s\n", to)
	}
}
