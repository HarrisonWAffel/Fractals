package pkg

import (
	"bytes"
	goffmpeg "github.com/u2takey/ffmpeg-go"
	"harrisonwaffel/fractals/pkg/ffmpeg"
	julia_set "harrisonwaffel/fractals/pkg/julia-set"

	"io"
	"sync"
	"testing"
)

func TestIt(t *testing.T) {

	js := julia_set.JuliaSet{
		InitialConstantReal: -0.8,
		ConstantImaginary:   0.156,
		TotalRange:          0.0012,
		StepSize:            0.000001,
		VideoHeight:         1600,
		VideoWidth:          1600,
	}

	b := new(bytes.Buffer)
	frameChan := js.GenerateSet(0.0, 0.0, 1.0)
	ffmpegProcessor := ffmpeg.Processor{}
	frameOutput := ffmpegProcessor.GetChunkReader(frameChan)

	wg := sync.WaitGroup{}
	wg.Add(1)
	errChan := make(chan error)
	go func(frameOutput *io.PipeReader, writer *bytes.Buffer, errChan chan error) {
		for {

			n, err := io.Copy(b, frameOutput)
			if n == 0 || err != nil {
				break
			}

		}
		wg.Done()
	}(frameOutput, b, errChan)
	wg.Wait()

	if len(b.Bytes()) == 0 {
		panic("no bytes")
	}

	err := goffmpeg.Input("pipe:").
		WithInput(b).
		Output("test.mp4", goffmpeg.KwArgs{
			"c:v":     "libx264",
			"pix_fmt": "yuv420p",
			"f":       "ismv",
		}).
		ErrorToStdOut().
		Run()
	if err != nil {
		panic(err)
	}
}
