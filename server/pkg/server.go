package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	julia_set "harrisonwaffel/fractals/pkg/julia-set"
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

		j, _ := json.MarshalIndent(js, "", " ")
		fmt.Println(string(j))

		err = js.StreamFuncOutput(0.0, 0.0, float32(zoomFloat), context.Writer)
		if err != nil {
			context.Status(http.StatusInternalServerError)
			return
		}
		fmt.Println("done")
	})

	return e.Run(":8989")
}
