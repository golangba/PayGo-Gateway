package accesstoken

import (
	"encoding/json"
	"fmt"
	"PayGo-Gateway2/gateways/helpers/sendrequest"
)

type resultCredentials struct {

}
type Credentials struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	url string
	err error
	response map[string]interface{}
}
func (c Credentials) GetUrl() (string, error) {
	if len(c.url) == 0{
		return "", fmt.Errorf("the url not yet implemented")
	}

	return c.url, nil
}

func (c Credentials) GetContentType() (string, error) {
	return "application/json;charset=utf-8", nil
}

func (c *Credentials) GetBody() ([]byte, error) {

	j, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (c *Credentials) SetResponse(b []byte) error {
	err := json.Unmarshal(b, &c.response)
	if err != nil {
		return err
	}
	return nil
}

func GetAccessToken(url string,c Credentials) (string, error){
	c.url = url
	_, err:=sendrequest.SendRequest(&c, "POST")
	if err != nil {
		return "", fmt.Errorf("Error: %s", err)
	}

	if m, ok := c.response["message"]; ok{
		return "", fmt.Errorf("mercado pago message: %v", m)
	}

	if token, ok:=c.response["access_token"]; ok{
		return fmt.Sprintf("%v",token), nil
	}

	return "", fmt.Errorf("some problems was encountered, please contact the developers")
}