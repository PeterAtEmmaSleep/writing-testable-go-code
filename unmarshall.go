package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

//go:generate mockgen -source=$GOFILE -destination=unmarshall_mock.go -package=$GOPACKAGE
type Unmarshaller interface {
	unmarshall(httpRequest *http.Request) (*AppRequest, error)
}

type unmarshaller struct{}

func NewUnmarshaller() Unmarshaller {
	return &unmarshaller{}
}

func (s *unmarshaller) unmarshall(httpRequest *http.Request) (*AppRequest, error) {
	var (
		request AppRequest
		body    []byte
	)

	defer func() { _ = httpRequest.Body.Close() }()

	body, err := io.ReadAll(httpRequest.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read request body")
	}

	if err := json.Unmarshal(body, &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshall request body")
	}
	return &request, nil
}
