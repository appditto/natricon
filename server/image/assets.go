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
	Body    IllustrationType = "Body"
	Hair    IllustrationType = "Hair"
	Mouth   IllustrationType = "Mouth"
	Eye     IllustrationType = "Eyes"
	Male    Sex              = "M"
	Female  Sex              = "F"
	Neutral Sex              = "N"
)

type Asset struct {
	IllustrationPath string           // Full path of illustration on the file system
	Type             IllustrationType // Type of illustration (body, hair, mouth, eye)
	SVGContents      []byte           // Full contents of SVG asset
	hairColored      bool             // Whether this asset should be colored the same as hair color
	bodyColored      bool             // Whether this asset should be colored the same as body color
}

var bodyIllustrations = [...]string{
	"NarrowRound-10.svg",
	"NarrowRound-11.svg",
	"NarrowRound-12.svg",
	"NarrowRound-13.svg",
	"NarrowRound-14.svg",
	"NarrowRound-15.svg",
	"NarrowRound-5.svg",
	"NarrowRound-6.svg",
	"NarrowRound-7.svg",
	"NarrowRound-8.svg",
	"NarrowRound-9.svg",
	"Square-100.svg",
	"Square-104.svg",
	"Square-108.svg",
	"Square-112.svg",
	"Square-116.svg",
	"Square-120.svg",
	"Square-124.svg",
	"Square-128.svg",
	"Square-48.svg",
	"Square-52.svg",
	"Square-56.svg",
	"Square-60.svg",
	"Square-64.svg",
	"Square-68.svg",
	"Square-72.svg",
	"Square-76.svg",
	"Square-80.svg",
	"Square-84.svg",
	"Square-88.svg",
	"Square-92.svg",
	"Square-96.svg",
}

var hairIllustrations = [...]string{
	"Bubble-1.svg",
	"Slick.svg",
	"Weird.svg",
}

var eyeIllustrations = [...]string{
	"Eyeglasses-1.svg",
	"Eyeglasses-2.svg",
	"Eyeglasses-3.svg",
	"Eyeglasses-4.svg",
	"Eyes-1.svg",
}

var mouthIllustrations = [...]string{
	"Mustache-Slick.svg",
	"Smile-Bigger.svg",
	"Smile-Simple.svg",
	"Smile-Teeth.svg",
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

// Singleton to keep assets loaded in memory
type assetManager struct {
	bodyAssets  [len(bodyIllustrations)]Asset
	hairAssets  [len(hairIllustrations)]Asset
	mouthAssets [len(mouthIllustrations)]Asset
	eyeAssets   [len(eyeIllustrations)]Asset
}

var singleton *assetManager
var once sync.Once

func GetAssets() *assetManager {
	once.Do(func() {
		var err error
		// Load body assets
		var bodyAssets [len(bodyIllustrations)]Asset
		for i, ba := range bodyIllustrations {
			bodyAssets[i] = Asset{}
			bodyAssets[i].IllustrationPath = getIllustrationPath(ba, Body)
			bodyAssets[i].SVGContents, err = ioutil.ReadFile(bodyAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", bodyAssets[i].IllustrationPath)
				panic(err.Error())
			}
			bodyAssets[i].hairColored = false
			bodyAssets[i].bodyColored = true
		}
		// Load hair assets
		var hairAssets [len(hairIllustrations)]Asset
		for i, ha := range hairIllustrations {
			hairAssets[i] = Asset{}
			hairAssets[i].IllustrationPath = getIllustrationPath(ha, Hair)
			hairAssets[i].SVGContents, err = ioutil.ReadFile(hairAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", hairAssets[i].IllustrationPath)
				panic(err.Error())
			}
			hairAssets[i].hairColored = true
			hairAssets[i].bodyColored = false
		}
		// Load mouth assets
		var mouthAssets [len(mouthIllustrations)]Asset
		for i, ma := range mouthIllustrations {
			mouthAssets[i] = Asset{}
			mouthAssets[i].IllustrationPath = getIllustrationPath(ma, Mouth)
			mouthAssets[i].SVGContents, err = ioutil.ReadFile(mouthAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", mouthAssets[i].IllustrationPath)
				panic(err.Error())
			}
			mouthAssets[i].hairColored = strings.Contains(ma, "_hc_")
			mouthAssets[i].bodyColored = false
		}
		// Load eye assets
		var eyeAssets [len(eyeIllustrations)]Asset
		for i, ea := range eyeIllustrations {
			eyeAssets[i] = Asset{}
			eyeAssets[i].IllustrationPath = getIllustrationPath(ea, Eye)
			eyeAssets[i].SVGContents, err = ioutil.ReadFile(eyeAssets[i].IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", eyeAssets[i].IllustrationPath)
				panic(err.Error())
			}
			eyeAssets[i].hairColored = false
			eyeAssets[i].bodyColored = false
		}
		// Create object
		singleton = &assetManager{
			bodyAssets:  bodyAssets,
			hairAssets:  hairAssets,
			mouthAssets: mouthAssets,
			eyeAssets:   eyeAssets,
		}
	})
	return singleton
}

// GetNBodyAssets - get # of body assets
func (sm *assetManager) GetNBodyAssets() int {
	return len(sm.bodyAssets)
}

// GetBodyAssets - get complete list of hair assets
func (sm *assetManager) GetBodyAssets() [len(bodyIllustrations)]Asset {
	return sm.bodyAssets
}

// GetNHairAssets - get # of hair assets
func (sm *assetManager) GetNHairAssets() int {
	return len(sm.hairAssets)
}

// GetHairAssets - get complete list of hair assets
func (sm *assetManager) GetHairAssets() [len(hairIllustrations)]Asset {
	return sm.hairAssets
}

// GetNMouthAssets - get # of mouth assets
func (sm *assetManager) GetNMouthAssets() int {
	return len(sm.mouthAssets)
}

// GetMouthAssets - Get mouth assets
func (sm *assetManager) GetMouthAssets() [len(mouthIllustrations)]Asset {
	return sm.mouthAssets
}

// GetNEyeAssets - Get # of eye assets
func (sm *assetManager) GetNEyeAssets() int {
	return len(sm.eyeAssets)
}

// GetEyeAssets - Get eye asset list
func (sm *assetManager) GetEyeAssets() [len(eyeIllustrations)]Asset {
	return sm.eyeAssets
}
