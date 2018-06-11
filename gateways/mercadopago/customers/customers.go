package customers

import (
	"encoding/json"
	"fmt"

	"github.com/golangba/PayGo-Gateway/gateways/helpers/sendrequest"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago/accesstoken"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago/config"
)

type Phone struct {
	AreaCode string `json:"area_code,omitempty" url:"area_code,omitempty"`
	Number   string `json:"number,omitempty" url:"number,omitempty"`
}

type Identification struct {
	Type   string `json:"type,omitempty" url:"type,omitempty"`
	Number string `json:"number,omitempty" url:"number,omitempty"`
}

type Address struct {
	ID           string `json:"id,omitempty" url:"id,omitempty"`
	ZipCode      string `json:"zip_code,omitempty" url:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty" url:"street_name,omitempty"`
	StreetNumber uint   `json:"street_number,omitempty" url:"street_number,omitempty"`
}

type Customer struct {
	mercadopago.MercadoPagoBase `json:"-" url:"-"`
	action                      action
	ID                          string         `json:"id,omitempty" url:"id,omitempty"`
	Email                       string         `json:"email,omitempty" url:"email,omitempty"`
	FirstName                   string         `json:"first_name,omitempty" url:"first_name,omitempty"`
	LastName                    string         `json:"last_name,omitempty" url:"last_name,omitempty"`
	Phone                       Phone          `json:"phone,omitempty" url:"phone,omitempty"`
	Identification              Identification `json:"identification,omitempty" url:"identification,omitempty"`
	Address                     Address        `json:"address,omitempty" url:"address,omitempty"`
	DateRegistered              string         `json:"date_registered,omitempty" url:"date_registered,omitempty"`
	Description                 string         `json:"description,omitempty" url:"description,omitempty"`
	DateCreated                 string         `json:"date_created,omitempty" url:"date_created,omitempty"`
	DateLastUpdated             string         `json:"date_last_updated,omitempty" url:"date_last_updated,omitempty"`
	Metadata                    interface{}    `json:"metadata,omitempty" url:"metadata,omitempty"`
	DefaultCard                 string         `json:"default_card,omitempty" url:"default_card,omitempty"`
	DefaultAddress              string         `json:"default_address,omitempty" url:"default_address,omitempty"`
	Cards                       interface{}    `json:"cards,omitempty" url:"cards,omitempty"`
	Addresses                   interface{}    `json:"addresses,omitempty" url:"addresses,omitempty"`
	LiveMode                    bool           `json:"live_mode,omitempty" url:"live_mode,omitempty"`
}

type action uint8

const (
	CREATE action = iota
	GET
	UPDATE
	SEARCH
)

var response map[string]interface{}

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

func (c Customer) GetMethod() string {
	switch c.action {
	case CREATE:
		return "POST"
	case UPDATE:
		return "PUT"
	default:
		return "GET"
	}
}

func (c *Customer) SetResponse(b []byte) error {
	err := json.Unmarshal(b, &response)
	if err != nil {
		return err
	} else if m, ok:= response["message"]; ok{ //verifying if have messages from the mercadopago
		return fmt.Errorf("mercado pago message: %v;\n Cause:%+v", m, response["cause"])
	}

	err = json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	return nil
}

//create and update customer
func SaveCustomer(action action,c *Customer) (bool, error) {
	switch action {
	case CREATE:
		c.action = CREATE
	case UPDATE:
		c.action = UPDATE
		if len(c.ID) == 0 {
			return false, fmt.Errorf("invalid Customer.ID")
		}
	}

	_, err := sendrequest.SendRequest(c)
	if err != nil {
		return false, err
	}
	return true, nil
}

//return a customer by id
func GetCustomer(cid string) (*Customer, error) {
	c := &Customer{action: GET, ID: cid}
	_, err := sendrequest.SendRequest(c)
	if err != nil {
		return nil, err
	}
	return c, nil
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

	urlParams := conf.ApiToken

	switch c.action {
	case GET, UPDATE:
		route, err := config.GetRoute(fmt.Sprintf("customers.%s", routeName))
		if err != nil {
			return "", err
		}
		url = fmt.Sprintf("%s/%s%s/%s?access_token=%s", conf.ApiUrl, conf.ApiVersion, route, c.ID, urlParams)
	case CREATE:
		route, err := config.GetRoute(fmt.Sprintf("customers.%s", routeName))
		if err != nil {
			return "", err
		}
		url = fmt.Sprintf("%s/%s%s?access_token=%s", conf.ApiUrl, conf.ApiVersion, route, urlParams)
	default:
		return "", fmt.Errorf("user action not allowed")
	}

	return url, nil
}
