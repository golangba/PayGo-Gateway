package sendRequest

import (
	"testing"
	"encoding/json"
)

type TestResquest struct {
	Title string `json:"title"`
	Body string `json:"body"`
	UserId int64 `json:"userId"`
	Response interface{}
}

func (t TestResquest) GetUrl() (string, error){
	return "https://jsonplaceholder.typicode.com/posts/", nil
}

func (t TestResquest) MarshalJSON() ([]byte, error){
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (t *TestResquest) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, t.Response)
	if err != nil {
		return err
	}
	return nil
}

func TestSendRequest(t *testing.T) {
	tr := new (TestResquest)
	tr.Title = "foo"
	tr.Body = "bar"
	tr.UserId = 1

	err := SendRequest(tr, "POST")
	if err != nil {
		t.Errorf("error on send Request: %s\n", err)
	}
}