package cfd

import (
	"math"
)

func Cfl(cGrid Grid, cflNum float64) float64 {
	var sMax float64
	for j := 0; j < int(cGrid.Points); j++ {
		sTmp := math.Abs(cGrid.U[j] + cGrid.C[j])
		if sTmp > sMax {
			sMax = sTmp
		}
	}
	return cflNum * cGrid.Dz / sMax
}
