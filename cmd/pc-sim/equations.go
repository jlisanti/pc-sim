package main

import (
	"math"

	"github.com/jlisanti/pc-sim/internal/cfd"
)

func F1(i int, cGrid cfd.Grid) float64 {
	// Forward difference
	A := 0.0
	if i > int(cGrid.Points)-1 {
		A = (cGrid.RhoU[i]*cGrid.Area[i] - cGrid.RhoU[i-1]*cGrid.Area[i-1]) / cGrid.Dz
	} else {
		A = (cGrid.RhoU[i+1]*cGrid.Area[i+1] - cGrid.RhoU[i]*cGrid.Area[i]) / cGrid.Dz
	}
	return A
}

func F2(i int, cGrid cfd.Grid) float64 {
	A, B, C, D := 0.0, 0.0, 0.0, 0.0
	if i > int(cGrid.Points)-1 {
		// Backward difference
		A = cGrid.Rho[i]*math.Pow(cGrid.RhoU[i+1]/cGrid.Rho[i], 2) + cGrid.P[i]*cGrid.Area[i]
		B = cGrid.Rho[i-1]*math.Pow(cGrid.RhoU[i-1]/cGrid.Rho[i-1], 2) + cGrid.P[i-1]*cGrid.Area[i-1]
		C = cGrid.P[i] * ((cGrid.Area[i] - cGrid.Area[i-1]) / cGrid.Dz)
		D = cGrid.Rho[i] * cGrid.Area[i] * ((4.0 * cGrid.Fric) / cGrid.ID[i]) * math.Pow(cGrid.U[i], 2) * 0.5 * cGrid.U[i] * (1.0 / math.Abs(cGrid.U[i]))

	} else {
		// Forward difference
		A = cGrid.Rho[i+1]*math.Pow(cGrid.RhoU[i+1]/cGrid.Rho[i+1], 2) + cGrid.P[i+1]*cGrid.Area[i+1]
		B = cGrid.Rho[i]*math.Pow(cGrid.RhoU[i]/cGrid.Rho[i], 2) + cGrid.P[i]*cGrid.Area[i]
		C = cGrid.P[i] * ((cGrid.Area[i+1] - cGrid.Area[i]) / cGrid.Dz)
		D = cGrid.Rho[i] * cGrid.Area[i] * ((4.0 * cGrid.Fric) / cGrid.ID[i]) * math.Pow(cGrid.U[i], 2) * 0.5 * cGrid.U[i] * (1.0 / math.Abs(cGrid.U[i]))
	}

	return ((A - B) / cGrid.Dz) + C - D
}

func F3(i int, cGrid cfd.Grid) float64 {
	A, B, C, D := 0.0, 0.0, 0.0, 0.0
	if i > int(cGrid.Points)-1 {
		// Backward difference
		A = cGrid.U[i] * (cGrid.Rho[i]*cGrid.Area[i]*cGrid.E[i] + cGrid.P[i]*cGrid.Area[i])
		B = cGrid.U[i-1] * (cGrid.Rho[i-1]*cGrid.Area[i-1]*cGrid.E[i-1] + cGrid.P[i-1]*cGrid.Area[i-1])
		C = 0.0 // Q value
		D = 4.0 * cGrid.ID[i] * cGrid.H * (cGrid.T[i] - cGrid.Tair)
	} else {
		// Forward difference
		A = cGrid.U[i+1] * (cGrid.Rho[i+1]*cGrid.Area[i+1]*cGrid.E[i+1] + cGrid.P[i+1]*cGrid.Area[i+1])
		B = cGrid.U[i] * (cGrid.Rho[i]*cGrid.Area[i]*cGrid.E[i] + cGrid.P[i]*cGrid.Area[i])
		C = 0.0 // Q value
		D = 4.0 * cGrid.ID[i] * cGrid.H * (cGrid.T[i] - cGrid.Tair)
	}

	return ((A - B) / cGrid.Dz) + C + D
}
