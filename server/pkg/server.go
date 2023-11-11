package pkg

import (
	"github.com/gin-gonic/gin"
	juliaset "harrisonwaffel/fractals/pkg/julia-set"
	mandelbrotset "harrisonwaffel/fractals/pkg/mandelbrot-set"
)

func StartServer() error {
	e := gin.Default()

	e.GET("/mandelbrot.png", mandelbrotset.MandelbrotPngHandler)

	e.GET("/mandelbrot.mp4", mandelbrotset.MandelbrotMp4Handler)

	e.GET("/julia-set.mp4", juliaset.JuliaSetMp4Handler)

	return e.Run(":8989")
}
