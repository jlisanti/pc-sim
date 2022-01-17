package finitedifference

import "github.com/jlisanti/pc-sim/internal/finitedifference"

func F1(i, int, dz float64, rho []float64, rhoU []float64, rhoE []float64) []float64 {
	return finitedifference.Forward1(i, rhoU, dz)
}
