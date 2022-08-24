package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Post - if Content-type header is missing, application/json is added as default
func Post[T any, I any](url string, request *T, headers map[string]string) (*I, error) {
	return execute[T, I](url, request, headers, http.MethodPost)
}

// Get - if Content-type header is missing, application/json is added as default
func Get[T any, I any](url string, request *T, headers map[string]string) (*I, error) {
	return execute[T, I](url, request, headers, http.MethodGet)
}

func execute[T any, I any](url string, requestData *T, headers map[string]string, method string) (*I, error) {

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(requestData)
	if err != nil {
		fmt.Printf("Error in create payload: %s\n", err)
		return nil, err
	}

	request, err := http.NewRequest(method, url, payloadBuf)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}

	if headers["Content-Type"] == "" {
		request.Header.Add("Content-Type", "application/json")
	}

	// map headers from caller
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("client: error read response.Body: %s\n", err)
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("client: error in response with code: %d, with error %s ", response.StatusCode, string(responseData))
	}

	var bodyConverted I
	err = json.Unmarshal(responseData, &bodyConverted)

	return &bodyConverted, err
}
