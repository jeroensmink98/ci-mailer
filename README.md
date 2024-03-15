# CI-Mailer

CI-Mailer is a command-line tool designed for sending emails directly from the command line. It's particularly useful in CI/CD pipelines. Pass a HTML file as email body and off you go.

## Features

- Send emails with custom `From`, `To`, `Subject`, and `Body`.
- Supports HTML and text email bodies.
- Configurable SMTP server settings, including server, port, user, and password.

## Getting Started

To run the program use the following

```bash
ci-mailer \
    --from "your_email@mail.com" \
    --to "someone@mail.com" \
    --subject "A new deployment" \
    --body "mail.html" \
    --smtp-server "your_smtp.com" \
    --smtp-port "587" \
    --smtp-user="username@mail.com" \
    --smtp-password="your_secure_password"
```

### Installation

Head over to the [releases page](https://github.com/jeroensmink98/ci-mailer/releases) at GitHub to download the latest release.

## Options

The program has the following options

```bash
USAGE:
   ci-mailer --from --to --subject --body --smtp-server --smtp-port --smtp-user smtp-password   

GLOBAL OPTIONS:
   --from value           from email address
   --to value             to email address
   --subject value        email subject
   --body value           reference to HTML file
   --smtp-server value    smtp server
   --smtp-port value      smtp port (default: 587)
   --smtp-user value      smtp user
   --smtp-password value  smtp password
   --help, -h             show help
   --version, -v          print the version
```
