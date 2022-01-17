package main

import (
	"math"

	"github.com/jlisanti/pc-sim/internal/cfd"
)

func F1(i int, cGrid cfd.Grid) float64 {
	// Forward difference
	value := (cGrid.RhoU[i+1]*cGrid.Area[i+1] - cGrid.RhoU[i]*cGrid.Area[i]) / cGrid.Dz
	return value
}

func F2(i int, cGrid cfd.Grid) float64 {
	// Forward difference
	A := cGrid.Rho[i+1]*math.Pow(cGrid.RhoU[i+1]/cGrid.Rho[i+1], 2) + cGrid.P[i+1]*cGrid.Area[i+1]
	B := cGrid.Rho[i+1]*math.Pow(cGrid.RhoU[i+1]/cGrid.Rho[i+1], 2) + cGrid.P[i+1]*cGrid.Area[i+1]
	C := cGrid.P[i] * ((cGrid.Area[i+1] - cGrid.Area[i]) / cGrid.Dz)
	D := cGrid.Rho[i] * cGrid.Area[i] * ((4.0 * cGrid.Fric) / cGrid.ID[i]) * math.Pow(cGrid.U[i], 2) * 0.5 * cGrid.U[i] * (1.0 / math.Abs(cGrid.U[i]))

	value := ((A - B) / cGrid.Dz) + C - D

	return value
}

func F3(i int, cGrid cfd.Grid) float64 {
	// Forward difference
	A := cGrid.U[i+1] * (cGrid.Rho[i+1]*cGrid.Area[i+1]*cGrid.E[i+1] + cGrid.P[i+1]*cGrid.Area[i+1])
	B := cGrid.U[i] * (cGrid.Rho[i]*cGrid.Area[i]*cGrid.E[i] + cGrid.P[i]*cGrid.Area[i])
	C := 0.0 // Q value
	D := 4.0 * cGrid.ID[i] * cGrid.H * (cGrid.T[i] - cGrid.Tair)

	value := ((A - B) / cGrid.Dz) + C + D

	return value
}
