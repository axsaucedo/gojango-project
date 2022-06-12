package main

import (
	gomail "github.com/ainsleyclark/go-mail"
	"github.com/robfig/cron"
	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
)

type All struct {
	Scheduler *cron.Cron
	Mailer    mail.Encryption
	Premailer *premailer.Premailer
	Conf      *gomail.Mailer
}
