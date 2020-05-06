package main

import (
	"syscall/js"
)

var seed string = "123456789"

func getNatriconStr(inputs []js.Value) {
	callback := inputs[len(inputs)-1:][0]
	message := inputs[0].String()
	if !nano.ValidateAddress(message) {
		callback.Invoke("Invalid address", js.Null())
		return
	}
	sha256 := nano.AddressSha256(message, seed)

	accessories, err := natricon.image.GetAccessoriesForHash(sha256)
	if err != nil {
		callback.Invoke("failure", js.Null())
		return
	}
	bodyHsv := accessories.BodyColor.ToHSV()
	hairHsv := accessories.HairColor.ToHSV()
	deltaHsv := color.HSV{}
	deltaHsv.H = hairHsv.H - bodyHsv.H
	deltaHsv.S = hairHsv.S - bodyHsv.S
	deltaHsv.V = hairHsv.V - bodyHsv.V
	svg, err := natricon.image.CombineSVG(accessories)
	if err != nil {
		callback.Invoke("failure", js.Null())
		return
	}

	callback.Invoke(js.Null(), string(svg))
}

func main() {
	js.Global().Set("getNatriconStr", js.FuncOf(getNatriconStr))
}
