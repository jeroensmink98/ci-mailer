package main

import (
	"crypto/tls"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	gomail "gopkg.in/mail.v2"
)

func main() {
	var to string
	var from string

	var subject string
	var body string

	var smtpServer string
	var smtpPort int
	var smtpUser string
	var smtpPassword string

	app := &cli.App{
		Name:    "ci-mailer",
		Version: "v1.1.0",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Jeroen Smink",
				Email: "info@devopsfrontier.com",
			},
		},
		Usage:     "sent emails from the command line",
		UsageText: "ci-mailer --from --to --subject --body --smtp-server --smtp-port --smtp-user smtp-password   ",
		Flags: []cli.Flag{
			// format: --from <email>
			&cli.StringFlag{
				Name:        "from",
				Usage:       "from email address",
				Destination: &from,
				Required:    true,
			},

			// format: --to <email>
			&cli.StringFlag{
				Name:        "to",
				Usage:       "to email address",
				Destination: &to,
				Required:    true,
			},

			// format: --subject <email>
			&cli.StringFlag{
				Name:        "subject",
				Usage:       "email subject",
				Destination: &subject,
				Required:    true,
			},

			// format: --body <email>
			&cli.StringFlag{
				Name:        "body",
				Usage:       "reference to HTML file",
				Destination: &body,
				Required:    true,
			},

			// format: --smtp-server <email>
			&cli.StringFlag{
				Name:        "smtp-server",
				Usage:       "smtp server",
				Destination: &smtpServer,
				Required:    true,
			},

			// format: --smtp-port <email>
			&cli.IntFlag{
				Name:        "smtp-port",
				Usage:       "smtp port",
				Destination: &smtpPort,
				DefaultText: "587",
				Required:    true,
			},

			// format: --smtp-user <email>
			&cli.StringFlag{
				Name:        "smtp-user",
				Usage:       "smtp user",
				Destination: &smtpUser,
				Required:    true,
			},

			// format: --smtp-password <email>
			&cli.StringFlag{
				Name:        "smtp-password",
				Usage:       "smtp password",
				Destination: &smtpPassword,
				Required:    true,
			},
		},

		Action: func(c *cli.Context) error {

			// construct a new email
			m := gomail.NewMessage()
			m.SetHeader("From", from)
			m.SetHeader("To", to)
			m.SetHeader("Subject", subject)
			m.SetBody("text/html", readFile(body))
			d := gomail.NewDialer(smtpServer, smtpPort, smtpUser, smtpPassword)

			d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

			err := d.DialAndSend(m)
			check(err)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func readFile(filename string) string {
	b, err := os.ReadFile(filename)
	check(err)
	return string(b)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
