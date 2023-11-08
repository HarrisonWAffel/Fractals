package mandelbrot_set

import (
	"bytes"
	"context"
	"harrisonwaffel/fractals/pkg/ffmpeg"
	"harrisonwaffel/fractals/pkg/util"
	"image"
	"image/color"
	"image/png"
	"sync"
)

type MandelbrotGenerator struct {
	Ctx context.Context

	ImageHeight      int
	ImageWidth       int
	ConvergenceLimit int

	MoveX        float64
	MoveY        float64
	Zoom         float64
	ZoomStepSize float64

	Palette []color.Color
}

// Unlike a julia set, we do not iterate with the mandelbrot set,
// we just translate and zoom around the image. A mandelbrot set will zoom infinitely
// and that's where the fun is

func (mg *MandelbrotGenerator) GenerateZoomVideo() chan ffmpeg.FrameChunk {
	frameChan := make(chan ffmpeg.FrameChunk)
	//30 second video
	go func(gen *MandelbrotGenerator, frameChan chan ffmpeg.FrameChunk) {
		videoLength := 60
		for i := 0; i < (videoLength * ffmpeg.FPS / 10); i++ {
			select {
			case <-mg.Ctx.Done():
				close(frameChan)
				return
			default:
			}
			wg := sync.WaitGroup{}
			frameChunk := ffmpeg.FrameChunk{}
			frameChunk.Frames = make([]ffmpeg.Frame, videoLength*ffmpeg.FPS)
			// ten routines per second
			for j := 0; j < 10; j++ {
				wg.Add(1)

				// There's a problem with this. As we continue to zoom
				// into the set, we need a larger and larger zoom step
				// in order for the video to seem smooth (as the set gets
				// smaller and smaller very quickly (but not exponentially)).
				mg.Zoom += mg.ZoomStepSize + (mg.Zoom / 128)

				go func(wg *sync.WaitGroup, chunk *ffmpeg.FrameChunk, gen *MandelbrotGenerator, index int, zoom float64) {
					frameBuff := new(bytes.Buffer)
					png.Encode(frameBuff, mg.GenerateImage(zoom))
					frameChunk.Frames[index] = ffmpeg.Frame{
						Frame: frameBuff.Bytes(),
					}
					wg.Done()
				}(&wg, &frameChunk, mg, j, mg.Zoom)
			}
			wg.Wait()
			frameChan <- frameChunk
		}
		close(frameChan)
	}(mg, frameChan)

	return frameChan
}

func (mg *MandelbrotGenerator) GenerateImage(zoom float64) *image.RGBA {
	if zoom == 0 {
		zoom = 1.
	}

	img := image.NewRGBA(image.Rect(0, 0, mg.ImageWidth, mg.ImageHeight))

	quarterWidth := 4.0 / (float64(mg.ImageWidth) * zoom)
	quarterHeight := 4.0 / (float64(mg.ImageHeight) * zoom)

	halfWidth := (float64(mg.ImageWidth) - (mg.MoveX * zoom)) / 2.0
	halfHeight := (float64(mg.ImageHeight) - (mg.MoveY * zoom)) / 2.0

	for y := 0; y < img.Bounds().Size().Y; y++ {
		for x := 0; x < img.Bounds().Size().X; x++ {
			constantReal := (float64(x) - halfWidth) * quarterWidth
			constantImaginary := (float64(y) - halfHeight) * quarterHeight

			iterations := 0
			var a, b float64
			for ; iterations < mg.ConvergenceLimit; iterations++ {
				newReal := (a * a) - (b * b) + constantReal
				b = (2 * a * b) + constantImaginary
				a = newReal
				m := a*a + b*b
				if m > 4 {
					break
				}
			}

			img.Set(x, y, mg.Palette[util.MapToRange64(float64(iterations))])
		}
	}

	return img
}
