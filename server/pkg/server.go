package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	julia_set "harrisonwaffel/fractals/pkg/julia-set"
	"net/http"
	"sync"
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

	e.GET("/fractal", func(context *gin.Context) {
		conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			fmt.Printf("bad upgrade: %v\n", err)
			return
		}
		defer conn.Close()

		mux := sync.Mutex{}

		for {
			_, m, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}
			locked := mux.TryLock()
			if !locked {
				continue
			}
			switch string(m) {
			case "julia-set":
				fmt.Println("beginning to generate a julia set")
				js := julia_set.JuliaSet{
					InitialConstantReal: -0.8,
					ConstantImaginary:   0.156,
					TotalRange:          0.0012,
					StepSize:            0.000001,
					VideoHeight:         800,
					VideoWidth:          800,
				}

				err = js.StreamFuncOutput(0.0, 0.0, 1.0, conn)
				if err != nil {
					context.Status(http.StatusInternalServerError)
					return
				}

			default:
				fmt.Println("invalid fractal name ", string(m))
			}
			mux.Unlock()
		}
		fmt.Println("conn was closed")
	})

	return e.Run(":8989")
}
