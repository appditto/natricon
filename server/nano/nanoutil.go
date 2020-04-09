package nano

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"

	"github.com/bbedward/nano/address"
	"github.com/bbedward/nano/types"
)

var nanoRegex = regexp.MustCompile("(?:xrb|nano)(?:_)(?:1|3)(?:[13456789abcdefghijkmnopqrstuwxyz]{59})")

func GenerateAddress() string {
	pub, _ := address.GenerateKey()
	return string(address.PubKeyToAddress(pub))
}

// ValidateAddress - Returns true if a nano address is valid
func ValidateAddress(account string) bool {
	if !nanoRegex.MatchString(account) {
		return false
	}
	return address.ValidateAddress(types.Account(account))
}

// AddressSha256 - Hashes an address excluding prefix
func AddressSha256(account string, seed string) string {
	var prefixRemoved string
	if len(account) == 64 {
		prefixRemoved = account[4:]
	} else if len(account) == 65 {
		prefixRemoved = account[5:]
	}
	hasher := sha256.New()
	hasher.Write([]byte(prefixRemoved))
	hasher.Write([]byte(seed))
	return hex.EncodeToString(hasher.Sum(nil))
}
