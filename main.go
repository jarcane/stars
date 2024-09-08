package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	stars := image.NewNRGBA(image.Rect(0, 0, 640, 480))

	for y := stars.Bounds().Min.Y; y < stars.Bounds().Max.Y; y++ {
		for x := stars.Bounds().Min.X; x < stars.Bounds().Max.X; x++ {
			stars.Set(x, y, color.Black)
		}
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
