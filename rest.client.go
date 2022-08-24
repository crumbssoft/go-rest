package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post[T any, I any](url string, request *T) (*I, error) {

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(request)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", payloadBuf)

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	fmt.Print(response)

	var accConverted I
	err = json.Unmarshal(responseData, &accConverted)
	return &accConverted, err
}
