package main

type Store struct{}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) GetUser() *User {
	// TODO: implement
	return nil
}
