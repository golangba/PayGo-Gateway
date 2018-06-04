package sendRequest

import (
	"encoding/json"
	"fmt"
)

type ToRequest interface {
	GetUrl() (string, error)
	json.Marshaler
	json.Unmarshaler
}

func SendRequest(rq ToRequest, method string) error{
	return fmt.Errorf("not implemented yet")
}