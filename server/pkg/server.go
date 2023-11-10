package pkg

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"harrisonwaffel/fractals/pkg/ffmpeg"
	julia_set "harrisonwaffel/fractals/pkg/julia-set"
	mandelbrot_set "harrisonwaffel/fractals/pkg/mandelbrot-set"
	"harrisonwaffel/fractals/pkg/util"
	"image/png"
	"strconv"
)

func StartServer() error {
	e := gin.Default()

	e.GET("/mandelbrot.png", func(context *gin.Context) {
		values := []string{
			"movex",
			"movey",
			"zoom",
		}

		output := getFloat64QueryParams(values, context)

		mandelBrot := mandelbrot_set.MandelbrotGenerator{
			Ctx:              context.Request.Context(),
			ImageHeight:      1000,
			MoveX:            output[0],
			MoveY:            output[1],
			Zoom:             output[2],
			ImageWidth:       1000,
			ConvergenceLimit: 255,
		}
		mandelBrot.Palette = util.InitPalette()

		// a cool one is
		// move X 148.5
		// move y 9.5
		// zoom 1000

		// 529 ms vs 647, so a little better

		img := mandelBrot.GenerateImage(mandelBrot.Zoom)

		frameBuff := new(bytes.Buffer)
		png.Encode(frameBuff, img)

		context.Writer.Write(frameBuff.Bytes())
	})

	e.GET("/mandelbrot.mp4", func(context *gin.Context) {
		values := []string{
			"movex",
			"movey",
			"zoom",
			"zoom-step",
			"seconds",
		}

		output := getFloat64QueryParams(values, context)

		mandelBrot := mandelbrot_set.MandelbrotGenerator{
			Ctx:              context.Request.Context(),
			ImageHeight:      1000,
			MoveX:            output[0],
			MoveY:            output[1],
			ZoomStepSize:     output[2],
			Zoom:             output[3],
			Duration:         int(output[5]),
			ImageWidth:       1000,
			ConvergenceLimit: 255,
		}
		mandelBrot.Palette = util.InitPalette()

		// a cool one is
		// move X 148.5
		// move y 9.5
		// zoom 1000

		p := ffmpeg.Processor{}

		frameChan := mandelBrot.GenerateZoomVideo()
		if err := p.StreamFuncOutput(frameChan, context.Writer); err != nil {
			fmt.Printf("err %v\n", err)
		}
	})

	e.GET("/julia-set.mp4", func(context *gin.Context) {

		querys := []string{
			"constant-real",
			"constant-imaginary",
			"total-range",
			"step-size",
			"zoom",
		}
		values := getFloat32QueryParams(querys, context)

		js := julia_set.JuliaSet{
			Ctx:                 context.Request.Context(),
			InitialConstantReal: values[0],
			ConstantImaginary:   values[1],
			TotalRange:          values[2],
			StepSize:            values[3],
			VideoHeight:         1000,
			VideoWidth:          1000,
		}

		ffmpegProcessor := ffmpeg.Processor{}

		frameChan := js.GenerateSet(0.0, 0.0, values[4])

		err := ffmpegProcessor.StreamFuncOutput(frameChan, context.Writer)
		if err != nil {
			fmt.Print("julia set err: %v\n", err)
		}

		fmt.Println("done")
	})

	return e.Run(":8989")
}

func getFloat64QueryParams(params []string, ctx *gin.Context) []float64 {
	var output []float64
	for _, e := range params {
		v := ctx.Query(e)
		if v == "" {
			output = append(output, 0.0)
		}
		out, _ := strconv.ParseFloat(v, 64)
		output = append(output, out)
	}
	return output
}

func getFloat32QueryParams(params []string, ctx *gin.Context) []float32 {
	var output []float32
	for _, e := range params {
		v := ctx.Query(e)
		if v == "" {
			output = append(output, 0.0)
		}
		out, _ := strconv.ParseFloat(v, 64)
		output = append(output, float32(out))
	}
	return output
}
