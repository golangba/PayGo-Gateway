package accesstoken

import (
	"encoding/json"
	"fmt"

	"github.com/golangba/PayGo-Gateway/gateways/helpers/sendrequest"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago"
	"github.com/golangba/PayGo-Gateway/gateways/mercadopago/config"
	"github.com/spf13/viper"
)

type credentials struct {
	mercadopago.MercadoPagoBase `json:"omitempty"`
	GrantType                   string `json:"grant_type"`
	ClientId                    string `json:"client_id"`
	ClientSecret                string `json:"client_secret"`
	response                    map[string]interface{}
}

func (c credentials) GetUrl() (string, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return "", err
	}
	route, err := config.GetRoute("accesstoken")
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s%s", conf.ApiUrl, route)
	return url, nil
}
func (c credentials) GetMethod() string {
	return "POST"
}

func (c *credentials) SetResponse(b []byte) error {
	err := json.Unmarshal(b, &c.response)
	if err != nil {
		return err
	}
	return nil
}

//returns access token that will be used in application
func GetAccessToken() (string, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return "", fmt.Errorf("Error: %s", err)
	}
	c := credentials{
		GrantType:    "client_credentials",
		ClientId:     conf.ApiClientID,
		ClientSecret: conf.ApiClientSecret,
	}

	_, err = sendrequest.SendRequest(&c)
	if err != nil {
		return "", err
	}

	if m, ok := c.response["message"]; ok {
		return "", fmt.Errorf("mercado pago message: %v", m)
	}

	if token, ok := c.response["access_token"]; ok {
		return fmt.Sprintf("%v", token), nil
	}

	return "", fmt.Errorf("some problems was encountered, please contact the developers")
}

func UpdateToken(c *config.Config) error {
	token, err := GetAccessToken()
	if err != nil { // Handle errors reading the config file
		return err
	}
	viper.Set("config.config.apiToken", token)
	c.ApiToken = token
	return nil
}
