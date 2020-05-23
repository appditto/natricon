package db

import "time"

type Donor struct {
	PubKey    string    `json:"pubkey"`
	ExpiresAt time.Time `json:"expires_at"`
}
