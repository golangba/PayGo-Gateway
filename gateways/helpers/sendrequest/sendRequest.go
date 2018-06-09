package sendrequest

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type PayGoRequest interface {
	GetUrl() (string, error)
	GetBody() ([]byte, error)
	GetReqType() (string, error)
	GetCharset() (string, error)
	SetResponse([]byte) error
}

func prepareRequest(rq PayGoRequest, method string) (*http.Request, error){
	url, err := rq.GetUrl() //Get Url from PayGoRequest Interface
	if err != nil {
		return nil, err
	}

	rqBody, err := rq.GetBody() //Turns interface into json
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(rqBody)) //Build a request
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8") //Setting request Header

	return request, nil
}

//
func SendRequest(rq PayGoRequest, method string) (int, error) {
	request, err := prepareRequest(rq, method)
	if err != nil {
		return 0, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body) // Reading body
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	err = rq.SetResponse(body) // Setting a response how user wants
	if err != nil {
		return 0, err
	}

	return response.StatusCode, nil
}
