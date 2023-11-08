package ffmpeg

import (
	"bytes"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
)

type Processor struct {
	TimeoutInSeconds int
}

type Frame struct {
	Frame []byte
}

type FrameChunk struct {
	Frames []Frame
}

func (fc *FrameChunk) ToByteArray() []byte {
	chunkBuff := bytes.Buffer{}
	for _, f := range fc.Frames {
		if f.Frame != nil && len(f.Frame) > 0 {
			chunkBuff.Write(f.Frame)
		}
	}
	return chunkBuff.Bytes()
}

const FPS = 30

func (p *Processor) CreateVideo(frameChan chan FrameChunk) *io.PipeReader {

	/*
		pipes!
		generator -> frameChan -> ffmpeg frameChan reader pipe -> ffmpeg -> frame output pipe writer -> frame output pipe reader -> whatever needs to read frames
	*/

	pr := p.GetChunkReader(frameChan)

	frameReader, frameWriter := io.Pipe()
	go func(frameReader *io.PipeReader, frameOutputWriter *io.PipeWriter) {
		ffmpeg.Input("pipe:").
			WithInput(pr).
			Output("pipe:", ffmpeg.KwArgs{
				"c:v":     "libx264",
				"pix_fmt": "yuv420p",
				"f":       "ismv",
			}).
			WithOutput(frameWriter, frameWriter).
			ErrorToStdOut().
			Run()
		frameWriter.Close()
		fmt.Println("ffmpeg is done")
	}(pr, frameWriter)

	return frameReader
}

func (p *Processor) GetChunkReader(frameChan chan FrameChunk) *io.PipeReader {
	pr, pw := io.Pipe()
	go func(frameChan chan FrameChunk, pw *io.PipeWriter) {
		for {
			select {
			case frameChunk, open := <-frameChan:
				if !open {
					pw.Close()
					return
				}

				pw.Write(frameChunk.ToByteArray())
			}
		}
	}(frameChan, pw)
	return pr
}
