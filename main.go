package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/samber/lo"
)

const (
	X_SIZE           = 640
	Y_SIZE           = 480
	MIN_STARS        = 50
	MAX_STARS        = 150
	MAX_RANGE        = 100
	DEFAULT_FILENAME = "stars.png"
)

type Star struct {
	x     int
	y     int
	class SpectralClass
}

func create_blank(max_x int, max_y int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, max_x, max_y))

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

func draw_stars(img *image.RGBA, stars []Star) {
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

func is_habitable(s Star) bool {
	return s.class == G || s.class == K || s.class == M
}

func is_in_range(orig Star, dest Star) bool {
	x_diff := math.Pow(float64(orig.x-dest.x), 2)
	y_diff := math.Pow(float64(orig.y-dest.y), 2)
	dist := math.Sqrt(x_diff + y_diff)

	return dist <= MAX_RANGE
}

func draw_routes(img *image.RGBA, stars []Star) {
	habitable := lo.Filter(stars, func(s Star, _ int) bool { return is_habitable(s) })

	gc := draw2dimg.NewGraphicContext(img)

	gc.SetFillColor(color.RGBA{128, 128, 128, 255})
	gc.SetStrokeColor(color.RGBA{128, 128, 128, 255})
	gc.SetLineWidth(0.2)

	for _, s := range habitable {
		in_range := lo.Filter(habitable, func(dest Star, _ int) bool { return is_in_range(s, dest) })

		for _, d := range in_range {
			gc.MoveTo(float64(s.x), float64(s.y))
			gc.LineTo(float64(d.x), float64(d.y))
			gc.Close()
			gc.FillStroke()
		}
	}
}

func write_image(img *image.RGBA, filename string) {
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

	draw_routes(img, stars)
	draw_stars(img, stars)
	write_image(img, *filename)
}
