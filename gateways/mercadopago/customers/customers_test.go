package customers

import (
	"fmt"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	c := Customer{
		FirstName:      "Vitor",
		LastName:       "Vicente Assunção",
		Email:          "vitorvicenteassuncao_1@signainfo.com.br",
		Phone:          Phone{AreaCode: "11", Number: "988524769"},
		Identification: Identification{Type: "CPF", Number: "05080749245"},
		Address: Address{
			ZipCode: "04613020",
			StreetName: "Rua Maria Peres Auge",
			StreetNumber: 624,
		},
		Description: "lorem ipsum",
	}
	r, err := CreateCustomer(&c)
	checkTestError(err, t)
	fmt.Println(r)
	fmt.Println(c)
}

func checkTestError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

