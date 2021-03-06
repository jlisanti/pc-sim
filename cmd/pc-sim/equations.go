package main

import (
	"math"

	"github.com/jlisanti/pc-sim/internal/cfd"
)

func F1(i int, t float64, cGrid cfd.Grid) float64 {
	// Forward difference
	A := 0.0
	if i > int(cGrid.Points)-2 {
		A = (cGrid.U[i].RhoV*cGrid.Area[i] - cGrid.U[i-1].RhoV*cGrid.Area[i-1]) / cGrid.Dz
	} else {
		A = (cGrid.U[i+1].RhoV*cGrid.Area[i+1] - cGrid.U[i].RhoV*cGrid.Area[i]) / cGrid.Dz
	}
	return -A
}

func F2(i int, t float64, cGrid cfd.Grid) float64 {
	A, B, C, D := 0.0, 0.0, 0.0, 0.0
	if i > int(cGrid.Points)-2 {
		// Backward difference
		A = cGrid.U[i].Rho*math.Pow(cGrid.W[i].V, 2)*cGrid.Area[i] +
			cGrid.W[i].P*cGrid.Area[i]
		B = cGrid.U[i-1].Rho*math.Pow(cGrid.W[i-1].V, 2)*cGrid.Area[i-1] +
			cGrid.W[i-1].P*cGrid.Area[i-1]
		C = 0.0 //cGrid.P[i] * ((cGrid.Area[i] - cGrid.Area[i-1]) / cGrid.Dz)
		D = 0.0 //cGrid.Rho[i] * cGrid.Area[i] * ((4.0 * cGrid.Fric) / cGrid.ID[i]) *
		//math.Pow(cGrid.U[i], 2) * 0.5 * cGrid.U[i] * (1.0 / math.Abs(cGrid.U[i]))
	} else {
		// Forward difference
		A = cGrid.U[i+1].Rho*math.Pow(cGrid.W[i+1].V, 2)*cGrid.Area[i+1] +
			cGrid.W[i+1].P*cGrid.Area[i+1]
		B = cGrid.U[i].Rho*math.Pow(cGrid.W[i].V, 2)*cGrid.Area[i] +
			cGrid.W[i].P*cGrid.Area[i]
		C = 0.0 //cGrid.P[i] * ((cGrid.Area[i+1] - cGrid.Area[i]) / cGrid.Dz)
		D = 0.0 //cGrid.Rho[i] * cGrid.Area[i] * ((4.0 * cGrid.Fric) / cGrid.ID[i]) * math.Pow(cGrid.U[i], 2) * 0.5 * cGrid.U[i] * (1.0 / math.Abs(cGrid.U[i]))
	}
	return -((A - B) / cGrid.Dz) + (C - D)
}

func F3(i int, t float64, cGrid cfd.Grid) float64 {

	A, B, C, D := 0.0, 0.0, 0.0, 0.0
	//width := 0.1
	//amp := 0.1
	//offset := 0.1
	if i > int(cGrid.Points)-2 {
		// Backward difference
		A = cGrid.W[i].V * (cGrid.U[i].Rho*cGrid.Area[i]*cGrid.W[i].E + cGrid.W[i].P*cGrid.Area[i])
		B = cGrid.W[i-1].V * (cGrid.U[i-1].Rho*cGrid.Area[i-1]*cGrid.W[i-1].E + cGrid.W[i-1].P*cGrid.Area[i-1])
		C = 0.0 // Q value
		D = 0.0 // 4.0 * cGrid.ID[i] * cGrid.H * (cGrid.T[i] - cGrid.Tair)
	} else {
		// Forward difference
		A = cGrid.W[i+1].V * (cGrid.U[i+1].Rho*cGrid.Area[i+1]*cGrid.W[i+1].E + cGrid.W[i+1].P*cGrid.Area[i+1])
		B = cGrid.W[i].V * (cGrid.U[i].Rho*cGrid.Area[i]*cGrid.W[i].E + cGrid.W[i].P*cGrid.Area[i])
		C = 0.0 //cfd.Q(i, cGrid, t, width, amp, offset)
		D = 0.0 //4.0 * cGrid.ID[i] * cGrid.H * (cGrid.T[i] - cGrid.Tair)
	}
	return -((A - B) / cGrid.Dz) + (C + D)
}
