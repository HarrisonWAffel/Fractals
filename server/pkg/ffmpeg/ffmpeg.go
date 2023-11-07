package ffmpeg

import (
	"bytes"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/png"
	"io"
)

type Processor struct {
}

func (p *Processor) CreateVideo(frames []*image.RGBA, writeCloser io.WriteCloser) error {

	videoBuffer := bytes.Buffer{}

	// todo; if we are going to stream it we need to
	// 	use an open ended byte provider, not just a buffer that can run dry
	for _, frame := range frames {
		frameBuff := new(bytes.Buffer)
		err := png.Encode(frameBuff, frame)
		if err != nil {
			return fmt.Errorf("encountered error converting frame to png: %v", err)
		}
		videoBuffer.Write(frameBuff.Bytes())
	}

	// jpeg is 24.8 mb png is 21.3

	// need to feed encoded png frames into here
	err := ffmpeg.Input("pipe:").
		WithInput(bytes.NewReader(videoBuffer.Bytes())).
		Output("pipe:", ffmpeg.KwArgs{
			"c:v":     "libx264",
			"pix_fmt": "yuv420p",
			"f":       "ismv",
		}).
		WithOutput(writeCloser).
		ErrorToStdOut().
		Run()
	if err != nil {
		return err
	}

	return writeCloser.Close()
}
