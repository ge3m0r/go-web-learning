package mail

import(
	"fmt"
	"golearning/pkg/logger"
	"net/smtp"
	
	emailPKG "github.com/jordan-wright/email"
)

type SMTP struct{}

func (s *SMTP)Send(email Email, config map[string]string) bool{
	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.HTML = email.HTML

	logger.DebugJson("发送邮件", "发件详情", e)

    err := e.Send(
        fmt.Sprintf("%v:%v", config["host"], config["port"]),

        smtp.PlainAuth(
            "",
            config["username"],
            config["password"],
            config["host"],
        ),
    )
    if err != nil {
        logger.ErrorString("发送邮件", "发件出错", err.Error())
        return false
    }

    logger.DebugString("发送邮件", "发件成功", "")
    return true

}