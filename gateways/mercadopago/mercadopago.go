package mercadopago

import (
	"encoding/json"

	"github.com/golangba/PayGo-Gateway/gateways/helpers/sendrequest"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago/config"
)

type MercadoPagoBase struct{}
type Action uint8

func (m MercadoPagoBase) GetContentType() (string, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return "", err
	}
	return conf.Charset, nil
}

func (m MercadoPagoBase) GetBody(pgr sendrequest.PayGoRequest) ([]byte, error) {
	j, err := json.Marshal(pgr)
	if err != nil {
		return nil, err
	}
	return j, nil
}
