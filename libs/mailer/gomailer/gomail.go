package gomailer

import (
	"GoTenancy/libs/mailer"
	"github.com/go-gomail/gomail"
)

// Sender gomail struct
type Sender struct {
	*Config
}

// Config gomail config
type Config struct {
	Sender gomail.Sender
}

// New initalize gomail sender with gomail.Dailer
func New(config *Config) *Sender {
	if config == nil {
		config = &Config{}
	}

	return &Sender{Config: config}
}

// Send send email with GoMail
func (sender *Sender) Send(email mailer.Email) error {
	var (
		to, cc, bcc []string
		m           = gomail.NewMessage()
	)

	if email.From != nil {
		m.SetHeader("From", email.From.String())
	}

	if email.ReplyTo != nil {
		m.SetHeader("ReplyTo", email.ReplyTo.String())
	}

	for _, address := range email.TO {
		to = append(to, address.String())
	}

	if len(to) > 0 {
		m.SetHeader("To", to...)
	}

	for _, address := range email.CC {
		cc = append(cc, address.String())
	}

	if len(cc) > 0 {
		m.SetHeader("Cc", cc...)
	}

	for _, address := range email.BCC {
		bcc = append(bcc, address.String())
	}

	if len(bcc) > 0 {
		m.SetHeader("Bcc", bcc...)
	}

	if email.Headers != nil {
		m.SetHeaders(email.Headers)
	}

	m.SetHeader("Subject", email.Subject)

	if email.Text != "" {
		m.AddAlternative("text/plain", email.Text)
	}

	if email.HTML != "" {
		m.AddAlternative("text/html", email.HTML)
	}

	for _, attachment := range email.Attachments {
		if attachment.Inline {
			m.Embed(attachment.FileName)
		} else {
			m.Attach(attachment.FileName)
		}
	}

	return gomail.Send(sender.Sender, m)
}
