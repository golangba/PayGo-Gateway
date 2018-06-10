package accesstoken

import (
	"fmt"
	"testing"

	"github.com/golangba/PayGo-Gateway/gateways/mercadopago/config"
)

func TestGetAccessToken(t *testing.T) {
	r, err := GetAccessToken()
	checkTestError(err, t)
	fmt.Println(r)
}

func TestUpdateToken(t *testing.T) {
	conf, err := config.GetConfig()
	checkTestError(err, t)

	err = UpdateToken(conf)
	checkTestError(err, t)
	if len(conf.ApiToken) == 0 {
		t.Errorf("Token not set")
	}
}
func checkTestError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
