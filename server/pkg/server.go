package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

	e.GET("/fractal", func(context *gin.Context) {
		conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			fmt.Printf("bad upgrade: %v\n", err)
			return
		}
		defer conn.Close()

		for {
			_, m, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("%v\n", string(m))
		}
		fmt.Println("conn was closed")
	})

	return e.Run(":8989")
}
