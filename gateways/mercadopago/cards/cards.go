package cards

import (
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago/config"
	"encoding/json"
	"fmt"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago/accesstoken"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago"
	"strings"
	"github.com/golangba/PayGo-Gateway/gateways/helpers/sendrequest"
)

type SecurityCode struct {
	Length       uint   `json:"length,omitempty" json:"length,omitempty"`
	CardLocation string `json:"card_location,omitempty" json:"card_location,omitempty"`
}

type Issuer struct {
	ID   uint   `json:"id,omitempty" json:"id,omitempty"`
	Name string `json:"name,omitempty" json:"name,omitempty"`
}

type Cardholder struct {
	Name           uint   `json:"name,omitempty" json:"name,omitempty"`
	Identification uint   `json:"identification,omitempty" json:"identification,omitempty"`
	Number         uint   `json:"number,omitempty" json:"number,omitempty"`
	Subtype        string `json:"subtype,omitempty" json:"subtype,omitempty"`
	Type           string `json:"type,omitempty" json:"type,omitempty"`
}

type Card struct {
	mercadopago.MercadoPagoBase `json:"-" url:"-"`
	action mercadopago.Action
	ID              string       `json:"id,omitempty" url:"id,omitempty"`
	CustomerId      string       `json:"customer_id,omitempty" url:"customer_id,omitempty"`
	ExpirationMonth uint         `json:"expiration_month,omitempty" url:"expiration_month,omitempty"`
	ExpirationYear  uint         `json:"expiration_year,omitempty" json:"expiration_year,omitempty"`
	FirstSixDigits  string       `json:"first_six_digits,omitempty" json:"first_six_digits,omitempty"`
	LastFourDigits  string       `json:"last_four_digits,omitempty" json:"last_four_digits,omitempty"`
	PaymentMethods  interface{}  `json:"payment_methods,omitempty" json:"payment_methods,omitempty"`
	SecurityCode    SecurityCode `json:"security_code,omitempty" json:"security_code,omitempty"`
	Issuer          Issuer       `json:"issuer,omitempty" json:"issuer,omitempty"`
	Cardholder      Cardholder   `json:"cardholder,omitempty" json:"cardholder,omitempty"`
	DateCreated     string       `json:"date_created,omitempty" json:"date_created,omitempty"`
	DateLastUpdate  string       `json:"date_last_update,omitempty" url:"date_last_update,omitempty"`
}

const (
	CREATE mercadopago.Action = iota
	UPDATE
	GET
	DELETE
	LIST
) 
var response map[string]interface{}

func (c Card) GetUrl() (string, error) {
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

func (c Card) GetMethod() string {
	switch c.action {
	case CREATE:
		return "POST"
	case UPDATE:
		return "PUT"
	case DELETE:
		return "DELETE"
	default:
		return "GET"
	}
}

func (c *Card) SetResponse(b []byte) error {
	err := json.Unmarshal(b, &response)
	if err != nil {
		return err
	} else if m, ok := response["message"]; ok { //verifying if have messages from the mercadopago
		return fmt.Errorf("mercado pago message: %v;\n Cause:%+v", m, response["cause"])
	}

	err = json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	return nil
}

func SaveCard(action mercadopago.Action, c *Card)(bool, error){
	switch action {
	case CREATE:
		c.action = CREATE
	case UPDATE:
		c.action = UPDATE
		if len(c.ID) == 0 {
			return false, fmt.Errorf("invalid Card.ID")
		}
	}

	_, err := sendrequest.SendRequest(c)
	if err != nil {
		return false, err
	}
	return true, nil
}

func mountUrl(c Card, conf *config.Config) (string, error) {
	var url string
	if len(c.CustomerId) ==0{
		return "", fmt.Errorf("you must to set Card.CustomerId")
	}
	if len(conf.ApiToken) == 0 {
		err := accesstoken.UpdateToken(conf)
		if err != nil {
			return "", err
		}
	}
	urlParams := conf.ApiToken
	route, err := config.GetRoute("customers.cards")
	if err != nil {
		return "", err
	}
	route = strings.Replace(route, "{$cid}", c.CustomerId, 1)

	switch c.action {
	case GET, UPDATE, DELETE:
		url = fmt.Sprintf("%s/%s%s/%s?access_token=%s", conf.ApiUrl, conf.ApiVersion, route, c.ID, urlParams)
	case CREATE, LIST:
		url = fmt.Sprintf("%s/%s%s?access_token=%s", conf.ApiUrl, conf.ApiVersion, route, urlParams)
	default:
		return "", fmt.Errorf("user action not allowed")
	}

	return url, nil
}
