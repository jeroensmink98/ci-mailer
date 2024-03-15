# CI-Mailer

CI-Mailer is a command-line tool designed for sending emails directly from the command line. It's particularly useful in CI/CD pipelines to notify team members of build statuses, deployments, or any other automated email notifications.

## Features

- Send emails with custom `From`, `To`, `Subject`, and `Body`.
- Supports HTML and text email bodies.
- Configurable SMTP server settings, including server, port, user, and password.
- TLS support with an option to bypass TLS certificate verification (useful for internal or self-signed certificates).

## Getting Started

To run the program use the following

```bash
ci-mailer \
    --from "your_email@mail.com" \
    --to "someone@mail.com" \
    --subject "A new deployment" \
    --body "path_to_html_file" \
    --smtp-server "your_smtp.com" \
    --smtp-port "587" \
    --smtp-user="username@mail.com" \
    --smtp-password="your_secure_password"
```

### Installation

To install CI-Mailer, run the following command in your terminal:

```bash
go get github.com/jeroensmink98/ci-mailer
```
