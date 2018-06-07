package accesstoken

import (
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
)

type Credentials struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}



func GetAccessToken(url string, c Credentials) (string, error){
	client := &http.Client{}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(rqBody)) //Build a request
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8") //Setting request Header

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	//Checking status from response todo: check all response formats from any gateway
	if response.StatusCode != http.StatusOK {
		return "",fmt.Errorf("ocorreu um erro! StatusCode: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body) // Reading body
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
}
