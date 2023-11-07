package julia_set

import (
	"fmt"
	"harrisonwaffel/fractals/pkg/ffmpeg"
	"net/http"

	"io"
	"sync"
)

func (js *JuliaSet) StreamFuncOutput(moveX, moveY, zoom float32, w http.ResponseWriter) error {
	frameChan := js.GenerateSet(moveX, moveY, zoom)
	ffmpegProcessor := ffmpeg.Processor{}
	frameOutput := ffmpegProcessor.CreateVideo(frameChan)

	wg := sync.WaitGroup{}
	wg.Add(1)
	errChan := make(chan error)
	go func(frameOutput *io.PipeReader, w http.ResponseWriter, errChan chan error) {
		for {
			n, err := io.Copy(w, frameOutput)
			if err != nil {
				// this err could also be an EOF
				errChan <- err
				wg.Done()
				close(errChan)
				return
			}
			if n == 0 {
				wg.Done()
				close(errChan)
				return
			}
		}
	}(frameOutput, w, errChan)
	wg.Wait()

	select {
	case err, open := <-errChan:
		if !open {
			return nil
		}
		return err
	default:
	}
	fmt.Println("ret")
	return nil
}
