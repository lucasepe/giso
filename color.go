package giso

import (
	"fmt"
	"image/color"
	"math"
	"strings"
)

func ParseHexColor(hexColor string) color.Color {
	hexColor = strings.TrimPrefix(hexColor, "#")
	a := uint8(255)

	var r, g, b uint8
	if len(hexColor) == 3 {
		format := "%1x%1x%1x"
		fmt.Sscanf(hexColor, format, &r, &g, &b)
		r |= r << 4
		g |= g << 4
		b |= b << 4
	}

	if len(hexColor) == 6 {
		format := "%02x%02x%02x"
		fmt.Sscanf(hexColor, format, &r, &g, &b)
	}

	if len(hexColor) == 8 {
		format := "%02x%02x%02x%02x"
		fmt.Sscanf(hexColor, format, &r, &g, &b, &a)
	}

	return color.NRGBA{R: r, G: g, B: b, A: a}
}

// Lighten increases the brightness of the color based on HSL mode. (`0.5` as `50%`)
func Lighten(c color.Color, percent float64) color.Color {
	percent = percent * 100

	r, g, b, a := c.RGBA()

	h, s, l := RGBToHSL(float64(r), float64(g), float64(b))

	l += percent
	if l > 100 {
		l = 100
	}

	r1, g1, b1 := HSLToRGB(h, s, l)

	return color.NRGBA{
		R: uint8(r1),
		G: uint8(g1),
		B: uint8(b1),
		A: uint8(a),
	}
}

// RGBToHSL converts RGB values to HSL (in [0.1] range).
// http://axonflux.com/handy-rgb-to-hsl-and-rgb-to-hsv-color-model-c
func RGBToHSL(r, g, b float64) (h, s, l float64) {
	r, g, b = r/255, g/255, b/255

	min, max := minMax(r, g, b)
	h = (max + min) / 2
	s = h
	l = h

	if max == min {
		// achromatic
		return 0, 0, l
	}

	d := (max - min)
	if l > 0.5 {
		s = d / (2 - max - min)
	} else {
		s = d / (max + min)
	}

	switch max {
	case r:
		h = (g - b) / d
		if g < b {
			h = h + 6
		}

	case g:
		h = (b-r)/d + 2
	case b:
		h = (r-g)/d + 4
	}

	h = h / 6

	return h, s, l
}

// HSLToRGB converts HSL values to RGB
// http://axonflux.com/handy-rgb-to-hsl-and-rgb-to-hsv-color-model-c
func HSLToRGB(h, s, l float64) (r, g, b float64) {
	if s == 0 {
		// it's gray
		return l, l, l
	}

	var p, q float64
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = (l + s) - (l * s)
	}
	p = 2*l - q

	r = hueToRGB(p, q, h+(1.0/3.0))
	g = hueToRGB(p, q, h)
	b = hueToRGB(p, q, h-(1.0/3.0))

	return math.Round(r * 255), math.Round(g * 255), math.Round(b * 255)
}

// hueToRGB helper func used to convert hue to rgb.
// http://axonflux.com/handy-rgb-to-hsl-and-rgb-to-hsv-color-model-c
func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t++
	}
	if t > 1 {
		t--
	}

	if t < 1/6 {
		return p + (q-p)*6*t
	}

	if t < 1/2 {
		return q
	}

	if t < 2/3 {
		return p + (q-p)*(2/3-t)*6
	}

	return p
}

// minMax helper function to find max e min of numbers.
func minMax(x ...float64) (float64, float64) {
	max := x[0]
	min := x[0]
	for _, value := range x {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
