package image

import (
	"sync"

	"github.com/appditto/natricon/server/db"
	"github.com/appditto/natricon/server/spc"
)

// BadgeService is a singleton providing badge/address data
type badgeService struct {
	PrincipalReps map[string]bool
	Exchanges     map[string]bool
	Services      map[string]bool
}

var bsingleton *badgeService
var bonce sync.Once

func GetBadgeSvc() *badgeService {
	bonce.Do(func() {
		// Grab cached principal reps
		principalReps := db.GetDB().GetPrincipalReps()
		// Translate everything into a map since
		// Translate all data into maps since lookup is O(1) instead of O(n)
		prMap := map[string]bool{}
		for i := 0; i < len(principalReps); i++ {
			prMap[principalReps[i]] = true
		}
		// Exchanges
		exchMap := map[string]bool{}
		for i := 0; i < len(spc.Exchanges); i++ {
			exchMap[spc.Exchanges[i]] = true
		}
		// Services
		svcMap := map[string]bool{}
		for i := 0; i < len(spc.Services); i++ {
			svcMap[spc.Services[i]] = true
		}
		bsingleton = &badgeService{
			PrincipalReps: prMap,
			Exchanges:     exchMap,
			Services:      svcMap,
		}
	})
	return bsingleton
}

// UpdatePrincipalReps - Update principal rep map
func (sm *badgeService) UpdatePrincipalReps(reps []string) {
	prMap := map[string]bool{}
	for i := 0; i < len(reps); i++ {
		prMap[reps[i]] = true
	}
	sm.PrincipalReps = prMap
}

// Getspc.BadgeType - Return badge type for a given PK
func (sm *badgeService) GetBadgeType(pk string) spc.BadgeType {
	if sm.Services[pk] {
		// Service
		return spc.BTService
	} else if sm.Exchanges[pk] {
		// Exchange
		return spc.BTExchange
	} else if sm.PrincipalReps[pk] {
		// Principal Rep
		return spc.BTNode
	} else if db.GetDB().HasDonorStatus(pk) {
		// Donor
		return spc.BTDonor
	}
	return spc.BTNone
}
