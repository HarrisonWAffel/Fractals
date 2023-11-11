package mandelbrot_set

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"harrisonwaffel/fractals/pkg/ffmpeg"
	"harrisonwaffel/fractals/pkg/util"
	"image/png"
	"log/slog"
)

func MandelbrotMp4Handler(context *gin.Context) {
	params := MandelbrotParams{}
	params.populate(context)
	j, _ := json.MarshalIndent(params, "", " ")
	fmt.Print(string(j))
	mandelBrot := MandelbrotGenerator{
		Ctx:              context.Request.Context(),
		ImageHeight:      1000,
		MoveX:            params.MoveX,
		MoveY:            params.MoveY,
		ZoomStepSize:     params.ZoomStepSize,
		Zoom:             params.Zoom,
		Duration:         params.Duration,
		ImageWidth:       1000,
		ConvergenceLimit: 255,
	}
	mandelBrot.Palette = util.InitPalette()

	frameChan := mandelBrot.GenerateZoomVideo()
	p := ffmpeg.Processor{}
	if err := p.StreamFuncOutput(frameChan, context.Writer); err != nil {
		slog.Info(err.Error())
	}
}

func MandelbrotPngHandler(context *gin.Context) {
	params := MandelbrotParams{}
	params.populate(context)

	mandelBrot := MandelbrotGenerator{
		Ctx:              context.Request.Context(),
		ImageHeight:      1000,
		MoveX:            params.MoveX,
		MoveY:            params.MoveY,
		ZoomStepSize:     params.ZoomStepSize,
		Zoom:             params.Zoom,
		Duration:         params.Duration,
		ImageWidth:       1000,
		ConvergenceLimit: 255,
	}
	mandelBrot.Palette = util.InitPalette()
	img := mandelBrot.GenerateImage(mandelBrot.Zoom)

	frameBuff := new(bytes.Buffer)
	png.Encode(frameBuff, img)
	context.Writer.Write(frameBuff.Bytes())
}

type MandelbrotParams struct {
	Duration     int
	Zoom         float64
	ZoomStepSize float64
	MoveX        float64
	MoveY        float64
}

func (mp *MandelbrotParams) populate(ctx *gin.Context) {
	mp.Duration = util.GinParamToInt(ctx.Query("duration"))
	mp.Zoom = util.GinParamToFloat64(ctx.Query("zoom"))
	mp.ZoomStepSize = util.GinParamToFloat64(ctx.Query("zoom-step-size"))
	mp.MoveX = util.GinParamToFloat64(ctx.Query("movex"))
	mp.MoveY = util.GinParamToFloat64(ctx.Query("movey"))
}
