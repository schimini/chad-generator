package gen

import (
	"image"
	"image/color"
	"log"
	"strings"

	"github.com/fogleman/gg"
	"github.com/schimini/chad-generator/pkg/util"
)

type Generator struct {
	f string
	i image.Image
}

const quote = "\""

func NewGenerator(imagePath string, fontPath string) *Generator {
	bgImage, err := gg.LoadImage(imagePath)
	bgImage = util.Darker(bgImage, 70) // Darken by 50%
	if err != nil {
		log.Fatal(err)
	}

	return &Generator{
		f: fontPath,
		i: bgImage,
	}
}

func (g *Generator) Generate(t string, o string) {
	// Quotes

	if !strings.HasPrefix(t, quote) {
		t = quote + t
	}
	if !strings.HasSuffix(t, quote) {
		t = t + quote
	}

	// Extract the width and height of the background image
	width := g.i.Bounds().Dx()
	height := g.i.Bounds().Dy()

	// Create a new image context with the same dimensions as the background image
	dc := gg.NewContext(width, height)

	// Draw the background
	dc.DrawImage(g.i, 0, 0)

	if err := dc.LoadFontFace(g.f, float64(48)); err != nil {
		log.Fatal(err)
	}

	dc.SetColor(color.White)
	// Draw the text on the image
	dc.DrawStringWrapped(t, float64(width)/2, float64(height)/2, 0.5, 0.5, float64(width)*3/4, 1.5, gg.AlignCenter)

	// Save the image
	dc.SavePNG(o)
}
