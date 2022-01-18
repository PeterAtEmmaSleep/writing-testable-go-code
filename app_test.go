package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	t.Run("should return an error if authorization fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		authService := NewMockAuthService(ctrl)
		app := app{
			AuthService:  authService,
			Unmarshaller: NewMockUnmarshaller(ctrl),
			Validator:    NewMockValidator(ctrl),
			Transformer:  NewMockTransformer(ctrl),
			ExtApiCall:   NewMockExtApiCall(ctrl),
		}

		expectedHttpRequest := &http.Request{}
		authError := errors.New("some error")

		authService.EXPECT().
			authorize(expectedHttpRequest).
			Return(authError).
			Times(1)

		rr := httptest.NewRecorder()
		app.handle(rr, &http.Request{})

		assert.Equal(t, 500, rr.Code)
		assert.Equal(t, "some error message", rr.Body.String())
	})

	t.Run("should return an error if unmarshalling fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		authService := NewMockAuthService(ctrl)
		unmarshaller := NewMockUnmarshaller(ctrl)
		app := app{
			AuthService:  authService,
			Unmarshaller: unmarshaller,
			Validator:    NewMockValidator(ctrl),
			Transformer:  NewMockTransformer(ctrl),
			ExtApiCall:   NewMockExtApiCall(ctrl),
		}

		expectedHttpRequest := &http.Request{}

		authService.EXPECT().
			authorize(expectedHttpRequest).
			Return(nil).
			Times(1)

		unmarshaller.EXPECT().
			unmarshall(gomock.Any()).
			Return(errors.New("some error")).
			Times(1)

		rr := httptest.NewRecorder()
		app.handle(rr, &http.Request{})

		assert.Equal(t, 500, rr.Code)
		assert.Equal(t, "some error message", rr.Body.String())
	})

	t.Run("should return an error if validation fails", func(t *testing.T) {
		// we need to make the authorization step pass -> configure the app, put valid user into the request
		// we need to make the parsing step pass
		// provide an invalid request object
	})

	t.Run("should return an error if transformation", func(t *testing.T) {
		// we need to make the authorization step pass -> configure the app, put valid user into the request
		// we need to make the parsing step pass
		// we need to make the validation step pass
		// provide invalid configuration
	})

	// ... and so on
}
