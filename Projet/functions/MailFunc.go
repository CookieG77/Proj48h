package functions

import (
	"encoding/json"
	"github.com/go-gomail/gomail"
	"os"
	"sync"
)

type SMTPServerConfig struct {
	SMTPHost string `json:"smtpHost"`
	SMTPPort int    `json:"smtpPort"`
	SMTPUser string `json:"smtpUser"`
	SMTPPass string `json:"smtpPass"`
}

// wg (WaitGroup) is a struct that waits for a collection of goroutines to finish
var wg sync.WaitGroup

var dialer *gomail.Dialer

// initialized is used to check if the SMTP service has been initialized or not.
var initialized bool = false

// InitMail initializes the mailer.
// If the SMTP server configuration file is not found, the function will log an error and return.
func InitMail(SMTPServerConfigFile string) {
	// Load the SMTP server configuration
	SMTPConfig, err := loadSMTPServerConfig(SMTPServerConfigFile)
	if err != nil {
		WarningPrintf("Could not load the SMTP server configuration file -> %v\n", err)
		return
	}
	dialer = gomail.NewDialer(SMTPConfig.SMTPHost, SMTPConfig.SMTPPort, SMTPConfig.SMTPUser, SMTPConfig.SMTPPass)
	initialized = true
	SuccessPrintf("SMTP server initialized -> %v\n", SMTPConfig)
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

// SendMail sends an email to the specified address.
// If the mailer has not been initialized, the function will log an error and return.
func SendMail(to string, subject string, content string) {
	if !initialized {
		ErrorPrintln("Mail Service not initialized, check the SMTP server configuration file")
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
	m.SetBody("text/html", content)

	if err := dialer.DialAndSend(m); err != nil {
		ErrorPrintf("Could not send mail to %s -> %v\n", to, err)
	} else {
		InfoPrintf("Mail sent to %s\n", to)
	}
}
