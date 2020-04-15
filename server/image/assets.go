package image

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/golang/glog"
)

type IllustrationType string
type Sex string

const (
	Body     IllustrationType = "Body"
	Hair     IllustrationType = "hair-front"
	HairBack IllustrationType = "hair-back"
	Mouth    IllustrationType = "Mouth"
	Eye      IllustrationType = "Eyes"
	Male     Sex              = "M"
	Female   Sex              = "F"
	Neutral  Sex              = "N"
)

type Asset struct {
	FileName         string           // File name of asset
	IllustrationPath string           // Full path of illustration on the file system
	Type             IllustrationType // Type of illustration (body, hair, mouth, eye)
	SVGContents      []byte           // Full contents of SVG asset
	HairColored      bool             // Whether this asset should be colored the same as hair color
	BodyColored      bool             // Whether this asset should be colored the same as body color
	Sex              Sex              // The Sex condition of this asset
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
	bodyAssets     [len(BodyIllustrations)]Asset
	hairAssets     [len(HairIllustrations)]Asset
	hairBackAssets [len(HairBackIllustrations)]Asset
	mouthAssets    [len(MouthIllustrations)]Asset
	eyeAssets      [len(EyeIllustrations)]Asset
}

var singleton *assetManager
var once sync.Once

func GetAssets() *assetManager {
	once.Do(func() {
		var err error
		// Load body assets
		var bodyAssets [len(BodyIllustrations)]Asset
		for i, ba := range BodyIllustrations {
			bodyAssets[i] = Asset{}
			bodyAssets[i].FileName = ba
			bodyAssets[i].IllustrationPath = getIllustrationPath(ba, Body)
			bodyAssets[i].SVGContents, err = ioutil.ReadFile(bodyAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", bodyAssets[i].IllustrationPath)
				panic(err.Error())
			}
			bodyAssets[i].HairColored = false
			bodyAssets[i].BodyColored = true
			bodyAssets[i].Sex = Neutral
		}
		// Load hair assets
		var hairAssets [len(HairIllustrations)]Asset
		for i, ha := range HairIllustrations {
			hairAssets[i] = Asset{}
			hairAssets[i].FileName = ha
			hairAssets[i].IllustrationPath = getIllustrationPath(ha, Hair)
			hairAssets[i].SVGContents, err = ioutil.ReadFile(hairAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", hairAssets[i].IllustrationPath)
				panic(err.Error())
			}
			hairAssets[i].HairColored = true
			hairAssets[i].BodyColored = false
			hairAssets[i].Sex = getSex(ha)
		}
		// Load hair back assets
		var hairBackAssets [len(HairBackIllustrations)]Asset
		for i, ha := range HairBackIllustrations {
			hairBackAssets[i] = Asset{}
			hairBackAssets[i].FileName = ha
			hairBackAssets[i].IllustrationPath = getIllustrationPath(ha, HairBack)
			hairBackAssets[i].SVGContents, err = ioutil.ReadFile(hairBackAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", hairBackAssets[i].IllustrationPath)
				panic(err.Error())
			}
			hairBackAssets[i].HairColored = true
			hairBackAssets[i].BodyColored = false
			hairBackAssets[i].Sex = getSex(ha)
		}
		// Load mouth assets
		var mouthAssets [len(MouthIllustrations)]Asset
		for i, ma := range MouthIllustrations {
			mouthAssets[i] = Asset{}
			mouthAssets[i].FileName = ma
			mouthAssets[i].IllustrationPath = getIllustrationPath(ma, Mouth)
			mouthAssets[i].SVGContents, err = ioutil.ReadFile(mouthAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", mouthAssets[i].IllustrationPath)
				panic(err.Error())
			}
			mouthAssets[i].HairColored = strings.Contains(ma, "_hc")
			mouthAssets[i].BodyColored = false
			mouthAssets[i].Sex = getSex(ma)
		}
		// Load eye assets
		var eyeAssets [len(EyeIllustrations)]Asset
		for i, ea := range EyeIllustrations {
			eyeAssets[i] = Asset{}
			eyeAssets[i].FileName = ea
			eyeAssets[i].IllustrationPath = getIllustrationPath(ea, Eye)
			eyeAssets[i].SVGContents, err = ioutil.ReadFile(eyeAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", eyeAssets[i].IllustrationPath)
				panic(err.Error())
			}
			eyeAssets[i].HairColored = false
			eyeAssets[i].BodyColored = false
			eyeAssets[i].Sex = getSex(ea)
		}
		// Create object
		singleton = &assetManager{
			bodyAssets:     bodyAssets,
			hairAssets:     hairAssets,
			hairBackAssets: hairBackAssets,
			mouthAssets:    mouthAssets,
			eyeAssets:      eyeAssets,
		}
	})
	return singleton
}

// GetNBodyAssets - get # of body assets
func (sm *assetManager) GetNBodyAssets() int {
	return len(sm.bodyAssets)
}

// GetBodyAssets - get complete list of hair assets
func (sm *assetManager) GetBodyAssets() [len(BodyIllustrations)]Asset {
	return sm.bodyAssets
}

// GetNHairAssets - get # of hair assets
func (sm *assetManager) GetNHairAssets() int {
	return len(sm.hairAssets)
}

// GetHairAssets - get complete list of hair assets
func (sm *assetManager) GetHairAssets() [len(HairIllustrations)]Asset {
	return sm.hairAssets
}

// GetBackHairAssets - get complete list of hair assets
func (sm *assetManager) GetBackHairAssets() [len(HairBackIllustrations)]Asset {
	return sm.hairBackAssets
}

// GetMouthAssets - Get mouth assets
func (sm *assetManager) GetMouthAssets(sex Sex) []Asset {
	var ret []Asset
	for _, v := range sm.mouthAssets {
		if v.Sex == sex || v.Sex == Neutral {
			ret = append(ret, v)
		}
	}
	return ret
}

// GetEyeAssets - Get eye asset list
func (sm *assetManager) GetEyeAssets(sex Sex) []Asset {
	var ret []Asset
	for _, v := range sm.eyeAssets {
		if v.Sex == sex || v.Sex == Neutral {
			ret = append(ret, v)
		}
	}
	return ret
}
