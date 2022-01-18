package main

import (
	"net/http"

	"github.com/pkg/errors"
)

type app struct {
	AuthService  AuthService
	Unmarshaller Unmarshaller
	Validator    Validator
	Transformer  Transformer
	ExtApiCall   ExtApiCall
}

func (app *app) handle(w http.ResponseWriter, req *http.Request) {
	_ = app.authorize(req)
	appRequest, _ := app.unmarshall(req)
	_ = app.validate(appRequest)
	extRequest, _ := app.transform(appRequest)
	_ = app.callExtApi(extRequest)
	_ = app.respondWithOK(w)
}

func (app *app) authorize(req *http.Request) error {
	if err := app.AuthService.authorize(req); err != nil {
		return errors.Wrap(err, "request authorization failed")
	}
	return nil
}

func (app *app) unmarshall(httpRequest *http.Request) (*AppRequest, error) {
	req, err := app.Unmarshaller.unmarshall(httpRequest)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshall request")
	}
	return req, nil
}

func (app *app) validate(appRequest *AppRequest) error {
	if err := app.Validator.validate(appRequest); err != nil {
		return errors.Wrap(err, "failed to validate request")
	}
	return nil
}

func (app *app) transform(appRequest *AppRequest) (*ExtRequest, error) {
	extRequest, err := app.Transformer.transform(appRequest)
	if err != nil {
		return nil, errors.Wrap(err, "failed to transform request")
	}
	return extRequest, nil
}

func (app *app) callExtApi(request *ExtRequest) error {
	if err := app.ExtApiCall.CallApi(request); err != nil {
		return errors.Wrap(err, "failed to call external API")
	}
	return nil
}

func (app *app) respondWithOK(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	return nil
}

func (app *app) respondWithInternalServicerError(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusInternalServerError)
	return nil
}
