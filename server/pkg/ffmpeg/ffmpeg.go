package ffmpeg

import (
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
)

type Processor struct {
}

type Frame struct {
	Index int
	Frame []byte
}

type FrameChunk struct {
	Index  int
	Frames []Frame
}

func (p *Processor) CreateVideo(frameChan chan Frame) *io.PipeReader {

	/*
		pipes!
		julia-set-generator -> frameChan -> ffmpeg frameChan reader pipe -> ffmpeg -> frame output pipe writer -> frame output pipe reader -> whatever needs to read frames
	*/

	pr, pw := io.Pipe()
	go func(frameChan chan Frame, pw *io.PipeWriter) {
		for {
			select {
			case frame, open := <-frameChan:
				if !open {
					pw.Close()
					return
				}
				pw.Write(frame.Frame)
			}
		}
	}(frameChan, pw)

	frameReader, frameWriter := io.Pipe()
	go func(frameReader *io.PipeReader, frameOutputWriter *io.PipeWriter) {
		ffmpeg.Input("pipe:", ffmpeg.KwArgs{
			"flush_packets": "1",
		}).
			WithInput(pr).
			Output("pipe:", ffmpeg.KwArgs{
				"flush_packets":        "1",
				"probesize":            "200000",
				"max_interleave_delta": "1000000",
				"c:v":                  "libx264",
				"pix_fmt":              "yuv420p",
				"f":                    "ismv",
			}).
			WithOutput(frameWriter, frameWriter).
			ErrorToStdOut().
			Run()
		frameWriter.Close()
		fmt.Println("ffmpeg is done")
	}(pr, frameWriter)

	return frameReader
}
