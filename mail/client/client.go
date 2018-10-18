package client

import (
	"github.com/dominikampletzer/hello/mail/registry"
	"github.com/dominikampletzer/hello/mail/senderInterface"
)

var Registry = registry.NewRegistry()

func SendMail(address, message string) {
	var sender = Registry.Get("senderInterface.Sender").(senderInterface.Sender)
	mailAddress := senderInterface.Address{Address: address}
	sender.SendMail(mailAddress, message)
}
