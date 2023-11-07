package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	julia_set "harrisonwaffel/fractals/pkg/julia-set"
	mandelbrot_set "harrisonwaffel/fractals/pkg/mandelbrot-set"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	// ignore the origin of the request
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartServer() error {
	e := gin.Default()

	e.GET("/mandelbrot.png", func(context *gin.Context) {
		moveX := context.Query("movex")
		moveY := context.Query("movey")
		zoom := context.Query("zoom")

		moveXFloat, err := strconv.ParseFloat(moveX, 64)
		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}
		moveYFloat, err := strconv.ParseFloat(moveY, 64)
		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}
		zoomFloat, err := strconv.ParseFloat(zoom, 64)
		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		mandelBrot := mandelbrot_set.MandelbrotGenerator{
			Ctx:              context.Request.Context(),
			ImageHeight:      1000,
			MoveX:            float32(moveXFloat),
			MoveY:            float32(moveYFloat),
			Zoom:             float32(zoomFloat),
			ImageWidth:       1000,
			ConvergenceLimit: 255,
		}

		// a cool one is
		// move X 20500
		// move y 2750
		// zoom 125

		img, err := mandelBrot.GenerateImage()
		if err != nil {
			context.Status(http.StatusInternalServerError)
			return
		}

		context.Writer.Write(img)
	})

	e.GET("/julia-set.mp4", func(context *gin.Context) {

		cReal := context.Query("constant-real")
		cImaginary := context.Query("constant-imaginary")
		totalRange := context.Query("total-range")
		stepSize := context.Query("step-size")
		zoom := context.Query("zoom")

		cRealFloat, err := strconv.ParseFloat(cReal, 64)
		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		cImaginaryFloat, err := strconv.ParseFloat(cImaginary, 64)
		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		totalRangeFloat, err := strconv.ParseFloat(totalRange, 64)
		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		stepSizeFloat, err := strconv.ParseFloat(stepSize, 64)
		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		zoomFloat, err := strconv.ParseFloat(zoom, 64)
		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		js := julia_set.JuliaSet{
			Ctx:                 context.Request.Context(),
			InitialConstantReal: float32(cRealFloat),
			ConstantImaginary:   float32(cImaginaryFloat),
			TotalRange:          float32(totalRangeFloat),
			StepSize:            float32(stepSizeFloat),
			VideoHeight:         1000,
			VideoWidth:          1000,
		}

		err = js.StreamFuncOutput(0.0, 0.0, float32(zoomFloat), context.Writer)
		if err != nil {
			context.Status(http.StatusInternalServerError)
			return
		}
		fmt.Println("done")
	})

	return e.Run(":8989")
}
