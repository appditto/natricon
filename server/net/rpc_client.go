package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/appditto/natricon/server/model"
	"github.com/golang/glog"
)

type RPCClient struct {
	Url string
}

func (client RPCClient) MakeAccountHistoryRequest(account string, count uint) (*model.AccountHistoryResponse, error) {
	request := model.AccountHistoryRequest{
		BaseRequest: model.AccountHistoryAction,
		Account:     account,
		Count:       count,
	}
	requestBody, _ := json.Marshal(request)
	resp, err := http.Post(client.Url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		glog.Fatalf("Error making RPC request", err)
		return nil, errors.New("Error")
	}
	defer resp.Body.Close()
	// Try to deserialize
	var historyResponse model.AccountHistoryResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Fatalf("Error decoding request body", err)
		return nil, errors.New("Error")
	}
	err = json.Unmarshal(body, &historyResponse)
	if err != nil {
		glog.Fatalf("Error unmarshaling response", err)
		return nil, errors.New("Error")
	}
	return &historyResponse, nil
}
