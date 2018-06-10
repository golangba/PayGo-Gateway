package identificationtype

import (
	"encoding/json"
	"fmt"

	"github.com/golangba/PayGo-Gateway/gateways/helpers/sendrequest"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago/accesstoken"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago/config"
)

type IdentificationType struct {
	mercadopago.MercadoPagoBase
	ID        string `json:"id, omnitempty"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	MinLength int    `json:"min_length"`
	MaxLength int    `json:"max_length"`
}

var response []IdentificationType

func (i IdentificationType) GetUrl() (string, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return "", err
	}
	route, err := config.GetRoute("identificationType")
	if err != nil {
		return "", err
	}
	if len(conf.ApiToken) == 0 {
		accesstoken.UpdateToken(conf)
	}
	url := fmt.Sprintf("%s%s?access_token=%s", conf.ApiUrl, route, conf.ApiToken)
	return url, nil
}

func (i IdentificationType) SetResponse(b []byte) error {
	err := json.Unmarshal(b, &response)
	if err != nil {
		return err
	}
	return nil
}

func GetIdentificationTypes() ([]IdentificationType, error) {
	i := IdentificationType{}
	_, err := sendrequest.SendRequest(&i, "GET")
	if err != nil {
		return nil, err
	}
	return response, nil
}
