package model

// Serializable/Deserializable models related to nano callback
type Callback struct {
	Hash   string `json:"hash"`
	Block  string `json:"block"`
	IsSend string `json:"is_send"`
	Amount string `json:"amount"`
}

type Block struct {
	LinkAsAccount string `json:"link_as_account"`
}

// RPC requests
type BaseRequest struct {
	Action string `json:"action"`
}

var AccountHistoryAction BaseRequest = BaseRequest{Action: "account_history"}

type AccountHistoryRequest struct {
	BaseRequest
	Account string `json:"account"`
	Count   uint   `json:"count"`
}

// RPC responses
type HistoryItem struct {
	Type           string `json:"type"`
	Account        string `json:"account"`
	Amount         string `json:"amount"`
	LocalTimestamp string `json:"local_timestamp"`
	Hash           string `json:"hash"`
}

type AccountHistoryResponse struct {
	Account string        `json:"account"`
	History []HistoryItem `json:"history"`
}
