package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

type Star struct {
	x     int
	y     int
	class SpectralClass
}

func main() {
	img := image.NewNRGBA(image.Rect(0, 0, 640, 480))

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.Set(x, y, color.Black)
		}
	}

	num_stars := 50 + rand.Intn(100)
	stars := make([]Star, num_stars)

	for n := 0; n < num_stars; n++ {
		stars[n] = Star{rand.Intn(640), rand.Intn(480), SpectralClass(rand.Intn(7))}
	}

	for n := 0; n < num_stars; n++ {
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

	f, err := os.Create("stars.png")
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
