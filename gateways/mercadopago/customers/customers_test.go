package customers

import (
	"fmt"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	c := Customer{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "fulanodetal@test.com",
		Phone:     Phone{AreaCode: "081", Number: "999999999"},
	}
	fmt.Println(c)
}
