package pkg

import (
	"bytes"
	"fmt"
	"harrisonwaffel/fractals/pkg/ffmpeg"
	julia_set "harrisonwaffel/fractals/pkg/julia-set"
	"io"
	"os"
	"testing"
	"time"
)

func TestIt(t *testing.T) {

	js := julia_set.JuliaSet{
		InitialConstantReal: -0.8,
		ConstantImaginary:   0.156,
		TotalRange:          0.0012,
		StepSize:            0.00001,
		VideoHeight:         1600,
		VideoWidth:          1600,
	}

	// todo; this part needs to become stream'd tomorrow
	frames := js.GenerateSet(0, 0, 2.5)

	p := ffmpeg.Processor{}
	pr1, pw1 := io.Pipe()
	b := new(bytes.Buffer)
	go func() {
		for {
			n, err := io.Copy(b, pr1)
			if err != nil {
				panic(err)
			}
			if n == 0 {
				return
			}

		}
	}()

	time.Sleep(250 * time.Millisecond)
	if err := p.CreateVideo(frames, pw1); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	os.WriteFile("test.mp4", b.Bytes(), os.ModePerm)
}
