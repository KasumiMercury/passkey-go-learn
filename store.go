package main

import "github.com/go-webauthn/webauthn/webauthn"

type Store struct {
	user    *User
	session *webauthn.SessionData
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) GetUser() *User {
	return s.user
}

func (s *Store) SaveUser(user *User) {
	s.user = user
}

func (s *Store) GetSession() *webauthn.SessionData {
	return s.session
}

func (s *Store) SaveSession(session *webauthn.SessionData) {
	s.session = session
}
