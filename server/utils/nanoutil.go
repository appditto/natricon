package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"regexp"

	"github.com/bbedward/nano/address"
	"github.com/bbedward/nano/types"
)

const rawPerNanoStr = "1000000000000000000000000000000"

var rawPerNano, _ = new(big.Float).SetString(rawPerNanoStr)

const nanoPrecision = 1000000 // 0.000001 NANO precision

const nanoRegexStr = "(?:xrb|nano)(?:_)(?:1|3)(?:[13456789abcdefghijkmnopqrstuwxyz]{59})"

var nanoRegex = regexp.MustCompile(nanoRegexStr)

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

// RawToNano - Converts Raw amount to usable Nano amount
func RawToNano(raw string) (float64, error) {
	rawBig, ok := new(big.Float).SetString(raw)
	if !ok {
		err := errors.New(fmt.Sprintf("Unable to convert %s to int", raw))
		return -1, err
	}
	asNano := rawBig.Quo(rawBig, rawPerNano)
	// Truncate precision beyond 0.000001
	bf := big.NewFloat(0).SetPrec(1000000).Set(asNano)
	bu := big.NewFloat(0).SetPrec(1000000).SetFloat64(0.000001)

	bf.Quo(bf, bu)

	// Truncate:
	i := big.NewInt(0)
	bf.Int(i)
	bf.SetInt(i)

	f, _ := bf.Mul(bf, bu).Float64()
	return f, nil
}
