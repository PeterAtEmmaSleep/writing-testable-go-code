package main

import (
	"errors"
	"net/http"
)

//go:generate mockgen -source=$GOFILE -destination=authorize_mock.go -package=$GOPACKAGE
type AuthService interface {
	authorize(req *http.Request) error
}

type authService struct {
	expectedUsername string
}

func NewAuthService(expectedUsername string) AuthService {
	return &authService{
		expectedUsername: expectedUsername,
	}
}

func (s *authService) authorize(req *http.Request) error {
	username := req.Header.Get("username")
	if username == s.expectedUsername {
		return nil
	}
	return errors.New("invalid credentials")
}
