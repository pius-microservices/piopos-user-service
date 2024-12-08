package utils

import (
	"strconv"

	"github.com/pius-microservices/piopos-user-service/config"
	"github.com/pius-microservices/piopos-user-service/package/database/models"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	Text    string
	Name    string
	Subject string
	OTPCode string
}

func SendMail(user *models.User, data *EmailData) error {
	envCfg := config.LoadConfig()

	// Create email message
	m := gomail.NewMessage()
	m.SetHeader("From", "PioPos Team <"+envCfg.MailerEmail+">")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", data.Text)

	// SMTP configuration
	port, _ := strconv.Atoi(envCfg.MailerPort)
	d := gomail.NewDialer(envCfg.MailerHost, port, envCfg.MailerEmail, envCfg.MailerPassword)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func EmailHTML(header, text1, text2, text3, footerText, year string) string {
	return `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<style>
					body {
						font-family: Arial, sans-serif;
						margin: 0;
						padding: 0;
						background-color: #f4f4f4;
					}
					.container {
						max-width: 600px;
						margin: 0 auto;
						padding: 20px;
						background-color: #ffffff;
						border-radius: 8px;
						box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
					}
					.header {
						background-color: #86c7ff;
						padding: 10px;
						text-align: center;
						color: #ffffff;
						border-top-left-radius: 8px;
						border-top-right-radius: 8px;
					}
					.content {
						padding: 20px;
						color: #333333;
						text-align: center;
					}
					.footer {
						text-align: center;
						padding: 10px;
						color: #777777;
						font-size: 12px;
					}
					a {
						text-decoration: none;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<div class="header">
						<h1>` + header + `</h1>
					</div>
					<div class="content">
						<p>` + text1 + `</p>
						<p>` + text2 + `</p>
						<h1>` + text3 + `</h1>
					</div>
					<div class="footer">
						<p>` + footerText + `</p>
						<p>  ` + year + ` <a href="#">PIO POS</a>. All rights reserved.</p>
					</div>
				</div>
			</body>
			</html>
		`
}
