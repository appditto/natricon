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
			bodyAsset.Sex = getSex(info.Name())
			encoded, _ := json.Marshal(bodyAsset)
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
			hairAssets.Sex = getSex(info.Name())
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
			encoded, _ := json.Marshal(hairBackAssets)
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
			encoded, _ := json.Marshal(mouthAssets)
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
