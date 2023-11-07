package julia_set

import (
	"fmt"
	"github.com/muesli/gamut"
	"harrisonwaffel/fractals/pkg/util"
	"image"
	"image/color"
	"math"
	"sync"
)

// we generate a video of the julia set by
// incrementing the InitialConstantReal
// by StepSize until we reach InitialConstantReal + TotalRange

type JuliaSet struct {
	InitialConstantReal float32
	ConstantImaginary   float32
	TotalRange          float32
	StepSize            float32

	// final video dimensions

	VideoHeight int
	VideoWidth  int
}

func (js *JuliaSet) GenerateSet(moveX, moveY, zoom float32) []*image.RGBA {

	if zoom == 0. {
		zoom = 1.0
	}

	gen := &JuliaSetGenerator{
		ConstantReal:      js.InitialConstantReal,
		ConstantImaginary: js.ConstantImaginary,
		Zoom:              zoom,
		MoveX:             moveX,
		MoveY:             moveY,
		// we use 255 since RGBA has at most 255 values,
		// this allows us to use the full color gradient
		ConvergeThreshold: 255,
	}

	gen.InitPalette(util.DefaultPalette...)

	frames := make([]*image.RGBA, int(math.Ceil(float64(js.TotalRange/js.StepSize))))
	// By incrementing the constant real we effectively iterate through time.
	// Controlling the step size controls the speed at which we move through time.
	// The total amount of time is determined by the js.TotalRange,
	fmt.Println(fmt.Sprintf("%v", math.Ceil(float64(js.TotalRange/js.StepSize))))
	wg := sync.WaitGroup{}

	// todo; handle the speed problem. if we want to stream it we likely need to chunk it
	for i := 0; i < int(math.Ceil(float64(js.TotalRange/js.StepSize))); i++ {
		go func(i int, frames []*image.RGBA, gen JuliaSetGenerator) {
			wg.Add(1)
			frames[i] = gen.CreateFrame(js.VideoWidth, js.VideoHeight)
			wg.Done()
		}(i, frames, *gen)
		gen.ConstantReal += js.StepSize
		// increment the set (go forward in time)
		fmt.Println(gen.ConstantReal)
	}
	wg.Wait()
	return frames
}

type JuliaSetGenerator struct {
	ConstantReal      float32
	ConstantImaginary float32
	// zoom into the center of the frame
	Zoom float32
	// pan the image on the X axis
	MoveX float32
	// pan the image on the Y axis
	MoveY float32
	// how many steps we should check before
	// we decide if the index converges or not
	ConvergeThreshold int
	// the color Palette we will use when generating frames
	Palette []color.Color
}

func (j *JuliaSetGenerator) InitPalette(colors ...string) {
	var cp []color.Color
	for i := 0; i < len(colors)-1; i++ {
		cp = append(cp, gamut.Blends(gamut.Hex(colors[i]), gamut.Hex(colors[i+1]), 128/(len(colors)-1)+1)...)
	}
	j.Palette = cp
}

func (j *JuliaSetGenerator) CreateFrame(fwidth, fheight int) *image.RGBA {
	var newReal, newImaginary, oldReal, oldImaginary float32
	maxIterations := 255 // 255 for RGB
	myImg := image.NewRGBA(image.Rect(0, 0, fwidth, fheight))
	height := float32(myImg.Bounds().Size().Y)
	width := float32(myImg.Bounds().Size().X)

	newRealpt2 := (0.5 * j.Zoom * width) + j.MoveX
	newImgpt2 := (0.5 * j.Zoom * height) + j.MoveY
	halfHeight := height / 2
	halfWidth := width / 2
	for i := 0; i < myImg.Bounds().Size().X; i++ {
		for k := 0; k < myImg.Bounds().Size().Y; k++ {
			newReal = 1.5 * (float32(k) - halfWidth) / newRealpt2
			newImaginary = (float32(i) - halfHeight) / newImgpt2

			iterations := 0
			for ; iterations < maxIterations; iterations++ {
				oldReal = newReal
				oldImaginary = newImaginary
				// z is a complex number described by newReal and newImaginary
				// a+bi where a = newReal, b = newImaginary

				// note: math.Pow is much slower than the below statements
				newReal = oldReal*oldReal - oldImaginary*oldImaginary + j.ConstantReal
				newImaginary = 2*oldReal*oldImaginary + j.ConstantImaginary
				if (newReal*newReal + newImaginary*newImaginary) > 2 {
					break
				}
			}
			myImg.Set(i, k, j.Palette[util.MapToRange(float32(iterations))])
		}
	}

	return myImg
}
