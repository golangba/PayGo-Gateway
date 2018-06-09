package accesstoken

import (
	"encoding/json"
	"fmt"
)

type Credentials struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	url string
}
func (c Credentials) GetUrl() (string, error) {
	if len(c.url) == 0{
		return "", fmt.Errorf("the url not yet implemented")
	}

	return c.url, nil
}

func (c Credentials) GetBody() ([]byte, error) {
	j, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (c *Credentials) SetResponse(b []byte) error {
	err := json.Unmarshal(b, &t.Response)
	if err != nil {
		return err
	}
	return nil
}


func GetAccessToken(url string, c Credentials) (string, error){

}
