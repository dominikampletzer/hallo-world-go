// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0
package client

import (
	"github.com/dominikampletzer/hello/mail/smtp"
	"testing"
)

// configure Registry to know which mail implementation is used.
func init() {
	Registry.Add("senderInterface.Sender", new(smtp.MailSenderImpl))
}

func TestMail(t *testing.T) {
	SendMail("dominik@test.de", "Es geht")
}
