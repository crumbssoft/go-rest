package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gqls/account/graph/model"
	"gqls/account/service"
	"io/ioutil"
	"net/http"
)

var baseUrl = "http://130.61.143.84:8080/msaccount/"

func Accounts(userId string) ([]*model.Account, error) {
	url := baseUrl + "v1/accounts"
	header := &service.Header{
		CallerId:        "goCaller",
		CallerProcessId: "goCallerFirstProcessId",
		CorrelationId:   "graphQlCorrelation",
	}

	accountBody := &service.AccountBody{UserId: userId}

	body := &service.AccountRequest{
		Header: header,
		Body:   accountBody,
	}

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(body)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", payloadBuf)
	fmt.Print(response)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var accConverted service.AccountResponse
	err = json.Unmarshal(responseData, &accConverted)
	return accConverted.Body, err
}

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
