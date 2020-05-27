package model

// Serializable/Deserializable models related to nano callback
type Callback struct {
	Hash   string `json:"hash"`
	Block  string `json:"block"`
	IsSend string `json:"is_send"`
	Amount string `json:"amount"`
}

type Block struct {
	Account       string `json:"account"`
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

// confirmation_quorum
var ConfirmationQuorumAction BaseRequest = BaseRequest{Action: "confirmation_quorum"}

type ConfirmationQuorumRequest struct {
	BaseRequest
}

// representatives
var RepresentativesAction BaseRequest = BaseRequest{Action: "representatives"}

type RepresentativesRequest struct {
	BaseRequest
}

// send
// representatives
var SendAction BaseRequest = BaseRequest{Action: "send"}

type SendRequest struct {
	BaseRequest
	Wallet      string `json:"wallet"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	AmountRaw   string `json:"amount"`
	ID          string `json:"id"`
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

// Confirmation_quorum

type ConfirmationQuorumResponse struct {
	OnlineWeightTotal string `json:"online_stake_total"`
}

// Representatives
type RepresentativeResponse struct {
	Representatives map[string]string `json:"representatives"`
}

// Send
type SendResponse struct {
	Block string `json:"block"`
}
