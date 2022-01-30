package main

import (
	"math"

	"github.com/jlisanti/pc-sim/internal/cfd"
)

func F1(i int, t float64, cGrid cfd.Grid) float64 {
	// Forward difference
	A := 0.0
	if i > int(cGrid.Points)-2 {
		A = (cGrid.RhoU[i]*cGrid.Area[i] - cGrid.RhoU[i-1]*cGrid.Area[i-1]) / cGrid.Dz
	} else {
		A = (cGrid.RhoU[i+1]*cGrid.Area[i+1] - cGrid.RhoU[i]*cGrid.Area[i]) / cGrid.Dz
	}

	/*
		if i > 998 {
			fmt.Println("rho")
			fmt.Println(i, cGrid.RhoU[i], cGrid.Area[i], cGrid.RhoU[i-1], cGrid.Area[i-1])
		}
	*/
	return -A
}

func F2(i int, t float64, cGrid cfd.Grid) float64 {
	A, B, C, D := 0.0, 0.0, 0.0, 0.0
	if i > int(cGrid.Points)-2 {
		// Backward difference
		A = cGrid.Rho[i]*math.Pow(cGrid.U[i], 2)*cGrid.Area[i] +
			cGrid.P[i]*cGrid.Area[i]
		B = cGrid.Rho[i-1]*math.Pow(cGrid.U[i-1], 2)*cGrid.Area[i-1] +
			cGrid.P[i-1]*cGrid.Area[i-1]
		C = 0.0 //cGrid.P[i] * ((cGrid.Area[i] - cGrid.Area[i-1]) / cGrid.Dz)
		D = 0.0 //cGrid.Rho[i] * cGrid.Area[i] * ((4.0 * cGrid.Fric) / cGrid.ID[i]) *
		//math.Pow(cGrid.U[i], 2) * 0.5 * cGrid.U[i] * (1.0 / math.Abs(cGrid.U[i]))
	} else {
		// Forward difference
		A = cGrid.Rho[i+1]*math.Pow(cGrid.U[i+1], 2)*cGrid.Area[i+1] +
			cGrid.P[i+1]*cGrid.Area[i+1]
		B = cGrid.Rho[i]*math.Pow(cGrid.U[i], 2)*cGrid.Area[i] +
			cGrid.P[i]*cGrid.Area[i]
		C = 0.0 //cGrid.P[i] * ((cGrid.Area[i+1] - cGrid.Area[i]) / cGrid.Dz)
		D = 0.0 //cGrid.Rho[i] * cGrid.Area[i] * ((4.0 * cGrid.Fric) / cGrid.ID[i]) * math.Pow(cGrid.U[i], 2) * 0.5 * cGrid.U[i] * (1.0 / math.Abs(cGrid.U[i]))
	}
	/*
		if i > 996 {
			fmt.Println("RhoU equation")
			fmt.Println(i, cGrid.Rho[i], math.Pow(cGrid.U[i], 2), cGrid.Area[i], cGrid.P[i])
			fmt.Println(i, cGrid.Rho[i-1], math.Pow(cGrid.U[i-1], 2), cGrid.Area[i-1], cGrid.P[i-1])
		}
	*/
	return -((A - B) / cGrid.Dz) + (C - D)
}

func F3(i int, t float64, cGrid cfd.Grid) float64 {

	A, B, C, D := 0.0, 0.0, 0.0, 0.0
	//width := 0.1
	//amp := 0.1
	//offset := 0.1
	if i > int(cGrid.Points)-2 {
		// Backward difference
		A = cGrid.U[i] * (cGrid.Rho[i]*cGrid.Area[i]*cGrid.E[i] + cGrid.P[i]*cGrid.Area[i])
		B = cGrid.U[i-1] * (cGrid.Rho[i-1]*cGrid.Area[i-1]*cGrid.E[i-1] + cGrid.P[i-1]*cGrid.Area[i-1])
		C = 0.0 // Q value
		D = 0.0 // 4.0 * cGrid.ID[i] * cGrid.H * (cGrid.T[i] - cGrid.Tair)
	} else {
		// Forward difference
		A = cGrid.U[i+1] * (cGrid.Rho[i+1]*cGrid.Area[i+1]*cGrid.E[i+1] + cGrid.P[i+1]*cGrid.Area[i+1])
		B = cGrid.U[i] * (cGrid.Rho[i]*cGrid.Area[i]*cGrid.E[i] + cGrid.P[i]*cGrid.Area[i])
		C = 0.0 //cfd.Q(i, cGrid, t, width, amp, offset)
		D = 0.0 //4.0 * cGrid.ID[i] * cGrid.H * (cGrid.T[i] - cGrid.Tair)
	}
	/*
		if i > 996 {

			fmt.Println("RhoE equation")
			fmt.Println(i, cGrid.U[i], cGrid.Rho[i], cGrid.Area[i], cGrid.E[i], cGrid.P[i])
			fmt.Println(i, cGrid.U[i-1], cGrid.Rho[i-1], cGrid.Area[i-1], cGrid.E[i-1], cGrid.P[i-1])
		}

	*/
	return -((A - B) / cGrid.Dz) + (C + D)
}
