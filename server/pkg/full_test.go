package pkg

import (
	"bytes"
	julia_set "harrisonwaffel/fractals/pkg/julia-set"
	"os"
	"testing"
)

func TestIt(t *testing.T) {

	js := julia_set.JuliaSet{
		InitialConstantReal: -0.8,
		ConstantImaginary:   0.156,
		TotalRange:          0.0012,
		StepSize:            0.00001,
		VideoHeight:         1200,
		VideoWidth:          1200,
	}

	b := new(bytes.Buffer)
	err := js.StreamFuncOutput(0., 0., 1.0, b)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	os.WriteFile("test.mp4", b.Bytes(), os.ModePerm)
}
