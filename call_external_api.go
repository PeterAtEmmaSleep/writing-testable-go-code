package main

//go:generate mockgen -source=$GOFILE -destination=call_external_api_mock.go -package=$GOPACKAGE
type ExtApiCall interface {
	CallApi(request *ExtRequest) error
}

type extApiCall struct{}

func NewExtApiCall() ExtApiCall {
	return &extApiCall{}
}

func (s *extApiCall) CallApi(_ *ExtRequest) error {
	return nil
}
