package main

import (
	"errors"
	"time"
)

type User struct {
	Id        string
	UserName  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type UserService struct {
	users map[string]*User
}

func NewUserService() *UserService {
	return &UserService{users: make(map[string]*User)}
}

func (s *UserService) CreateUser(username, email, password string) (*User, error) {
	id := generateID()
	if _, exists := s.users[id]; exists {
		return nil, errors.New("user already exists")
	}

	user := &User{
		Id:        id,
		UserName:  username,
		Email:     email,
		Password:  hashPassword(password),
		CreatedAt: time.Now(),
	}

	s.users[id] = user
	return user, nil
}

func (s *UserService) AuthenticateUser(username, password string) (string, error) {
	for _, user := range s.users {
		if user.UserName == username && user.Password == hashPassword(password) {
			return user.Id, nil
		}
	}

	return "", errors.New("invalid credentials")
}

func (s *UserService) GetUser(userID string) (*User, error) {
	user, exixts := s.users[userID]
	if !exixts {
		return nil, errors.New("user not found")
	}

	return user, nil
}
