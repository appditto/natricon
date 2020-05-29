package spc

import (
	"github.com/appditto/natricon/server/color"
)

// Badge types
type BadgeType string

const (
	BTNone     BadgeType = ""
	BTDonor    BadgeType = "donor"
	BTExchange BadgeType = "exchange"
	BTNode     BadgeType = "node"
	BTService  BadgeType = "service"
)

// Vanity
type Vanity struct {
	// Optional fields
	Hash  string    // Will generate the natricon with specific hash
	Badge BadgeType // Will generate natricon with specified badge
	// If using any of the below then ALL of them are required
	BodyColor    *color.RGB
	HairColor    *color.RGB
	BodyAssetID  int
	HairAssetID  int
	MouthAssetID int
	EyeAssetID   int
}

// Stats
type StatsService string
