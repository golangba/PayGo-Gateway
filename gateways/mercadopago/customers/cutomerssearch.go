package customers

import (
	"encoding/json"
	"fmt"
	"github.com/golangba/PayGo-Gateway/gateways/helpers/sendrequest"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago/config"
	"github.com/google/go-querystring/query"
)

type SearchCustomersParams struct {
	Customer
	Offset  uint       `json:"offset,omitempty" url:"offset,omitempty"`
	Limit   uint       `json:"limit,omitempty" url:"limit,omitempty"`
	Results []Customer `json:"results,omitempty" url:"-"`
}

func (sc SearchCustomersParams) GetUrl() (string, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return "", err
	}

	urlParams := conf.ApiToken
	v, err := query.Values(sc)
	if err != nil {
		return "", err
	}
	urlParams += fmt.Sprintf("&%s", v.Encode())

	route, err := config.GetRoute(fmt.Sprintf("customers.%s", "search"))
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s%s?access_token=%s", conf.ApiUrl, conf.ApiVersion, route, urlParams)
	return url, nil
}
func (sc *SearchCustomersParams) SetResponse(b []byte) error {
	err := json.Unmarshal(b, &sc)
	if err != nil {
		return err
	}
	return nil
}

func SearchCustomers(sc SearchCustomersParams) (*SearchCustomersParams, error) {
	sc.action = search
	_, err := sendrequest.SendRequest(&sc)
	if err != nil {
		return nil, err
	}
	return &sc, nil
}
