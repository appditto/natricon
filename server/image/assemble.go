package image

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"strings"
	"sync"

	svg "github.com/ajstarks/svgo"
	"github.com/appditto/natricon/server/color"
	"github.com/golang/glog"
	minify "github.com/tdewolff/minify/v2"
	minifysvg "github.com/tdewolff/minify/v2/svg"
)

const DefaultSize = 512   // Default SVG width/height attribute
const opacityLower = 0.15 // Minimum lower opacity threshold
const opacityUpper = 0.6  // Maximum upper opacity threshold

type SVG struct {
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	Doc    string `xml:",innerxml"`
}

func CombineSVG(accessories Accessories) ([]byte, error) {
	var (
		body         SVG
		hair         SVG
		mouth        SVG
		eye          SVG
		backHair     SVG
		bodyOutline  SVG
		hairOutline  SVG
		mouthOutline SVG
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
	if accessories.BodyOutlineAsset != nil {
		if err := xml.Unmarshal(accessories.BodyOutlineAsset.SVGContents, &bodyOutline); err != nil {
			glog.Fatalf("Unable to parse body outline SVG %v", err)
			return nil, err
		}
	}
	if accessories.HairOutlineAsset != nil {
		if err := xml.Unmarshal(accessories.HairOutlineAsset.SVGContents, &hairOutline); err != nil {
			glog.Fatalf("Unable to parse hair outline SVG %v", err)
			return nil, err
		}
	}
	if accessories.MouthOutlineAsset != nil {
		if err := xml.Unmarshal(accessories.MouthOutlineAsset.SVGContents, &mouthOutline); err != nil {
			glog.Fatalf("Unable to parse mouth outline SVG %v", err)
			return nil, err
		}
	}
	// Create new SVG writer
	var b bytes.Buffer
	canvas := svg.New(&b)
	canvas.Startraw(fmt.Sprintf("viewBox=\"0 0 %d %d\"", DefaultSize, DefaultSize))
	// Add body outline
	if accessories.BodyOutlineAsset != nil {
		canvas.Gid("bodyOutline")
		bodyOutline.Doc = strings.ReplaceAll(bodyOutline.Doc, "black", accessories.OutlineColor.ToHTML(true))
		io.WriteString(canvas.Writer, bodyOutline.Doc)
		canvas.Gend()
	}
	// Add mouth outline
	if accessories.MouthOutlineAsset != nil {
		canvas.Gid("mouthOutline")
		mouthOutline.Doc = strings.ReplaceAll(mouthOutline.Doc, "black", accessories.OutlineColor.ToHTML(true))
		io.WriteString(canvas.Writer, mouthOutline.Doc)
		canvas.Gend()
	}
	// Add hair outline
	if accessories.HairOutlineAsset != nil {
		canvas.Gid("hairOutline")
		hairOutline.Doc = strings.ReplaceAll(hairOutline.Doc, "black", accessories.OutlineColor.ToHTML(true))
		io.WriteString(canvas.Writer, hairOutline.Doc)
		canvas.Gend()
	}
	// Add back hair
	if accessories.BackHairAsset != nil {
		canvas.Gid("backhair")
		if accessories.HairAsset.HairColored {
			backHair.Doc = strings.ReplaceAll(backHair.Doc, "#FF0000", accessories.HairColor.ToHTML(true))
			backHair.Doc = strings.ReplaceAll(backHair.Doc, "fill-opacity=\"0.15\"", fmt.Sprintf("fill-opacity=\"%f\"", GetTargetOpacity(accessories.HairColor.ToHSV())))
		}
		io.WriteString(canvas.Writer, backHair.Doc)
		canvas.Gend()
	}
	// Body group
	canvas.Gid("body")
	if accessories.BodyAsset.BodyColored {
		body.Doc = strings.ReplaceAll(body.Doc, "#00FFFF", accessories.BodyColor.ToHTML(true))
		body.Doc = strings.ReplaceAll(body.Doc, "fill-opacity=\"0.15\"", fmt.Sprintf("fill-opacity=\"%f\"", GetTargetOpacity(accessories.BodyColor.ToHSV())))
	}
	io.WriteString(canvas.Writer, body.Doc)
	canvas.Gend()
	// Hair Group
	canvas.Gid("hair")
	if accessories.HairAsset.HairColored {
		hair.Doc = strings.ReplaceAll(hair.Doc, "#FF0000", accessories.HairColor.ToHTML(true))
		hair.Doc = strings.ReplaceAll(hair.Doc, "fill-opacity=\"0.15\"", fmt.Sprintf("fill-opacity=\"%f\"", GetTargetOpacity(accessories.HairColor.ToHSV())))
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

	// Minify
	var ret []byte
	ret, _ = getMinifier().minifier.Bytes("image/svg+xml", b.Bytes())

	return ret, nil
}

func GetTargetOpacity(color color.HSV) float64 {
	brightness := color.V
	ret := (1.0-brightness)*(opacityUpper-opacityLower)/(1.0-MinBrightness) + opacityLower
	// Return result rounded to 2 places
	return math.Round(ret*100) / 100
}

// Singleton to get minifier
type minifySingleton struct {
	minifier *minify.M
}

var mSingleton *minifySingleton
var onceM sync.Once

func getMinifier() *minifySingleton {
	onceM.Do(func() {
		minifier := minify.New()
		minifier.AddFunc("image/svg+xml", minifysvg.Minify)
		mSingleton = &minifySingleton{
			minifier: minifier,
		}
	})
	return mSingleton
}
