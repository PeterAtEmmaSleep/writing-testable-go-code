package main

import "github.com/pkg/errors"

//go:generate mockgen -source=$GOFILE -destination=validate_mock.go -package=$GOPACKAGE
type Validator interface {
	validate(appRequest *AppRequest) error
}

type validator struct{}

func NewValidator() Validator {
	return &validator{}
}

func (v *validator) validate(appRequest *AppRequest) error {
	if appRequest.someField == "" {
		return errors.New("missing required field 'someField'")
	}
	return nil
}
