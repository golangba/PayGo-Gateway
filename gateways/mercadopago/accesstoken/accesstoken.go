package accesstoken

import (
	"encoding/json"
	"fmt"
	"paygo/gateways/helpers/sendrequest"
	"paygo/gateways/mercadopago"
	"github.com/spf13/viper"
)

type credentials struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	response     map[string]interface{}
}

func (c credentials) GetUrl() (string, error) {
	conf, err := mercadopago.GetConfig()
	if err != nil {
		return "", err
	}
	route, err := mercadopago.GetRoute("accesstoken")
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s%s", conf.ApiUrl, route)
	return url, nil
}

func (c credentials) GetContentType() (string, error) {
	conf, err := mercadopago.GetConfig()
	if err != nil {
		return "", err
	}
	return conf.Charset, nil
}

func (c *credentials) GetBody() ([]byte, error) {
	j, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (c *credentials) SetResponse(b []byte) error {
	err := json.Unmarshal(b, &c.response)
	if err != nil {
		return err
	}
	return nil
}

func GetAccessToken() (string, error) {
	conf, err := mercadopago.GetConfig()
	if err != nil {
		return "", fmt.Errorf("Error: %s", err)
	}
	c := new(credentials)
	c.GrantType = "client_credentials"
	c.ClientId = conf.ApiClientID
	c.ClientSecret = conf.ApiClientSecret

	_, err = sendrequest.SendRequest(c, "POST")
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

func UpdateToken(c *mercadopago.Config) error{
	token, err := GetAccessToken()
	if err != nil { // Handle errors reading the config file
		return err
	}
	viper.Set("mercadopago.config.apiToken", token)
	c.ApiToken = token
	return nil
}
