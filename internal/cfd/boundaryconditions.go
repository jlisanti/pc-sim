package cfd

import "math"

func PressureInlet(c float64, u float64, rho float64, Po float64, Ps float64, To float64) (rhoIn float64, uIn float64, pIn float64) {
	if u > float64(0.0) {
		pIn = Ps

		gamma := float64(1.4)
		R := float64(287.15)

		//Inflow, compute velocity from isentropic relation
		A := math.Log(Po/Ps) * (gamma - 1.0) / gamma
		B := 2.0 / (gamma - 1.0)

		uIn = math.Sqrt(B*(math.Exp(A)-1.0)) * c

		Msqr := math.Pow(uIn/c, 2)
		TsIn := (2.0 * To) / (2.0 + Msqr*gamma - Msqr)
		rhoIn = Ps / (R * TsIn)

	} else {
		pIn = Ps
		uIn = u
		rhoIn = rho
	}
	return rhoIn, uIn, pIn
}

func PressureOutlet()
