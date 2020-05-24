package image

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
)

type IllustrationType string
type Sex string

const (
	Body         IllustrationType = "body"
	BodyOutline  IllustrationType = "body-outline"
	Badge        IllustrationType = "badge"
	Hair         IllustrationType = "hair-front"
	HairBack     IllustrationType = "hair-back"
	HairOutline  IllustrationType = "hair-outline"
	Mouth        IllustrationType = "mouth"
	MouthOutline IllustrationType = "mouth-outline"
	Eye          IllustrationType = "eyes"
	Male         Sex              = "M"
	Female       Sex              = "F"
	Neutral      Sex              = "N"
)

type Asset struct {
	FileName         string           // File name of asset
	IllustrationPath string           // Full path of illustration on the file system
	Type             IllustrationType // Type of illustration (body, hair, mouth, eye)
	SVGContents      []byte           // Full contents of SVG asset
	HairColored      bool             // Whether this asset should be colored the same as hair color
	BodyColored      bool             // Whether this asset should be colored the same as body color
	Sex              Sex              // The Sex condition of this asset
	LightOnly        bool             // Whether this asset can only be used on light colors
	DarkColored      bool             // Whether this asset gets adjusted on dark colors
	DarkBWColored    bool             // Whether this asset has a secondary color adjustmetn on dark backgrounds
	BLK299           bool             // Opacity replacements for _blk299 assets
}

// getIllustrationPath - get full path of image
func getIllustrationPath(illustration string, iType IllustrationType) string {
	wd, err := os.Getwd()
	if err != nil {
		panic("Can't get working directory")
	}

	fPath := path.Join(wd, "assets", "illustrations", string(iType), illustration)
	if _, err := os.Stat(fPath); err != nil {
		// File does not exist
		panic(fmt.Sprintf("File %s does not exist", fPath))
	}
	return fPath
}

// getSex - get Sex based on image name
func getSex(name string) Sex {
	if strings.Contains(name, "_f") {
		return Female
	} else if strings.Contains(name, "_m") {
		return Male
	}
	return Neutral
}

// Singleton to keep assets loaded in memory
type assetManager struct {
	bodyAssets         []Asset
	bodyOutlineAssets  []Asset
	donorBadgeAssets   []Asset
	exchBadgeAssets    []Asset
	nodeBadgeAssets    []Asset
	svcBadgeAssets     []Asset
	hairAssets         []Asset
	hairBackAssets     []Asset
	hairOutlineAssets  []Asset
	mouthAssets        []Asset
	mouthOutlineAssets []Asset
	eyeAssets          []Asset
}

var singleton *assetManager
var once sync.Once

func GetAssets() *assetManager {
	once.Do(func() {
		var err error
		// Load body assets
		var bodyAssets []Asset
		for _, ba := range BodyIllustrations {
			var a Asset
			err = json.Unmarshal(ba, &a)
			bodyAssets = append(bodyAssets, a)
		}
		// Load body outlines
		var bodyOutlineAssets []Asset
		for _, ba := range BodyOutlineIllustrations {
			var a Asset
			err = json.Unmarshal(ba, &a)
			bodyOutlineAssets = append(bodyOutlineAssets, a)
		}
		// Load badges
		var donorBadgeAssets []Asset
		for _, ba := range DonorBadgeIllustrations {
			var a Asset
			err = json.Unmarshal(ba, &a)
			donorBadgeAssets = append(donorBadgeAssets, a)
		}
		var exchBadgeAssets []Asset
		for _, ba := range ExchangeBadgeIllustrations {
			var a Asset
			err = json.Unmarshal(ba, &a)
			exchBadgeAssets = append(exchBadgeAssets, a)
		}
		var nodeBadgeAssets []Asset
		for _, ba := range NodeBadgeIllustrations {
			var a Asset
			err = json.Unmarshal(ba, &a)
			nodeBadgeAssets = append(nodeBadgeAssets, a)
		}
		var svcBadgeAssets []Asset
		for _, ba := range ServiceBadgeIllustrations {
			var a Asset
			err = json.Unmarshal(ba, &a)
			svcBadgeAssets = append(svcBadgeAssets, a)
		}
		// Load hair assets
		var hairAssets []Asset
		for _, ha := range HairIllustrations {
			var a Asset
			err = json.Unmarshal(ha, &a)
			hairAssets = append(hairAssets, a)
		}
		// Load hair back assets
		var hairBackAssets []Asset
		for _, ha := range HairBackIllustrations {
			var a Asset
			err = json.Unmarshal(ha, &a)
			hairBackAssets = append(hairBackAssets, a)
		}
		// Hair outline
		var hairOutlineAssets []Asset
		for _, ha := range HairOutlineIllustrations {
			var a Asset
			err = json.Unmarshal(ha, &a)
			hairOutlineAssets = append(hairOutlineAssets, a)
		}
		// Load mouth assets
		var mouthAssets []Asset
		for _, ma := range MouthIllustrations {
			var a Asset
			err = json.Unmarshal(ma, &a)
			mouthAssets = append(mouthAssets, a)
		}
		// Mouth Outline
		var mouthOutlineAssets []Asset
		for _, ma := range MouthOutlineIllustrations {
			var a Asset
			err = json.Unmarshal(ma, &a)
			mouthOutlineAssets = append(mouthOutlineAssets, a)
		}
		// Load eye assets
		var eyeAssets []Asset
		for _, ea := range EyeIllustrations {
			var a Asset
			err = json.Unmarshal(ea, &a)
			eyeAssets = append(eyeAssets, a)
		}
		if err != nil {
			panic("Failed to decode assets")
		}
		// Create object
		singleton = &assetManager{
			bodyAssets:         bodyAssets,
			bodyOutlineAssets:  bodyOutlineAssets,
			donorBadgeAssets:   donorBadgeAssets,
			exchBadgeAssets:    exchBadgeAssets,
			nodeBadgeAssets:    nodeBadgeAssets,
			svcBadgeAssets:     svcBadgeAssets,
			hairAssets:         hairAssets,
			hairBackAssets:     hairBackAssets,
			hairOutlineAssets:  hairOutlineAssets,
			mouthAssets:        mouthAssets,
			mouthOutlineAssets: mouthOutlineAssets,
			eyeAssets:          eyeAssets,
		}
	})
	return singleton
}

// GetNBodyAssets - get # of body assets
func (sm *assetManager) GetNBodyAssets() int {
	return len(sm.bodyAssets)
}

// GetBodyAssets - get complete list of hair assets
func (sm *assetManager) GetBodyAssets() []Asset {
	return sm.bodyAssets
}

// GetBodyOutlineAssets
func (sm *assetManager) GetBodyOutlineAssets() []Asset {
	return sm.bodyOutlineAssets
}

// GetNHairAssets - get # of hair assets
func (sm *assetManager) GetNHairAssets() int {
	return len(sm.hairAssets)
}

// GetBadges - get badge assets
func (sm *assetManager) GetBadgeAssets(btype BadgeType) []Asset {
	switch btype {
	case BTDonor:
		return sm.donorBadgeAssets
	case BTExchange:
		return sm.exchBadgeAssets
	case BTNode:
		return sm.nodeBadgeAssets
	case BTService:
		return sm.svcBadgeAssets
	default:
		return sm.donorBadgeAssets
	}
}

// GetHairAssets - get complete list of hair assets
func (sm *assetManager) GetHairAssets(sex Sex) []Asset {
	var ret []Asset
	for _, v := range sm.hairAssets {
		if sex == Neutral {
			ret = append(ret, v)
		} else if v.Sex == sex || v.Sex == Neutral {
			ret = append(ret, v)
		}
	}
	return ret
}

// GetHairOutlineAssets
func (sm *assetManager) GetHairOutlineAssets() []Asset {
	return sm.hairOutlineAssets
}

// GetBackHairAssets - get complete list of hair assets
func (sm *assetManager) GetBackHairAssets() []Asset {
	return sm.hairBackAssets
}

// GetMouthAssets - Get mouth assets
func (sm *assetManager) GetMouthAssets(sex Sex, luminosity float64) []Asset {
	var ret []Asset
	luminosityInt := int(luminosity)
	for _, v := range sm.mouthAssets {
		if LightToDarkSwitchPoint > luminosityInt && v.LightOnly {
			continue
		}
		if sex == Neutral {
			ret = append(ret, v)
		} else if v.Sex == sex || v.Sex == Neutral {
			ret = append(ret, v)
		}
	}
	return ret
}

// GetMouthOutlineAssets
func (sm *assetManager) GetMouthOutlineAssets() []Asset {
	return sm.mouthOutlineAssets
}

// GetEyeAssets - Get eye asset list
func (sm *assetManager) GetEyeAssets(sex Sex, luminosity float64) []Asset {
	var ret []Asset
	luminosityInt := int(luminosity)
	for _, v := range sm.eyeAssets {
		if LightToDarkSwitchPoint > luminosityInt && v.LightOnly {
			continue
		}
		if sex == Neutral {
			ret = append(ret, v)
		} else if v.Sex == sex || v.Sex == Neutral {
			ret = append(ret, v)
		}
	}
	return ret
}
