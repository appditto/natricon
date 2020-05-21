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
