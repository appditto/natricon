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
		body     SVG
		hair     SVG
		mouth    SVG
		eye      SVG
		backHair SVG
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
	if err := xml.Unmarshal(accessories.MouthAsset.SVGContents, &mouth); err != nil {
		glog.Fatalf("Unable to parse mouth SVG %v", err)
		return nil, err
	}
	if err := xml.Unmarshal(accessories.EyeAsset.SVGContents, &eye); err != nil {
		glog.Fatalf("Unable to parse eye SVG %v", err)
		return nil, err
	}
	if accessories.BackHairAsset != nil {
		if err := xml.Unmarshal(accessories.BackHairAsset.SVGContents, &backHair); err != nil {
			glog.Fatalf("Unable to parse back hair SVG %v", err)
			return nil, err
		}
	}
	// Create new SVG writer
	var b bytes.Buffer
	canvas := svg.New(&b)
	canvas.Start(defaultSize, defaultSize)
	// Add back hair first
	if accessories.BackHairAsset != nil {
		canvas.Gid("backhair")
		if accessories.HairAsset.HairColored {
			backHair.Doc = strings.ReplaceAll(backHair.Doc, "#FF0000", accessories.HairColor.ToHTML(true))
		}
		io.WriteString(canvas.Writer, backHair.Doc)
		canvas.Gend()
	}
	// Body group
	canvas.Gid("body")
	if accessories.BodyAsset.BodyColored {
		body.Doc = strings.ReplaceAll(body.Doc, "#00FFFF", accessories.BodyColor.ToHTML(true))
	}
	io.WriteString(canvas.Writer, body.Doc)
	canvas.Gend()
	// Hair Group
	canvas.Gid("hair")
	if accessories.HairAsset.HairColored {
		hair.Doc = strings.ReplaceAll(hair.Doc, "#FF0000", accessories.HairColor.ToHTML(true))
	}
	io.WriteString(canvas.Writer, hair.Doc)
	canvas.Gend()
	// Mouth group
	canvas.Gid("mouth")
	if accessories.HairAsset.HairColored {
		mouth.Doc = strings.ReplaceAll(mouth.Doc, "#FFFF00", accessories.HairColor.ToHTML(true))
	}
	io.WriteString(canvas.Writer, mouth.Doc)
	canvas.Gend()
	// Eye group
	canvas.Gid("eye")
	io.WriteString(canvas.Writer, eye.Doc)
	canvas.Gend()
	// End document
	canvas.End()

	return b.Bytes(), nil
}
