package smtp

import (
	"github.com/dominikampletzer/hello/mail/senderInterface"
	"log"
)

type MailSenderImpl struct {
}

func (mail *MailSenderImpl) SendMail(address senderInterface.Address, message string) {
	log.Println("NAchricht wird gesendet an :" + address.Address + "; Nachricht:" + message)

}
