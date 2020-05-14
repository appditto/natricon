package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/appditto/natricon/server/image"
	"github.com/golang/glog"
)

// getSex - get Sex based on image name
func getSex(name string) image.Sex {
	if strings.Contains(name, "_f") {
		return image.Female
	} else if strings.Contains(name, "_m") {
		return image.Male
	}
	return image.Neutral
}

// getLightOnly - get whether t his asset should only be available on light backgrounds
func getLightOnly(name string) bool {
	if strings.Contains(name, "_lod") || strings.Contains(name, "_ld") {
		return false
	}
	return true
}

// getDarkColored - get whether this asset should have colors altered
func getDarkColored(name string) bool {
	if strings.Contains(name, "_lod") {
		return true
	}
	return false
}

// getDarkBWColored - get whether this asset should have colors altered
func getDarkBWColored(name string) bool {
	if strings.Contains(name, "_lod_bw") {
		return true
	}
	return false
}

func getBlk299(name string) bool {
	if strings.Contains(name, "_blk299") {
		return true
	}
	return false
}

func LoadAssetsToArray() {
	wd, err := os.Getwd()
	if err != nil {
		panic("Can't get working directory")
	}

	ret := "package image\n\n"

	var bodyAsset image.Asset
	ret += "var BodyIllustrations = [][]byte{\n"
	fPath := path.Join(wd, "assets", "illustrations", string(image.Body))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			bodyAsset = image.Asset{}
			bodyAsset.FileName = info.Name()
			bodyAsset.IllustrationPath = path
			bodyAsset.SVGContents, err = ioutil.ReadFile(path)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", path)
				panic(err.Error())
			}
			bodyAsset.HairColored = false
			bodyAsset.BodyColored = true
			bodyAsset.LightOnly = false
			bodyAsset.DarkColored = false
			bodyAsset.DarkBWColored = false
			bodyAsset.Sex = getSex(info.Name())
			bodyAsset.BLK299 = false
			encoded, _ := json.Marshal(bodyAsset)
			ret += strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(encoded), "[", "{"), "]", "}"), " ", ", ") + ","
		}
		return nil
	})
	ret += "}\n"

	var bodyOutlineAsset image.Asset
	ret += "\nvar BodyOutlineIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", string(image.BodyOutline))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			bodyOutlineAsset = image.Asset{}
			bodyOutlineAsset.FileName = info.Name()
			bodyOutlineAsset.IllustrationPath = path
			bodyOutlineAsset.SVGContents, err = ioutil.ReadFile(path)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", path)
				panic(err.Error())
			}
			bodyOutlineAsset.HairColored = false
			bodyOutlineAsset.BodyColored = false
			bodyOutlineAsset.DarkBWColored = false
			bodyOutlineAsset.Sex = getSex(info.Name())
			bodyOutlineAsset.BLK299 = false
			encoded, _ := json.Marshal(bodyOutlineAsset)
			ret += strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(encoded), "[", "{"), "]", "}"), " ", ", ") + ","
		}
		return nil
	})
	ret += "}\n"

	var hairAssets image.Asset
	ret += "\nvar HairIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", string(image.Hair))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			hairAssets = image.Asset{}
			hairAssets.FileName = info.Name()
			hairAssets.IllustrationPath = path
			hairAssets.SVGContents, err = ioutil.ReadFile(hairAssets.IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", hairAssets.IllustrationPath)
				panic(err.Error())
			}
			hairAssets.HairColored = true
			hairAssets.BodyColored = false
			hairAssets.LightOnly = false
			hairAssets.DarkColored = false
			hairAssets.DarkBWColored = false
			hairAssets.Sex = getSex(info.Name())
			hairAssets.BLK299 = false
			encoded, _ := json.Marshal(hairAssets)
			ret += strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(encoded), "[", "{"), "]", "}"), " ", ", ") + ","
		}
		return nil
	})
	ret += "}\n"

	var hairBackAssets image.Asset
	ret += "\nvar HairBackIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", string(image.HairBack))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			hairBackAssets = image.Asset{}
			hairBackAssets.FileName = info.Name()
			hairBackAssets.IllustrationPath = path
			hairBackAssets.SVGContents, err = ioutil.ReadFile(hairBackAssets.IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", hairBackAssets.IllustrationPath)
				panic(err.Error())
			}
			hairBackAssets.HairColored = true
			hairBackAssets.BodyColored = false
			hairBackAssets.Sex = getSex(info.Name())
			hairBackAssets.LightOnly = false
			hairBackAssets.DarkColored = false
			hairBackAssets.DarkBWColored = false
			hairBackAssets.BLK299 = false
			encoded, _ := json.Marshal(hairBackAssets)
			ret += strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(encoded), "[", "{"), "]", "}"), " ", ", ") + ","
		}
		return nil
	})
	ret += "}\n"

	var hairOutlineAsset image.Asset
	ret += "\nvar HairOutlineIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", string(image.HairOutline))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			hairOutlineAsset = image.Asset{}
			hairOutlineAsset.FileName = info.Name()
			hairOutlineAsset.IllustrationPath = path
			hairOutlineAsset.SVGContents, err = ioutil.ReadFile(path)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", path)
				panic(err.Error())
			}
			hairOutlineAsset.HairColored = false
			hairOutlineAsset.BodyColored = false
			hairOutlineAsset.Sex = getSex(info.Name())
			hairOutlineAsset.LightOnly = false
			hairOutlineAsset.DarkColored = false
			hairOutlineAsset.DarkBWColored = false
			hairOutlineAsset.BLK299 = false
			encoded, _ := json.Marshal(hairOutlineAsset)
			ret += strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(encoded), "[", "{"), "]", "}"), " ", ", ") + ","
		}
		return nil
	})
	ret += "}\n"

	var eyeAssets image.Asset
	ret += "\nvar EyeIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", string(image.Eye))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			eyeAssets = image.Asset{}
			eyeAssets.FileName = info.Name()
			eyeAssets.IllustrationPath = path
			eyeAssets.SVGContents, err = ioutil.ReadFile(eyeAssets.IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", eyeAssets.IllustrationPath)
				panic(err.Error())
			}
			eyeAssets.HairColored = false
			eyeAssets.BodyColored = false
			eyeAssets.Sex = getSex(info.Name())
			eyeAssets.LightOnly = getLightOnly(info.Name())
			eyeAssets.DarkColored = getDarkColored(info.Name())
			eyeAssets.DarkBWColored = getDarkBWColored(info.Name())
			eyeAssets.BLK299 = getBlk299(info.Name())
			encoded, _ := json.Marshal(eyeAssets)
			ret += strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(encoded), "[", "{"), "]", "}"), " ", ", ") + ","
		}
		return nil
	})
	ret += "}\n"

	var mouthAssets image.Asset
	ret += "\nvar MouthIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", string(image.Mouth))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			mouthAssets = image.Asset{}
			mouthAssets.FileName = info.Name()
			mouthAssets.IllustrationPath = path
			mouthAssets.SVGContents, err = ioutil.ReadFile(mouthAssets.IllustrationPath)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", mouthAssets.IllustrationPath)
				panic(err.Error())
			}
			mouthAssets.HairColored = strings.Contains(info.Name(), "_hc")
			mouthAssets.BodyColored = false
			mouthAssets.Sex = getSex(info.Name())
			mouthAssets.LightOnly = getLightOnly(info.Name())
			mouthAssets.DarkColored = getDarkColored(info.Name())
			mouthAssets.DarkBWColored = getDarkBWColored(info.Name())
			mouthAssets.BLK299 = getBlk299(info.Name())
			encoded, _ := json.Marshal(mouthAssets)
			ret += strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(encoded), "[", "{"), "]", "}"), " ", ", ") + ","
		}
		return nil
	})
	ret += "}\n"

	var mouthOutlineAsset image.Asset
	ret += "\nvar MouthOutlineIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", string(image.MouthOutline))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			mouthOutlineAsset = image.Asset{}
			mouthOutlineAsset.FileName = info.Name()
			mouthOutlineAsset.IllustrationPath = path
			mouthOutlineAsset.SVGContents, err = ioutil.ReadFile(path)
			if err != nil {
				glog.Fatalf("Couldn't load file %s", path)
				panic(err.Error())
			}
			mouthOutlineAsset.HairColored = false
			mouthOutlineAsset.BodyColored = false
			mouthOutlineAsset.Sex = getSex(info.Name())
			mouthOutlineAsset.LightOnly = false
			mouthOutlineAsset.DarkColored = false
			mouthOutlineAsset.DarkBWColored = false
			mouthOutlineAsset.BLK299 = false
			encoded, _ := json.Marshal(mouthOutlineAsset)
			ret += strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(encoded), "[", "{"), "]", "}"), " ", ", ") + ","
		}
		return nil
	})
	ret += "}"

	output := path.Join(wd, "image", "illustrations.go")
	outputF, err := os.Create(output)
	defer outputF.Close()
	if err != nil {
		fmt.Printf("Failed to open file for writing %s", output)
	}
	outputF.WriteString(ret)
}
