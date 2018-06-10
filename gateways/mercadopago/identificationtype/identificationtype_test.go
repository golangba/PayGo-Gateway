package identificationtype

import (
	"fmt"
	"testing"
)

func TestGetIdentificationTypes(t *testing.T) {
	i, err := GetIdentificationTypes()
	checkTestError(err, t)
	fmt.Println("ola")
	fmt.Println(i)
}

func checkTestError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
