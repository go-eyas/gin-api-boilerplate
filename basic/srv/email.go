package srv

import (
  "github.com/go-eyas/toolkit/email"
)

type EmailSrv struct {}

var Email *email.Email

func EmailInit(conf *EmailConfig) {
  Email = email.New(conf)
}