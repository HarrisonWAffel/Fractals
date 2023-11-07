package mandelbrot_set

import (
	"bytes"
	"context"
	"harrisonwaffel/fractals/pkg/util"
	"image"
	"image/color"
	"image/png"
)

type MandelbrotGenerator struct {
	Ctx context.Context

	ImageHeight      int
	ImageWidth       int
	ConvergenceLimit int

	MoveX float32
	MoveY float32
	Zoom  float32

	Palette []color.Color
}

// Unlike a julia set, we do not iterate with the mandelbrot set,
// we just translate and zoom around the image. A mandelbrot set will zoom infinitely
// and that's where the fun is

func (mg *MandelbrotGenerator) GenerateImage() ([]byte, error) {

	img := image.NewRGBA(image.Rect(0, 0, mg.ImageWidth, mg.ImageHeight))
	mg.Palette = util.InitPalette()

	quarterWidth := 4.0 / (float32(mg.ImageWidth) * mg.Zoom)
	quarterHeight := 4.0 / (float32(mg.ImageHeight) * mg.Zoom)

	halfWidth := (float32(mg.ImageWidth) - float32(mg.MoveX)) / 2.0
	halfHeight := (float32(mg.ImageHeight) - float32(mg.MoveY)) / 2.0

	for y := 0; y < img.Bounds().Size().Y; y++ {
		for x := 0; x < img.Bounds().Size().X; x++ {

			constantReal := (float32(x) - halfWidth) * quarterWidth
			constantImaginary := (float32(y) - halfHeight) * quarterHeight

			iterations := 0
			var a, b float32
			for ; iterations < mg.ConvergenceLimit; iterations++ {
				newReal := (a * a) - (b * b) + constantReal
				b = (2 * a * b) + constantImaginary
				a = newReal
				m := a*a + b*b
				if m > 4 {
					break
				}
			}

			img.Set(x, y, mg.Palette[util.MapToRange(float32(iterations))])
		}
	}

	imgBuff := new(bytes.Buffer)
	err := png.Encode(imgBuff, img)
	if err != nil {
		return nil, err
	}

	return imgBuff.Bytes(), nil
}
