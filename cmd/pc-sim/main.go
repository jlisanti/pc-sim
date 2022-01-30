package main

import (
	"fmt"
	"math"

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
	t := float64(0.0)
	dt := float64(0.000001)
	cflNum := float64(0.01)

	tMax := float64(0.000000001)
	NaNquit := false

	// time iterate
	for {
		dt = cfd.Cfl(cGrid, cflNum)
		t += dt
		fmt.Println("t = ", t)
		fmt.Println("Umax = ", cGrid.Umax, " at: ", cGrid.Umaxi, " Pmax = ", cGrid.Pmax, " at: ", cGrid.Pmaxi,
			"Tmax = ", cGrid.Tmax, " at: ", cGrid.Tmaxi, " rhomax = ", cGrid.Rhomax, " at: ", cGrid.Rhomaxi)
		fmt.Println("Umin = ", cGrid.Umin, " at: ", cGrid.Umini, " Pmin = ", cGrid.Pmin, " at: ", cGrid.Pmini,
			"Tmin = ", cGrid.Tmin, " at: ", cGrid.Tmini, " rhomin = ", cGrid.Rhomin, " at: ", cGrid.Rhomini)
		// loop over rk coefficients
		for {
			// loop over grid
			//fmt.Println("rkIndex = ", rkIndex)
			for i := 0; i < int(cGrid.Points); i++ {
				cGrid.Rhoi[i] = cGrid.Rho[i]
				cGrid.RhoUi[i] = cGrid.RhoU[i]
				cGrid.RhoEi[i] = cGrid.RhoE[i]
			}
			if rkIndex == 3 {
				for i := 0; i < int(cGrid.Points); i++ {
					for j := range rkCoeffs[i].k {
						if math.IsNaN(rkCoeffs[i].k[j]) {
							fmt.Println("rkCoeff k[", j, "] is Nan at grid point: ", i)
							NaNquit = true
						}
						if math.IsNaN(rkCoeffs[i].l[j]) {
							fmt.Println("rkCoeff l[", j, "] is Nan at grid point: ", i)
							NaNquit = true
						}
						if math.IsNaN(rkCoeffs[i].n[j]) {
							fmt.Println("rkCoeff n[", j, "] is Nan at grid point: ", i)
							NaNquit = true
						}
						if NaNquit {
							panic("cannot recover, quiting...")
						}
					}
					cfd.UpdateStep(i, &cGrid, rkCoeffs[i].k, rkCoeffs[i].l, rkCoeffs[i].n)
				}
				break
			} else {
				if rkIndex != 0 {
					for i := 0; i < int(cGrid.Points); i++ {
						rkCoeffs[i].k[rkIndex] = dt * F1(i, t, cGrid)
						rkCoeffs[i].l[rkIndex] = dt * F2(i, t, cGrid)
						rkCoeffs[i].n[rkIndex] = dt * F3(i, t, cGrid)
					}
					for i := 0; i < int(cGrid.Points); i++ {
						for j := range rkCoeffs[i].k {
							if math.IsNaN(rkCoeffs[i].k[j]) {
								fmt.Println("rkCoeff k[", j, "] is Nan at grid point: ", i)
								NaNquit = true
							}
							if math.IsNaN(rkCoeffs[i].l[j]) {
								fmt.Println("rkCoeff l[", j, "] is Nan at grid point: ", i)
								NaNquit = true
							}
							if math.IsNaN(rkCoeffs[i].n[j]) {
								fmt.Println("rkCoeff n[", j, "] is Nan at grid point: ", i)
								NaNquit = true
							}
							if NaNquit {
								panic("cannot recover, quiting...")
							}
						}
						cfd.UpdateSubStep(i, &cGrid, rkCoeffs[i].k[rkIndex], rkCoeffs[i].l[rkIndex], rkCoeffs[i].n[rkIndex])
					}
					rkIndex++
				} else {
					for i := 0; i < int(cGrid.Points); i++ {
						rkCoeffs[i].k[rkIndex] = dt * F1(i, t, cGrid)
						rkCoeffs[i].l[rkIndex] = dt * F2(i, t, cGrid)
						rkCoeffs[i].n[rkIndex] = dt * F3(i, t, cGrid)
					}
					rkIndex++
				}
			}
		}
		rkIndex = 0
		if t >= tMax {
			fmt.Println("t = ", t, " max time: ", tMax, " reached. Exiting time loop...")
			break
		}
	}
}
