package julia_set

import (
	"github.com/gin-gonic/gin"
	"harrisonwaffel/fractals/pkg/ffmpeg"
	"harrisonwaffel/fractals/pkg/util"
	"log/slog"
	"net/http"
)

func JuliaSetMp4Handler(context *gin.Context) {
	params := JuliaSetQueryParams{}
	params.populate(context)

	js := JuliaSet{
		Ctx:                 context.Request.Context(),
		InitialConstantReal: params.ConstantReal,
		ConstantImaginary:   params.ConstantImaginary,
		TotalRange:          params.TotalRange,
		StepSize:            params.StepSize,
		VideoHeight:         1000,
		VideoWidth:          1000,
	}

	frameChan := js.GenerateSet(params.MoveX, params.MoveY, params.Zoom)

	ffmpegProcessor := ffmpeg.Processor{}
	err := ffmpegProcessor.StreamFuncOutput(frameChan, context.Writer)
	if err != nil {
		slog.Error("error encountered streaming video: %v\n", err)
		context.Status(http.StatusInternalServerError)
	}
}

type JuliaSetQueryParams struct {
	ConstantReal      float32
	ConstantImaginary float32
	TotalRange        float32
	Zoom              float32
	StepSize          float32
	MoveX             float32
	MoveY             float32
}

func (jp *JuliaSetQueryParams) populate(context *gin.Context) {
	jp.ConstantReal = util.GinParamToFloat32(context.Query("constant-real"))
	jp.ConstantImaginary = util.GinParamToFloat32(context.Query("constant-imaginary"))
	jp.TotalRange = util.GinParamToFloat32(context.Query("total-range"))
	jp.StepSize = util.GinParamToFloat32(context.Query("step-size"))
	jp.MoveX = util.GinParamToFloat32(context.Query("movex"))
	jp.MoveY = util.GinParamToFloat32(context.Query("movey"))
	jp.Zoom = util.GinParamToFloat32(context.Query("zoom"))
}
