package customers

import (
	"encoding/json"
	"fmt"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago/config"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago/accesstoken"
	"github.com/golangba/PayGo-Gateway/gateways/helpers/sendrequest"
)

type Phone struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Address struct {
	ID           string `json:"id, ominitempty"`
	ZipCode      string `json:"zip_code"`
	StreetName   string `json:"street_name"`
	StreetNumber uint `json:"street_number, omnitempty"`
}

type Customer struct {
	mercadopago.MercadoPagoBase
	action          action
	ID              string                      `json:"id, ominitempty"`
	Email           string                      `json:"email"`
	FirstName       string                      `json:"first_name"`
	LastName        string                      `json:"last_name"`
	Phone           Phone                       `json:"phone"`
	Identification  Identification              `json:"identification"`
	Address         Address                     `json:"address"`
	DateRegistered  string                      `json:"date_registered"`
	Description     string                      `json:"description"`
	DateCreated     string                      `json:"date_created, omnitempty"`
	DateLastUpdated string                      `json:"date_last_updated, omnitempty"`
	Metadata        interface{}                      `json:"metadata, omnitempty"`
	DefaultCard     string                      `json:"default_card, omnitempty"`
	DefaultAddress  string                      `json:"default_address, omnitempty"`
	Cards           map[string]interface{} `json:"cards, omnitempty"`
	Addresses       map[string]interface{} `json:"addresses, omnitempty"`
	LiveMode        bool                        `json:"live_mode, omnitempty"`
}

type action uint8

const create action = 1
const get action = 2
const update action = 3
const search action = 4

func (c Customer) GetUrl() (string, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return "", err
	}
	url, err := mountUrl(c, conf)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (c *Customer) SetResponse(b []byte) error {
	err := json.Unmarshal(b, &c)
	if err != nil {
		return err
	}
	return nil
}

func CreateCustomer(c *Customer) (bool, error) {
	c.action = create
	_, err :=sendrequest.SendRequest(c, "POST")
	if err != nil {
		return false, err
	}
	return true, nil
}

func mountUrl(c Customer, conf *config.Config) (string, error) {
	var url string
	routeName := "common"

	if len(conf.ApiToken) == 0 {
		err := accesstoken.UpdateToken(conf)
		if err != nil {
			return "", err
		}
	}
	switch c.action {
	case get, update:
		route, err := config.GetRoute(fmt.Sprintf("customers.%s", routeName))
		if err != nil {
			return "", err
		}
		url = fmt.Sprintf("%s/%s%s/%s?access_token=%s", conf.ApiUrl, conf.ApiVersion, route, c.ID, conf.ApiToken)
	case search:
		routeName = "search"
		fallthrough
	case create:
		route, err := config.GetRoute(fmt.Sprintf("customers.%s", routeName))
		if err != nil {
			return "", err
		}
		url = fmt.Sprintf("%s/%s%s?access_token=%s", conf.ApiUrl, conf.ApiVersion, route, conf.ApiToken)
	default:
		return "", fmt.Errorf("user action not allowed")
	}
	return url, nil
}
