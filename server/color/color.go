package color

import (
	"errors"
	"fmt"
	"math"
)

type RGB struct {
	R, G, B float64
}

// Takes a string like '#123456' or 'ABCDEF' and returns an RGB between 0..255
func HTMLToRGB(in string) (RGB, error) {
	if in[0] == '#' {
		in = in[1:]
	}

	if len(in) != 6 {
		return RGB{}, errors.New("Invalid string length")
	}

	var r, g, b byte
	if n, err := fmt.Sscanf(in, "%2x%2x%2x", &r, &g, &b); err != nil || n != 3 {
		return RGB{}, err
	}

	return RGB{float64(r), float64(g), float64(b)}, nil
}

// A nudge to make truncation round to nearest number instead of flooring
const delta = 1 / 512.0

func (c RGB) ToHTML(withHash bool) string {
	if withHash {
		return fmt.Sprintf("#%02x%02x%02x", byte((c.R + delta)), byte((c.G + delta)), byte((c.B + delta)))
	}
	return fmt.Sprintf("%02x%02x%02x", byte((c.R + delta)), byte((c.G + delta)), byte((c.B + delta)))
}

type HSV struct {
	H, S, V float64
}

// ToHSV - Returns HSV as (0..360, 0..1, 0..1)
func (c RGB) ToHSV() HSV {
	var h, s, v float64

	r := c.R
	g := c.G
	b := c.B

	min := math.Min(math.Min(r, g), b)
	max := math.Max(math.Max(r, g), b)
	delta := max - min

	if max != 0.0 {
		s = delta / max
		v = max / 255
	}

	// hue
	if delta != 0 {
		if r == max {
			h = (g - b) / delta
		} else if g == max {
			h = 2.0 + ((b - r) / delta)
		} else {
			h = 4 + ((r - g) / delta)
		}
	}
	h = h * 60
	if h < 0 {
		h = h + 360
	}

	return HSV{h, s, v}
}

// ToRGB - convert HSV to RGB
func (c HSV) ToRGB() RGB {
	Hp := c.H / 60.0
	C := c.V * c.S
	X := C * (1.0 - math.Abs(math.Mod(Hp, 2.0)-1.0))

	m := c.V - C
	r, g, b := 0.0, 0.0, 0.0

	switch {
	case 0.0 <= Hp && Hp < 1.0:
		r = C
		g = X
	case 1.0 <= Hp && Hp < 2.0:
		r = X
		g = C
	case 2.0 <= Hp && Hp < 3.0:
		g = C
		b = X
	case 3.0 <= Hp && Hp < 4.0:
		g = X
		b = C
	case 4.0 <= Hp && Hp < 5.0:
		r = X
		b = C
	case 5.0 <= Hp && Hp <= 6.0:
		r = C
		b = X
	}

	return RGB{255 * (m + r), 255 * (m + g), 255 * (m + b)}
}

func (c HSV) ToHTML(withHash bool) string {
	return c.ToRGB().ToHTML(withHash)
}
