package julia_set

import (
	"bytes"
	"context"
	"fmt"
	"github.com/muesli/gamut"
	"harrisonwaffel/fractals/pkg/ffmpeg"
	"harrisonwaffel/fractals/pkg/util"
	"image"
	"image/color"
	"image/png"
	"math"
	"sync"
)

// we generate a video of the julia set by
// incrementing the InitialConstantReal
// by StepSize until we reach InitialConstantReal + TotalRange

type JuliaSet struct {
	Ctx                 context.Context
	InitialConstantReal float32
	ConstantImaginary   float32
	TotalRange          float32
	StepSize            float32

	// final video dimensions

	VideoHeight int
	VideoWidth  int
}

func (js *JuliaSet) GenerateSet(moveX, moveY, zoom float32) chan ffmpeg.FrameChunk {
	if zoom == 0. {
		zoom = 1.0
	}

	gen := &JuliaSetGenerator{
		Ctx:               js.Ctx,
		ConstantReal:      js.InitialConstantReal,
		ConstantImaginary: js.ConstantImaginary,
		Zoom:              zoom,
		MoveX:             moveX,
		MoveY:             moveY,
		// we use 255 since RGBA has at most 255 values,
		// this allows us to use the full color gradient
		ConvergeThreshold: 175,
	}
	gen.InitPalette(util.DefaultPalette...)

	frameChan := make(chan ffmpeg.FrameChunk)
	steps := int(math.Ceil(float64(js.TotalRange / js.StepSize)))
	fmt.Println("intiial steps", steps)
	adjustedSteps := int(math.Ceil(float64(js.TotalRange/js.StepSize) / 24))
	fmt.Println("adjusted steps", adjustedSteps)

	go func(gen *JuliaSetGenerator, frameChan chan ffmpeg.FrameChunk) {
		for i := 0; i < int(math.Ceil(float64(js.TotalRange/js.StepSize)/24)); i++ {

			select {
			case <-gen.Ctx.Done():
				fmt.Println("ctx dead")
				close(frameChan)
				return
			default:
			}

			wg := sync.WaitGroup{}
			chunk := ffmpeg.FrameChunk{
				Frames: make([]ffmpeg.Frame, ffmpeg.FPS+1),
			}

			// Concurrently process as many frames per second as we want
			for j := 1; j <= ffmpeg.FPS; j++ {
				wg.Add(1)
				go func(wg *sync.WaitGroup, chunk *ffmpeg.FrameChunk, gen JuliaSetGenerator, index, frameIndex int) {
					frameBuff := new(bytes.Buffer)
					png.Encode(frameBuff, gen.CreateFrame(js.VideoWidth, js.VideoHeight))
					chunk.Frames[index] = ffmpeg.Frame{
						Index: index,
						Frame: frameBuff.Bytes(),
					}
					wg.Done()
				}(&wg, &chunk, *gen, j, j*i)

				// increment the set (go forward in time)
				gen.ConstantReal += js.StepSize
			}
			wg.Wait()

			frameChan <- chunk
		}

		close(frameChan)
	}(gen, frameChan)

	return frameChan
}

type JuliaSetGenerator struct {
	Ctx               context.Context
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
		cp = append(cp, gamut.Blends(gamut.Hex(colors[i]), gamut.Hex(colors[i+1]), int(util.MapToRangeEnd)/(len(colors)-1)+1)...)
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
