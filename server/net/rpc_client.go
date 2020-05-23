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

// Nano account_history request
func (client RPCClient) MakeAccountHistoryRequest(account string, count uint) (*model.AccountHistoryResponse, error) {
	// Build request
	request := model.AccountHistoryRequest{
		BaseRequest: model.AccountHistoryAction,
		Account:     account,
		Count:       count,
	}
	requestBody, _ := json.Marshal(request)
	// HTTP post
	resp, err := http.Post(client.Url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		glog.Errorf("Error making RPC request %s", err)
		return nil, errors.New("Error")
	}
	defer resp.Body.Close()
	// Try to decode+deserialize
	var historyResponse model.AccountHistoryResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorf("Error decoding request body %s", err)
		return nil, errors.New("Error")
	}
	err = json.Unmarshal(body, &historyResponse)
	if err != nil {
		glog.Errorf("Error unmarshaling response %s", err)
		return nil, errors.New("Error")
	}
	return &historyResponse, nil
}

// Nano confirmation_quorum request
func (client RPCClient) MakeConfirmationQuorumRequest() (*model.ConfirmationQuorumResponse, error) {
	// Build request
	request := model.ConfirmationQuorumRequest{
		BaseRequest: model.ConfirmationQuorumAction,
	}
	requestBody, _ := json.Marshal(request)
	// HTTP post
	resp, err := http.Post(client.Url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		glog.Errorf("Error making RPC request %s", err)
		return nil, errors.New("Error")
	}
	defer resp.Body.Close()
	// Try to decode+deserialize
	var quorumResponse model.ConfirmationQuorumResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorf("Error decoding request body %s", err)
		return nil, errors.New("Error")
	}
	err = json.Unmarshal(body, &quorumResponse)
	if err != nil {
		glog.Errorf("Error unmarshaling response %s", err)
		return nil, errors.New("Error")
	}
	return &quorumResponse, nil
}

// representatives
func (client RPCClient) MakeRepresentativesRequest() (*model.RepresentativeResponse, error) {
	// Build request
	request := model.RepresentativesRequest{
		BaseRequest: model.RepresentativesAction,
	}
	requestBody, _ := json.Marshal(request)
	// HTTP post
	resp, err := http.Post(client.Url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		glog.Errorf("Error making RPC request %s", err)
		return nil, errors.New("Error")
	}
	defer resp.Body.Close()
	// Try to decode+deserialize
	var repResponse model.RepresentativeResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorf("Error decoding request body %s", err)
		return nil, errors.New("Error")
	}
	err = json.Unmarshal(body, &repResponse)
	if err != nil {
		glog.Errorf("Error unmarshaling response %s", err)
		return nil, errors.New("Error")
	}
	return &repResponse, nil
}
