package db

import "time"

type Donator struct {
	Address   string    `json:"address"`
	ExpiresAt time.Time `json:"expires_at"`
}
