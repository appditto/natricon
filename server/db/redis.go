package db

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
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
func (r *redisManager) UpdateDonorStatus(hash string, account string, durationDays uint, maxDays uint) {
	hashKey := fmt.Sprintf("%s:processedHashes", keyPrefix)
	key := fmt.Sprintf("%s:donator:%s", keyPrefix, account)
	// See if this hash was already processed
	_, err := r.hget(hashKey, hash)
	if err == nil {
		glog.Infof("Hash already processed %s", hash)
		return
	}
	// Get current donator if exists
	cur, err := r.get(key)
	var donator *Donator
	if err == nil {
		json.Unmarshal([]byte(cur), donator)
	}
	// Calculate new expiry
	curDate := time.Now().UTC()
	existingHours := 0.0
	if donator != nil {
		existingHours = donator.ExpiresAt.Sub(curDate).Hours()
		if existingHours < 0 {
			existingHours = 0.0
		}
	}
	// Calculate newExpiry
	newExpiryHours := time.Duration(math.Min(float64(maxDays*24), existingHours+float64(durationDays*24)))
	newExpiry := curDate.Add(newExpiryHours * time.Hour)
	// Set new donator
	newDonor := Donator{
		Address:   account,
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

// HasDonorStatus - check if an account has donor status
func (r *redisManager) HasDonorStatus(account string) bool {
	account = strings.ReplaceAll(account, "xrb_", "nano_")
	key := fmt.Sprintf("%s:donator:%s", keyPrefix, account)
	raw, err := r.get(key)
	if err != nil {
		return false
	}
	var donator *Donator
	json.Unmarshal([]byte(raw), donator)
	// See if expired
	curDate := time.Now().UTC()
	if donator.ExpiresAt.Sub(curDate).Seconds() < 0 {
		r.del(key)
		return false
	}
	return true
}
