package image

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/golang/glog"
)

const defaultSize = 1000

type SVG struct {
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	Doc    string `xml:",innerxml"`
}

func CombineSVG(accessories Accessories) ([]byte, error) {
	var (
		body SVG
		hair SVG
	)
	// Parse all SVG assets
	if err := xml.Unmarshal(accessories.BodyAsset.SVGContents, &body); err != nil {
		glog.Fatalf("Unable to parse body SVG %v", err)
		return nil, err
	}
	if err := xml.Unmarshal(accessories.HairAsset.SVGContents, &hair); err != nil {
		glog.Fatalf("Unable to parse hair SVG %v", err)
		return nil, err
	}
	// Create new SVG writer
	var b bytes.Buffer
	canvas := svg.New(&b)
	canvas.Start(defaultSize, defaultSize)
	// Body group
	canvas.Gid("body")
	if accessories.BodyAsset.bodyColored {
		strings.ReplaceAll(body.Doc, "#ffffff", "#ffffff")
	}
	io.WriteString(canvas.Writer, body.Doc)
	canvas.Gend()
	// Hair Group
	canvas.Gid("hair")
	io.WriteString(canvas.Writer, hair.Doc)
	canvas.Gend()
	// End document
	canvas.End()

	return b.Bytes(), nil
}
