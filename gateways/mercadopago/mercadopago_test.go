package mercadopago

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	config, err := GetConfig()
	checkTestError(err, t)
	fmt.Println(config)
}

func TestGetRoute(t *testing.T) {
	route, err := GetRoute("accesstoken")
	checkTestError(err, t)
	if route != "/oauth/token" {
		t.Errorf("Wrong route")
	}
	fmt.Println(route)
}

func checkTestError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
