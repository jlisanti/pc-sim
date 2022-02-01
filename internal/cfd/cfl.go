package cfd

import (
	"math"
)

func Cfl(cGrid Grid, cflNum float64) float64 {
	var sMax float64
	for j := 0; j < int(cGrid.Points); j++ {
		sTmp := math.Abs(cGrid.W[j].V + cGrid.W[j].C)
		if sTmp > sMax {
			sMax = sTmp
		}
	}
	return cflNum * cGrid.Dz / sMax
}
