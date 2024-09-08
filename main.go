package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

func main() {
	stars := image.NewNRGBA(image.Rect(0, 0, 640, 480))

	for y := stars.Bounds().Min.Y; y < stars.Bounds().Max.Y; y++ {
		for x := stars.Bounds().Min.X; x < stars.Bounds().Max.X; x++ {
			stars.Set(x, y, color.Black)
		}
	}

	num_stars := rand.Intn(150)

	for n := 0; n < num_stars; n++ {
		x := rand.Intn(640)
		y := rand.Intn(480)
		c := color.RGBA{uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255}

		stars.Set(x, y, c)
		stars.Set(x+1, y, c)
		stars.Set(x, y+1, c)
		stars.Set(x-1, y, c)
		stars.Set(x, y-1, c)
	}

	f, err := os.Create("stars.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, stars); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
