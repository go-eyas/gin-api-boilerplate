package srv

// 邮件发送

import (
	"github.com/go-eyas/toolkit/email"
)

type EmailSrv struct{}

var Email *email.Email

func EmailInit(conf *EmailConfig) {
	Email = email.New(conf)
}
