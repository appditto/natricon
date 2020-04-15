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

	fmt.Println("var bodyIllustrations = [...]string{")
	fPath := path.Join(wd, "assets", "illustrations", string(image.Body))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			fmt.Printf("\t\"%s\"\n", info.Name())
		}
		return nil
	})
	fmt.Println("}")

	fmt.Println("\nvar hairIllustrations = [...]string{")
	fPath = path.Join(wd, "assets", "illustrations", string(image.Hair))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			fmt.Printf("\t\"%s\"\n", info.Name())
		}
		return nil
	})
	fmt.Println("}")

	fmt.Println("\nvar eyeIllustrations = [...]string{")
	fPath = path.Join(wd, "assets", "illustrations", string(image.Eye))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			fmt.Printf("\t\"%s\"\n", info.Name())
		}
		return nil
	})
	fmt.Println("}")

	fmt.Println("\nvar mouthIllustrations = [...]string{")
	fPath = path.Join(wd, "assets", "illustrations", string(image.Hair))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			fmt.Printf("\t\"%s\"\n", info.Name())
		}
		return nil
	})
	fmt.Println("}")
}
