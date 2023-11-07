package util

import (
	"math"
)

// DefaultPalette is the default color gradient to use when generating frames
var DefaultPalette = []string{"#000764", "#206acb", "#edffff", "#ffaa00", "#0002000"}

// MapToRange maps one range to another, e.g. maps
// a float64 range of [0 ,1] to any other range such as [0,360]
func MapToRange(input float32) int {
	inputStart := 0.
	inputEnd := 255.
	outputStart := 0.
	outputEnd := 128.

	slope := 1.0 * (outputEnd - outputStart) / (inputEnd - inputStart)
	return int(outputStart + math.Round(slope*(float64(input)-inputStart)))
}
