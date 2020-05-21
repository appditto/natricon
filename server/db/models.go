package db

import "time"

type Donor struct {
	Address   string    `json:"address"`
	ExpiresAt time.Time `json:"expires_at"`
}
