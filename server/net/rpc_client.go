package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/appditto/natricon/server/model"
	"github.com/appditto/natricon/server/utils"
	"github.com/golang/glog"
)

type RPCClient struct {
	Url        string
	httpClient *http.Client
}

func NewRPCClient(url string) *RPCClient {
	return &RPCClient{
		Url: url,
		httpClient: &http.Client{
			Timeout: time.Second * 30, // Set a timeout for all requests
		},
	}
}

// Base request
func (client *RPCClient) makeRequest(request interface{}) ([]byte, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := client.httpClient.Post(client.Url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("non-200 response from server")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
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
