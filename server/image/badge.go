package image

import "github.com/appditto/natricon/server/db"

type BadgeType string

const (
	BTNone     BadgeType = ""
	BTDonor    BadgeType = "donor"
	BTExchange BadgeType = "exchange"
	BTNode     BadgeType = "node"
	BTService  BadgeType = "service"
)

// Exchanges
var Exchanges = []string{
	"16aa39b37529b7bb50f345ce97e1b34088c1930b973eedd4b2301943a2c001da",
	"d368b6c13ad91139e933e310d5c1218add1909fdeb46cb50e2fcaa9e9a24d047",
	"4b9fca980437c128235bcd0b1d4a5f1c2dc51d8abb5703710e44980d8c8d1c83",
	"c798cff4f1131204f65c4d22c3e6316f26f380ee0616aadbabea1268fd75fb05",
	"f45b8087702b867f9736ae82628708e57780b1eb004e123cf2822a5cb935af17",
	"095b645b6c0cccb52dd65218de613ce13cea58a850a80c3f704291b698a50417",
	"f614e657a196f51a55ecd55200d75317076ed96b984cc2123dae64e325e9dcec",
	"aadf2d5e7be0692d52952466216b3bb4cccc3a2cad63e126f34a52a22717269c",
	"bf6822000278519a5886903f81274339a3c8a8167ba80dc531c1284c0969b1cd",
	"3c20b242915b7fdf8f08a85eda94c285cce2c29987f6dba1711329e669fe194e",
	"85b8bdaff3c70bd1870b4ab44c67465d49f3a0af254d22ad8aa25cb935959393",
	"e3c750bfecc71a25505494656eb475bbfc11a8d87fd3fc0c76aaac040938eaf3",
	"c28e28a213a462130fc17b1bc1dbfdf1ccc940b5d4c143dd997dff4203ed7f05",
}

// Services
var Services = []string{
	"c58384724ee9dae70fabc3d357caf4f40cf4eaaf7a68e5ee104d093ce76af05a",
	"2f087568a509807680c666813f1156d32a226bd8e29fd03a67d7382e6968dfc5",
	"5a70e35e5ad2faf9523fa5d2fb44f7ac055f8fbd3d7126dbfa7da9f8928cf1ca",
	"52148a0b8784ac7b6c7dea55c0e1d9f7ce34a6c3369f50bfaf00438254f349a3",
	"e2a9aa91f7b37d66871d81da8344a6364f6a77697a4a395ed452ad30cf061d11",
	"511ac43730543f18c07836bb2f61032b16eda46f10779ca0f330c9b663881060",
	"559487622beacd53269720f2bf1ce924c5e56b0537c35b4d47afdcc4718bf645",
	"022d4fc557a97d8c7bf34b680df8127b2ff2997a76cc52e3b7a001fe5ead2fe2",
	"d154db1790b9f28aaa20acc72048cb04fecfc8c84848c6e2d18ddb09e05700c2",
	"e382dd09ec8cafc2427cf817e9afe1f372ce81085ab4feb1f3de1f25ee818e5d",
	"60e6f8e0017f59c8ce5447c1f1e951cad302661dc40e44c4ecea2f7f835d3e7b",
	"cb5e4fffe00bfd72495ca5ea50063b7245ae53368dc978db5285396ecfcdc3e3",
	"8ad883d7de7c3b15b26bc69400ac5dc7cce4abd973dd9876420e1a361c3c2efc",
	"f4aa8c2b743dd91dacf67b24b62d7d24585ca9262cc153da72a6ba06e984ae48",
	"caeae4206c202abac3ccb7ac89a9f72961c5f73062a2aec400491a271521d583",
	"d4bbfa50649d80e5f63fc396c6f4cf6321cabd7c1480e964c2701d56aafeb5e3",
	"511ac43730543f18c07836bb2f61032b16eda46f10779ca0f330c9b663881060",
	"e315b46176f6d3c6255ab222bea7305b6cd848d3b9f0a59f51332d5d70629868",
	"2994d330022a052df83e10fce1b3e140496cdcd7e0c0f2ff6de2670291b88011",
	"69f0a3b369c2d66d1cac6a40ab561df1ba6b69b15f67ec91ba9ff286d9624254",
	"1793e59c41d19b79b66134e76129d53446fd3794882563788437482e356f0a87",
}

// BadgeService is a singleton providing badge/address data
type badgeService struct {
	principalReps []string
}

var bsingleton *badgeService

func GetBadgeSvc() *badgeService {
	once.Do(func() {
		// Grab cached principal reps
		// TODO
		principalReps := []string{}
		// Create object
		bsingleton = &badgeService{
			principalReps: principalReps,
		}
	})
	return bsingleton
}

// GetBadgeType - Return badge type for a given PK
func (sm *badgeService) GetBadgeType(pk string) BadgeType {
	// Exchange
	for _, a := range Exchanges {
		if a == pk {
			return BTExchange
		}
	}
	// Service
	for _, a := range Services {
		if a == pk {
			return BTService
		}
	}
	// Donor
	if db.GetDB().HasDonorStatus(pk) {
		return BTDonor
	}
	// TODO - implement Node type
	return BTNone
}
