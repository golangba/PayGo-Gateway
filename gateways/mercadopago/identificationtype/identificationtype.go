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
	ID        string `json:"id, omitempty"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	MinLength uint   `json:"min_length"`
	MaxLength uint   `json:"max_length"`
}

var responseId []IdentificationType

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
		err = accesstoken.UpdateToken(conf)
		if err != nil {
			return "", err
		}
	}
	url := fmt.Sprintf("%s%s?access_token=%s", conf.ApiUrl, route, conf.ApiToken)
	return url, nil
}
func (i IdentificationType) GetMethod() string {
	return "GET"
}

func (i IdentificationType) SetResponse(b []byte) error {
	err := json.Unmarshal(b, &responseId)
	if err != nil {
		return err
	}
	return nil
}

//returns all personal identifications allowed to the client
func GetIdentificationTypes() ([]IdentificationType, error) {
	i := IdentificationType{}
	_, err := sendrequest.SendRequest(&i)
	if err != nil {
		return nil, err
	}
	return responseId, nil
}
