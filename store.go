package main

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
)

type Store struct {
	user    *User
	session map[uuid.UUID]*webauthn.SessionData
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
	// TODO: implement
	testId, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	return s.session[testId]
}

func (s *Store) SaveSession(session *webauthn.SessionData) {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	s.session[id] = session
}
