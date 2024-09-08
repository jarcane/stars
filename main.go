package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

const (
	X_SIZE           = 1920
	Y_SIZE           = 1080
	MIN_STARS        = 100
	MAX_STARS        = 200
	MAX_RANGE        = 300
	MAX_CONNECTIONS  = 4
	DEFAULT_FILENAME = "stars.png"
)

func create_blank(max_x int, max_y int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, max_x, max_y))

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.Set(x, y, color.Black)
		}
	}

	return img
}

func draw_stars(img *image.RGBA, stars []Star) {
	gc := draw2dimg.NewGraphicContext(img)
	gc.SetStrokeColor(color.RGBA64{128, 128, 128, 255})
	gc.SetLineWidth(1)

	for n := 0; n < len(stars); n++ {
		star := stars[n]
		x := star.x
		y := star.y
		c := spectralClasses[star.class].toRGBA(255)

		gc.SetFillColor(c)

		gc.BeginPath()
		draw2dkit.Circle(gc, float64(x), float64(y), 3+float64(star.class/2))
		gc.FillStroke()
	}
}

func draw_routes(img *image.RGBA, habitable []Star) {
	gc := draw2dimg.NewGraphicContext(img)

	gc.SetFillColor(color.RGBA{128, 128, 128, 255})
	gc.SetStrokeColor(color.RGBA{128, 128, 128, 255})
	gc.SetLineWidth(1)

	for _, s := range habitable {
		filtered := find_neighbors(s, habitable)

		for _, d := range filtered {
			gc.MoveTo(float64(s.x), float64(s.y))
			gc.LineTo(float64(d.x), float64(d.y))
			gc.Close()
			gc.FillStroke()
		}
	}
}

func draw_names(img *image.RGBA, stars []Star) {
	gc := draw2dimg.NewGraphicContext(img)

	draw2d.SetFontFolder("./resource/font")
	draw2d.SetFontNamer(func(fontData draw2d.FontData) string { return "FiraCode-Bold.ttf" })
	gc.SetFontSize(9)
	gc.SetFillColor(color.White)

	for _, star := range stars {
		gc.FillStringAt(star.name, float64(star.x), float64(star.y-7))
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
	habitable := habitables(stars)

	draw_routes(img, habitable)
	draw_stars(img, stars)
	draw_names(img, habitable)
	write_image(img, *filename)
}
