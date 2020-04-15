package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/appditto/natricon/image"
)

func LoadAssetsToArray() {
	wd, err := os.Getwd()
	if err != nil {
		panic("Can't get working directory")
	}

	ret := "package image\n\n"

	ret += "var BodyIllustrations = [...]string{\n"
	fPath := path.Join(wd, "assets", "illustrations", string(image.Body))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			ret += fmt.Sprintf("\t\"%s\",\n", info.Name())
		}
		return nil
	})
	ret += "}\n"

	ret += "\nvar HairIllustrations = [...]string{\n"
	fPath = path.Join(wd, "assets", "illustrations", string(image.Hair))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			ret += fmt.Sprintf("\t\"%s\",\n", info.Name())
		}
		return nil
	})
	ret += "}\n"

	ret += "\nvar HairBackIllustrations = [...]string{\n"
	fPath = path.Join(wd, "assets", "illustrations", string(image.HairBack))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			ret += fmt.Sprintf("\t\"%s\",\n", info.Name())
		}
		return nil
	})
	ret += "}\n"

	ret += "\nvar EyeIllustrations = [...]string{\n"
	fPath = path.Join(wd, "assets", "illustrations", string(image.Eye))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			ret += fmt.Sprintf("\t\"%s\",\n", info.Name())
		}
		return nil
	})
	ret += "}\n"

	ret += "\nvar MouthIllustrations = [...]string{\n"
	fPath = path.Join(wd, "assets", "illustrations", string(image.Mouth))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			ret += fmt.Sprintf("\t\"%s\",\n", info.Name())
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
