package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/appditto/natricon/server/model"
	"github.com/appditto/natricon/server/utils"
	"github.com/golang/glog"
)

type RPCClient struct {
	Url string
}

// Base request
func (client RPCClient) makeRequest(request interface{}) ([]byte, error) {
	requestBody, _ := json.Marshal(request)
	// HTTP post
	resp, err := http.Post(client.Url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		glog.Errorf("Error making RPC request %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	// Try to decode+deserialize
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorf("Error decoding response body %s", err)
		return nil, err
	}
	return body, nil
}

// Nano account_history request
func (client RPCClient) MakeAccountHistoryRequest(account string, count uint) (*model.AccountHistoryResponse, error) {
	// Build request
	request := model.AccountHistoryRequest{
		BaseRequest: model.AccountHistoryAction,
		Account:     account,
		Count:       count,
	}
	response, err := client.makeRequest(request)
	if err != nil {
		glog.Errorf("Error making request %s", err)
		return nil, err
	}
	var historyResponse model.AccountHistoryResponse
	err = json.Unmarshal(response, &historyResponse)
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
	response, err := client.makeRequest(request)
	if err != nil {
		glog.Errorf("Error making request %s", err)
		return nil, err
	}
	// Try to decode+deserialize
	var quorumResponse model.ConfirmationQuorumResponse
	err = json.Unmarshal(response, &quorumResponse)
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
	response, err := client.makeRequest(request)
	if err != nil {
		glog.Errorf("Error making request %s", err)
		return nil, err
	}
	// Try to decode+deserialize
	var repResponse model.RepresentativeResponse
	err = json.Unmarshal(response, &repResponse)
	if err != nil {
		glog.Errorf("Error unmarshaling response %s", err)
		return nil, errors.New("Error")
	}
	return &repResponse, nil
}

// send
func (client RPCClient) MakeSendRequest(source string, destination string, amountRaw string, id string, wallet string) (*model.SendResponse, error) {
	// Build request
	request := model.SendRequest{
		BaseRequest: model.SendAction,
		Source:      source,
		Destination: destination,
		AmountRaw:   amountRaw,
		ID:          id,
		Wallet:      wallet,
	}
	bpowKey := utils.GetEnv("BPOW_KEY", "")
	if bpowKey != "" {
		request.BpowKey = &bpowKey
	}
	response, err := client.makeRequest(request)
	if err != nil {
		glog.Errorf("Error making request %s", err)
		return nil, err
	}
	// Try to decode+deserialize
	var sendResponse model.SendResponse
	err = json.Unmarshal(response, &sendResponse)
	if err != nil {
		glog.Errorf("Error unmarshaling response %s", err)
		return nil, errors.New("Error")
	}
	return &sendResponse, nil
}
