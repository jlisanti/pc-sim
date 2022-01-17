package main

import (
	"github.com/jlisanti/pc-sim/internal/cfd"
)

type rkCoeff struct {
	k [4]float64
	l [4]float64
	n [4]float64
}

func main() {

	cGrid := *cfd.NewGrid(1000)

	rkZero := [4]float64{0.0, 0.0, 0.0, 0.0}
	rkCoeffs := []rkCoeff{}

	for i := 0; i <= int(cGrid.Points); i++ {
		rkCoeffs = append(rkCoeffs, rkCoeff{k: rkZero, l: rkZero, n: rkZero})
	}

	// integrate in time
	rkIndex := 0

	// time iterate
	for {
		// loop over rk coefficients
		for rkIndex < 4 {
			// loop over grid
			if rkIndex != 4 {
				if rkIndex != 0 {
					for i := 0; i <= int(cGrid.Points); i++ {
						rkCoeffs[i].k[rkIndex] = F1(i, cGrid)
						rkCoeffs[i].l[rkIndex] = F2(i, cGrid)
						rkCoeffs[i].n[rkIndex] = F3(i, cGrid)
						cfd.UpdateSubStep(i, cGrid, rkCoeffs[i].k[rkIndex], rkCoeffs[i].l[rkIndex], rkCoeffs[i].n[rkIndex])
					}
				} else {
					for i := 0; i <= int(cGrid.Points); i++ {
						rkCoeffs[i].k[rkIndex] = F1(i, cGrid)
						rkCoeffs[i].l[rkIndex] = F2(i, cGrid)
						rkCoeffs[i].n[rkIndex] = F3(i, cGrid)
					}
				}
			} else {
				for i := 0; i <= int(cGrid.Points); i++ {
					cfd.UpdateStep(i, cGrid, rkCoeffs[i].k, rkCoeffs[i].l, rkCoeffs[i].n)
				}
				rkIndex++
			}
		}
	}
}
