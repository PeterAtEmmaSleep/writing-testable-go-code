package main

import "net/http"

func main() {
	authService := NewAuthService("expected username")
	unmarshaller := NewUnmarshaller()
	validator := NewValidator()
	transformer := NewTransformer("some config value")
	extApiCall := NewExtApiCall()

	app := app{
		AuthService:  authService,
		Unmarshaller: unmarshaller,
		Validator:    validator,
		Transformer:  transformer,
		ExtApiCall:   extApiCall,
	}

	http.HandleFunc("/path", app.handle)
	_ = http.ListenAndServe(":8090", nil)
}
