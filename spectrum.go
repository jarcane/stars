package main

import "image/color"

type RGB struct {
	r uint8
	g uint8
	b uint8
}

func (c RGB) toRGBA(a uint8) color.RGBA {
	return color.RGBA{c.r, c.g, c.b, a}
}

type SpectralClass int

const (
	O = iota
	B
	A
	F
	G
	K
	M
)

// Spectral RGB values sourced from Mitchell Charity: http://www.vendian.org/mncharity/dir3/starcolor/
var spectralClasses = map[SpectralClass]RGB{
	O: {155, 176, 255},
	B: {170, 191, 255},
	A: {202, 215, 255},
	F: {248, 247, 255},
	G: {255, 244, 234},
	K: {255, 210, 161},
	M: {255, 204, 111},
}

type Star struct {
	x     int
	y     int
	class SpectralClass
}
