package senderInterface

type Address struct {
	Address string
}
type Sender interface {
	// Send an email to a given address with a  message.
	SendMail(address Address, message string)
}
