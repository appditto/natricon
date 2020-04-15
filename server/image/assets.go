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
	IllustrationPath string           // Full path of illustration on the file system
	Type             IllustrationType // Type of illustration (body, hair, mouth, eye)
	SVGContents      []byte           // Full contents of SVG asset
	hairColored      bool             // Whether this asset should be colored the same as hair color
	bodyColored      bool             // Whether this asset should be colored the same as body color
	sex              Sex              // The sex condition of this asset
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

// getSex - get sex based on image name
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
			bodyAssets[i].IllustrationPath = getIllustrationPath(ba, Body)
			bodyAssets[i].SVGContents, err = ioutil.ReadFile(bodyAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", bodyAssets[i].IllustrationPath)
				panic(err.Error())
			}
			bodyAssets[i].hairColored = false
			bodyAssets[i].bodyColored = true
			bodyAssets[i].sex = Neutral
		}
		// Load hair assets
		var hairAssets [len(HairIllustrations)]Asset
		for i, ha := range HairIllustrations {
			hairAssets[i] = Asset{}
			hairAssets[i].IllustrationPath = getIllustrationPath(ha, Hair)
			hairAssets[i].SVGContents, err = ioutil.ReadFile(hairAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", hairAssets[i].IllustrationPath)
				panic(err.Error())
			}
			hairAssets[i].hairColored = true
			hairAssets[i].bodyColored = false
			hairAssets[i].sex = getSex(ha)
		}
		// Load hair back assets
		var hairBackAssets [len(HairBackIllustrations)]Asset
		for i, ha := range HairIllustrations {
			hairBackAssets[i] = Asset{}
			hairBackAssets[i].IllustrationPath = getIllustrationPath(ha, HairBack)
			hairBackAssets[i].SVGContents, err = ioutil.ReadFile(hairBackAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", hairBackAssets[i].IllustrationPath)
				panic(err.Error())
			}
			hairBackAssets[i].hairColored = true
			hairBackAssets[i].bodyColored = false
			hairBackAssets[i].sex = getSex(ha)
		}
		// Load mouth assets
		var mouthAssets [len(MouthIllustrations)]Asset
		for i, ma := range MouthIllustrations {
			mouthAssets[i] = Asset{}
			mouthAssets[i].IllustrationPath = getIllustrationPath(ma, Mouth)
			mouthAssets[i].SVGContents, err = ioutil.ReadFile(mouthAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", mouthAssets[i].IllustrationPath)
				panic(err.Error())
			}
			mouthAssets[i].hairColored = strings.Contains(ma, "_hc")
			mouthAssets[i].bodyColored = false
			mouthAssets[i].sex = getSex(ma)
		}
		// Load eye assets
		var eyeAssets [len(EyeIllustrations)]Asset
		for i, ea := range EyeIllustrations {
			eyeAssets[i] = Asset{}
			eyeAssets[i].IllustrationPath = getIllustrationPath(ea, Eye)
			eyeAssets[i].SVGContents, err = ioutil.ReadFile(eyeAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", eyeAssets[i].IllustrationPath)
				panic(err.Error())
			}
			eyeAssets[i].hairColored = false
			eyeAssets[i].bodyColored = false
			eyeAssets[i].sex = getSex(ea)
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

// GetNMouthAssets - get # of mouth assets
func (sm *assetManager) GetNMouthAssets() int {
	return len(sm.mouthAssets)
}

// GetMouthAssets - Get mouth assets
func (sm *assetManager) GetMouthAssets() [len(MouthIllustrations)]Asset {
	return sm.mouthAssets
}

// GetNEyeAssets - Get # of eye assets
func (sm *assetManager) GetNEyeAssets() int {
	return len(sm.eyeAssets)
}

// GetEyeAssets - Get eye asset list
func (sm *assetManager) GetEyeAssets() [len(EyeIllustrations)]Asset {
	return sm.eyeAssets
}
