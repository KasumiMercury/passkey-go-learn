package main

import "github.com/go-webauthn/webauthn/webauthn"

type Store struct{}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) GetUser() *User {
	// TODO: implement
	return nil
}

func (s *Store) SaveUser(user *User) {
	// TODO: implement
}

func (s *Store) GetSession() *webauthn.SessionData {
	// TODO: implement
	return nil
}

func (s *Store) SaveSession(session *webauthn.SessionData) {
	// TODO: implement
}
