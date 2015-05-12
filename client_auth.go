package vnc

import (
	"net"
)

const (
	secTypeInvalid = iota
	secTypeNone
	secTypeVNCAuth
)

// A ClientAuth implements a method of authenticating with a remote server.
type ClientAuth interface {
	// SecurityType returns the byte identifier sent by the server to
	// identify this authentication scheme.
	SecurityType() uint8

	// Handshake is called when the authentication handshake should be
	// performed, as part of the general RFB handshake. (see 7.1.2)
	Handshake(net.Conn) error
}

// ClientAuthNone is the "none" authentication. See 7.1.2
type ClientAuthNone struct{}

func (*ClientAuthNone) SecurityType() uint8 {
	return secTypeNone
}

func (*ClientAuthNone) Handshake(net.Conn) error {
	return nil
}
