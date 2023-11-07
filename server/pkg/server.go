package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	julia_set "harrisonwaffel/fractals/pkg/julia-set"
	"net/http"
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

	e.GET("/fractal.mp4", func(context *gin.Context) {

		js := julia_set.JuliaSet{
			Ctx:                 context.Request.Context(),
			InitialConstantReal: 0.280,
			ConstantImaginary:   0.01,
			TotalRange:          0.005,
			StepSize:            0.00001,
			VideoHeight:         1000,
			VideoWidth:          1000,
		}

		err := js.StreamFuncOutput(0.0, 0.0, 1, context.Writer)
		if err != nil {
			context.Status(http.StatusInternalServerError)
			return
		}
		fmt.Println("done")
	})

	return e.Run(":8989")
}
