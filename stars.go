package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/samber/lo"
)

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
	name  string
	class SpectralClass
}

func is_habitable(s Star) bool {
	return s.class == G || s.class == K || s.class == M
}

func habitables(stars []Star) []Star {
	return lo.Filter(stars, func(s Star, _ int) bool { return is_habitable(s) })
}

func is_in_range(orig Star, dest Star) bool {
	x_diff := math.Pow(float64(orig.x-dest.x), 2)
	y_diff := math.Pow(float64(orig.y-dest.y), 2)
	dist := math.Sqrt(x_diff + y_diff)

	return dist <= MAX_RANGE
}

func find_neighbors(origin Star, stars []Star) []Star {
	in_range := lo.Filter(stars, func(dest Star, _ int) bool { return is_in_range(origin, dest) })
	filtered := lo.Slice(lo.Uniq(in_range), 0, MAX_CONNECTIONS)

	return filtered
}

func generate_stars(max_x int, max_y int) []Star {
	num_stars := MIN_STARS + rand.Intn(MAX_STARS-MIN_STARS)
	stars := make([]Star, num_stars)

	for n := 0; n < num_stars; n++ {
		stars[n] = Star{rand.Intn(max_x), rand.Intn(max_y), randomStarName(n), SpectralClass(rand.Intn(7))}
	}

	return stars
}
