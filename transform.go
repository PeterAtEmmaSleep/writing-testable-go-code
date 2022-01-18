package main

//go:generate mockgen -source=$GOFILE -destination=transform_mock.go -package=$GOPACKAGE
type Transformer interface {
	transform(appRequest *AppRequest) (*ExtRequest, error)
}

type transformer struct {
	someCfgValue string
}

func NewTransformer(someCfgValue string) Transformer {
	return &transformer{
		someCfgValue: someCfgValue,
	}
}

func (t *transformer) transform(appRequest *AppRequest) (*ExtRequest, error) {
	return &ExtRequest{
		someField: appRequest.someField + t.someCfgValue,
	}, nil
}
