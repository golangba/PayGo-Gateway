package sendrequest

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//Interface used for send requests, prepare the data and set a response like the client wants
type PayGoRequest interface {
	GetUrl() (string, error)                      //returns the url that will be used
	GetBody(request PayGoRequest) ([]byte, error) //returns the body of request [nil case GET or DELETE]
	GetContentType() (string, error)              //returns the content-type from the application
	GetMethod() string                            //method of the request
	SetResponse([]byte) error                     //set the response and treat it
}

var bodyMethodsAllowed map[string]bool

func init() {
	bodyMethodsAllowed = make(map[string]bool)
	bodyMethodsAllowed["POST"] = true
	bodyMethodsAllowed["PUT"] = true
}

func prepareRequest(rq PayGoRequest, method string) (*http.Request, error) {
	url, err := rq.GetUrl() //Get Url from PayGoRequest Interface
	if err != nil {
		return nil, err
	}

	bma, ok := bodyMethodsAllowed[method]
	var rqBody []byte
	if ok || bma {
		rqBody, err = rq.GetBody(rq) //Turns interface into json
		if err != nil {
			return nil, err
		}
	} else {
		rqBody = nil
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(rqBody)) //Build a request
	if err != nil {
		return nil, err
	}

	contentType, err := rq.GetContentType()
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", contentType) //Setting request Header

	return request, nil
}

//Send the request to the gateway and set the response to client
func SendRequest(rq PayGoRequest) (int, error) {
	request, err := prepareRequest(rq, rq.GetMethod())
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
