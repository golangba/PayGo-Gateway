package customers

import (
	"fmt"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	//creation test
	 c := Customer{
		FirstName:      "Vitor",
		LastName:       "Vicente Assunção",
		Email:          "vitorvicenteassuncao_1@signainfo.com.br",
		Phone:          Phone{AreaCode: "11", Number: "988524769"},
		Identification: Identification{Type: "CPF", Number: "05080749245"},
		Address: Address{
			ZipCode:      "04613020",
			StreetName:   "Rua Maria Peres Auge",
			StreetNumber: 624,
		},
		Description: "lorem ipsum",
	}
	r, err := SaveCustomer(CREATE, &c)
	checkTestError(err, t)
	fmt.Printf("Request status: %t\n", r)
	fmt.Printf("Costumer: %+v\n", c)

	//update test
	 c = Customer{
		ID:             "327319823-TjF62w7HJBkxIy",
		LastName:       "Vicente Assunção Souza",
		Description: "updated",
	}
	r, err = SaveCustomer(UPDATE, &c)
	checkTestError(err, t)
	fmt.Printf("Request status: %t\n", r)
	fmt.Printf("Costumer: %+v\n", c)
}

func TestSearchCustomer(t *testing.T) {
	sc := SearchCustomersParams{
		Offset: 0,
		Limit:  5,
	}
	sc.Identification = Identification{Number: "05080749245"}
	r, err := SearchCustomers(sc)
	checkTestError(err, t)
	fmt.Println("Users found")
	for _, result := range r.Results {
		fmt.Println(result.ID)
	}

}

func TestGetCustomer(t *testing.T) {
	cid := "327319823-TjF62w7HJBkxIy"
	c, err := GetCustomer(cid)
	checkTestError(err, t)
	fmt.Printf("User found: %+v", c)
}

func checkTestError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
