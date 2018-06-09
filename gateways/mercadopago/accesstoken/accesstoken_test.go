package accesstoken

import (
	"testing"
	"fmt"
)

func TestGetAccessToken(t *testing.T) {
	c := Credentials{
		GrantType: "client_credentials",
		ClientId:"7130376828807209",
		ClientSecret:"ylLotexirqboqPmqgtjOMXWfYOFh3VY4",
	}
	url := "https://api.mercadopago.com/oauth/token"

	r, err:= GetAccessToken(url, c)
	checkTestError(err, t)
	fmt.Println(r)
}

func checkTestError(err error, t *testing.T){
	if err != nil{
		t.Errorf("Error: %s", err)
	}
}

