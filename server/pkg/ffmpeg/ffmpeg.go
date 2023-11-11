package ffmpeg

import (
	"bytes"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
)

type Processor struct{}

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

const FPS = 60

func (p *Processor) CreateVideo(frameChan chan FrameChunk) *io.PipeReader {

	/*
		pipes!
		generator -> frameChan -> ffmpeg frameChan reader pipe -> ffmpeg -> frame output pipe writer -> frame output pipe reader -> whatever needs to read frames
	*/

	pr := p.GetChunkReader(frameChan)

	// ffmpeg flags should be updated at some point,
	// resulting artifacts using this encoder of flags is very large
	frameReader, frameWriter := io.Pipe()
	go func(frameReader *io.PipeReader, frameOutputWriter *io.PipeWriter) {
		err := ffmpeg.Input("pipe:").
			WithInput(pr).
			Output("pipe:", ffmpeg.KwArgs{
				"c:v":     "libx264",
				"crf":     "28",
				"preset":  "veryfast",
				"pix_fmt": "yuv420p",
				"f":       "ismv",
			}).
			WithOutput(frameWriter, frameWriter).
			ErrorToStdOut().
			Run()
		if err != nil {
			panic(err)
		}
		frameWriter.Close()
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
