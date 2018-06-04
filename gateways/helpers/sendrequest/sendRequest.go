package sendrequest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"fmt"
)

type PayGoRequest interface {
	GetUrl() (string, error)
	GetJSON() ([]byte, error)
	SetResponse([]byte) error
}

func SendJSONRequest(rq PayGoRequest, method string) error {
	url, err := rq.GetUrl() //Get Url from PayGoRequest Interface
	if err != nil {
		return err
	}

	rqBody, err := rq.GetJSON() //Turns interface into json
	if err != nil {
		return err
	}

	client := &http.Client{}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(rqBody)) //Build a request
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8") //Setting request Header

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	//Checking status from response todo: check all response formats from any gateway
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("something was wrong! StatusCode: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body) // Reading body
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = rq.SetResponse(body) // Setting a response how user wants
	if err != nil {
		return err
	}

	return nil
}
