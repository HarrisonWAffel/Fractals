package julia_set

import (
	"github.com/gorilla/websocket"
	"harrisonwaffel/fractals/pkg/ffmpeg"

	"io"
	"sync"
)

func (js *JuliaSet) StreamFuncOutput(moveX, moveY, zoom float32, conn *websocket.Conn) error {
	frameChan := js.GenerateSet(moveX, moveY, zoom)
	ffmpegProcessor := ffmpeg.Processor{}
	frameOutput := ffmpegProcessor.CreateVideo(frameChan)

	wg := sync.WaitGroup{}
	wg.Add(1)
	errChan := make(chan error)
	go func(frameOutput *io.PipeReader, conn *websocket.Conn, errChan chan error) {
		for {
			newWriter, err := conn.NextWriter(websocket.BinaryMessage)
			b := make([]byte, 4096)
			if err != nil {
				panic(err)
			}

			n, err := frameOutput.Read(b)
			if err != nil {
				// this err could also be an EOF
				errChan <- err
				wg.Done()
				close(errChan)
				newWriter.Close()
				return
			}
			if n == 0 {
				wg.Done()
				close(errChan)
				newWriter.Close()
				return
			}
			newWriter.Write(b)
			newWriter.Close()
		}
	}(frameOutput, conn, errChan)
	wg.Wait()

	select {
	case err, open := <-errChan:
		if !open {
			return nil
		}
		return err
	default:
	}
	return nil
}
