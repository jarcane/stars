package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

const (
	X_SIZE           = 640
	Y_SIZE           = 480
	MIN_STARS        = 50
	MAX_STARS        = 150
	DEFAULT_FILENAME = "stars.png"
)

type Star struct {
	x     int
	y     int
	class SpectralClass
}

func create_blank(max_x int, max_y int) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, max_x, max_y))

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.Set(x, y, color.Black)
		}
	}

	return img
}

func generate_stars(max_x int, max_y int) []Star {
	num_stars := MIN_STARS + rand.Intn(MAX_STARS-MIN_STARS)
	stars := make([]Star, num_stars)

	for n := 0; n < num_stars; n++ {
		stars[n] = Star{rand.Intn(max_x), rand.Intn(max_y), SpectralClass(rand.Intn(7))}
	}

	return stars
}

func draw_stars(img *image.NRGBA, stars []Star) {
	for n := 0; n < len(stars); n++ {
		star := stars[n]
		x := star.x
		y := star.y
		c := spectralClasses[star.class].toRGBA(255)

		img.Set(x, y, c)
		img.Set(x+1, y, c)
		img.Set(x, y+1, c)
		img.Set(x-1, y, c)
		img.Set(x, y-1, c)
	}
}

func write_image(img *image.NRGBA, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	filename := flag.String("filename", DEFAULT_FILENAME, "the filename of the output")
	x := flag.Int("x", X_SIZE, "Horizontal size of image")
	y := flag.Int("y", Y_SIZE, "Vertical size of image")
	flag.Parse()

	img := create_blank(*x, *y)

	stars := generate_stars(*x, *y)

	draw_stars(img, stars)

	write_image(img, *filename)
}
