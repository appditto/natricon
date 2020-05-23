package db

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/appditto/natricon/server/utils"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v7"
	"github.com/golang/glog"
)

// Prefix for all keys
const keyPrefix = "natricon"

// Singleton to keep assets loaded in memory
type redisManager struct {
	Client *redis.Client
	Locker *redislock.Client
}

var singleton *redisManager
var once sync.Once

func GetDB() *redisManager {
	once.Do(func() {
		redis_port, err := strconv.Atoi(utils.GetEnv("REDIS_PORT", "6379"))
		if err != nil {
			panic("Invalid REDIS_PORT specified")
		}
		redis_db, err := strconv.Atoi(utils.GetEnv("REDIS_DB", "0"))
		if err != nil {
			panic("Invalid REDIS_DB specified")
		}
		client := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", utils.GetEnv("REDIS_HOST", "localhost"), redis_port),
			DB:   redis_db,
		})
		// Create locker
		// Create object
		singleton = &redisManager{
			Client: client,
			Locker: redislock.New(client),
		}
	})
	return singleton
}

// del - Redis DEL
func (r *redisManager) del(key string) (int64, error) {
	val, err := r.Client.Del(key).Result()
	return val, err
}

// get - Redis GET
func (r *redisManager) get(key string) (string, error) {
	val, err := r.Client.Get(key).Result()
	return val, err
}

// set - Redis SET
func (r *redisManager) set(key string, value string) error {
	err := r.Client.Set(key, value, 0).Err()
	return err
}

// hget - Redis HGET
func (r *redisManager) hget(key string, field string) (string, error) {
	val, err := r.Client.HGet(key, field).Result()
	return val, err
}

// hset - Redis HSET
func (r *redisManager) hset(key string, field string, value string) error {
	err := r.Client.HSet(key, field, value, 0).Err()
	return err
}

// UpdateDonorStatus - Update donor status with given duration in days
func (r *redisManager) UpdateDonorStatus(hash string, acct string, durationDays uint, maxDays uint) {
	pubkey := utils.AddressToPub(acct)
	hashKey := fmt.Sprintf("%s:processedHashes", keyPrefix)
	key := fmt.Sprintf("%s:donor:%s", keyPrefix, pubkey)
	// See if this hash was already processed
	_, err := r.hget(hashKey, hash)
	if err == nil {
		glog.Infof("Hash already processed %s", hash)
		return
	}
	// Get current donor if exists
	cur, err := r.get(key)
	var donor *Donor
	if err == nil {
		json.Unmarshal([]byte(cur), donor)
	}
	// Calculate new expiry
	curDate := time.Now().UTC()
	existingHours := 0.0
	if donor != nil {
		existingHours = donor.ExpiresAt.Sub(curDate).Hours()
		if existingHours < 0 {
			existingHours = 0.0
		}
	}
	// Calculate newExpiry
	newExpiryHours := time.Duration(math.Min(float64(maxDays*24), existingHours+float64(durationDays*24)))
	newExpiry := curDate.Add(newExpiryHours * time.Hour)
	// Set new donor
	newDonor := Donor{
		PubKey:    pubkey,
		ExpiresAt: newExpiry,
	}
	// Marshal
	marshaled, err := json.Marshal(newDonor)
	if err != nil {
		glog.Errorf("Couldn't serialize donor %s", err)
		return
	}
	// Save new status
	r.set(key, string(marshaled))
	r.hset(hashKey, hash, "1")
}

// HasDonorStatus - check if a public key has donor status
func (r *redisManager) HasDonorStatus(pubkey string) bool {
	key := fmt.Sprintf("%s:donor:%s", keyPrefix, pubkey)
	raw, err := r.get(key)
	if err != nil {
		return false
	}
	var donor Donor
	err = json.Unmarshal([]byte(raw), &donor)
	if err != nil {
		glog.Errorf("Error unmarshalling donor json %s", err)
		return false
	}
	// See if expired
	curDate := time.Now().UTC()
	if donor.ExpiresAt.Sub(curDate).Seconds() < 0 {
		r.del(key)
		return false
	}
	return true
}

// SetPrincipalRepRequirement - set voting weight requirement to be principal rep
func (r *redisManager) SetPrincipalRepRequirement(amount float64) {
	key := fmt.Sprintf("%s:principal_rep_requirement", keyPrefix)
	r.set(key, fmt.Sprintf("%f", amount))
}

// GetPrincipalRepRequirement - get voting weight requirement to be principal rep
func (r *redisManager) GetPrincipalRepRequirement() float64 {
	key := fmt.Sprintf("%s:principal_rep_requirement", keyPrefix)
	amount, err := r.get(key)
	if err != nil {
		// Return approximation
		return 94737.0
	}
	converted, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		// Return approximation
		return 94737.0
	}
	return converted
}

// SetPrincipalReps - Cache principal reps
func (r *redisManager) SetPrincipalReps(reps []string) {
	key := fmt.Sprintf("%s:principal_reps", keyPrefix)
	marshalled, err := json.Marshal(reps)
	if err != nil {
		r.set(key, string(marshalled))
	}
}

// GetPrincipalReps - Get cached principal reps
func (r *redisManager) GetPrincipalReps() []string {
	key := fmt.Sprintf("%s:principal_reps", keyPrefix)
	reps, err := r.get(key)
	if err != nil {
		// Return empty set
		return []string{}
	}
	var repsU []string
	err = json.Unmarshal([]byte(reps), &repsU)
	if err != nil {
		return []string{}
	}
	return repsU
}
