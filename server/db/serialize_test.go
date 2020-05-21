package db

import (
	"encoding/json"
	"testing"
	"time"
)

func TestSerializeDonor(t *testing.T) {
	time := time.Date(2020, 11, 19, 19, 19, 19, 752097, time.UTC)
	donor := Donor{
		Address:   "1234",
		ExpiresAt: time,
	}
	expected := `{"address":"1234","expires_at":"2020-11-19T19:19:19.000752097Z"}`
	jsonB, _ := json.Marshal(donor)
	if expected != string(jsonB) {
		t.Errorf("Expected %s but got %s", expected, string(jsonB))
	}
}

func TestUnSerializeDonor(t *testing.T) {
	donorStr := `{"address":"1234","expires_at":"2020-11-19T19:19:19.000752097Z"}`
	var donor Donor
	json.Unmarshal([]byte(donorStr), &donor)
	if donor.Address != "1234" {
		t.Errorf("Expected address %s but got %s", "1234", donor.Address)
	}
	if donor.ExpiresAt.Year() != 2020 || donor.ExpiresAt.Month() != 11 || donor.ExpiresAt.Day() != 19 || donor.ExpiresAt.Hour() != 19 || donor.ExpiresAt.Minute() != 19 || donor.ExpiresAt.Second() != 19 {
		t.Error("Bad date after unmarshal")
	}
}
