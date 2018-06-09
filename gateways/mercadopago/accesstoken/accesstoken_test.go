package accesstoken

import (
	"fmt"
	"testing"
	"paygo/gateways/mercadopago"
)

func TestGetAccessToken(t *testing.T) {
	r, err := GetAccessToken()
	checkTestError(err, t)
	fmt.Println(r)
}

func TestUpdateToken(t *testing.T) {
	config, err := mercadopago.GetConfig()
	checkTestError(err, t)

	err = UpdateToken(config)
	checkTestError(err, t)
	if len(config.ApiToken) == 0 {
		t.Errorf("Token not set")
	}
}
func checkTestError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
