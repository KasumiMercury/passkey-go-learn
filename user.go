package main

import (
	"encoding/binary"
	"github.com/go-webauthn/webauthn/webauthn"
)

type User struct {
	id          uint64
	name        string
	displayName string
	credentials []webauthn.Credential
}

func (u *User) WebAuthnID() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.BigEndian.PutUint64(buf, u.id)
	return buf
}

func (u *User) WebAuthnName() string {
	return u.name
}

func (u *User) WebAuthnDisplayName() string {
	return u.displayName
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.credentials
}
