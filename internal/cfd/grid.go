package cfd

import (
	"math"
)

type Grid struct {
	Points  int64
	Area    []float64
	ID      []float64
	Rho     []float64
	RhoU    []float64
	RhoE    []float64
	Rhoi    []float64
	RhoUi   []float64
	RhoEi   []float64
	P       []float64
	E       []float64
	T       []float64
	U       []float64
	C       []float64
	Dz      float64
	Tair    float64
	Fric    float64
	H       float64
	Cv      float64
	R       float64
	Gamma   float64
	Umax    float64
	Rhomax  float64
	Pmax    float64
	Tmax    float64
	Umin    float64
	Rhomin  float64
	Pmin    float64
	Tmin    float64
	Umaxi   int
	Rhomaxi int
	Pmaxi   int
	Tmaxi   int
	Umini   int
	Rhomini int
	Pmini   int
	Tmini   int
}

func NewGrid(length int64) *Grid {
	var rhotmp []float64
	var rhoUtmp []float64
	var rhoEtmp []float64
	var rhotmpi []float64
	var rhoUtmpi []float64
	var rhoEtmpi []float64
	var utmp []float64
	var Ttmp []float64
	var Etmp []float64
	var Ptmp []float64
	var Ctmp []float64
	var A []float64
	var ID []float64
	Cv := 718.0
	gasConstant := 287.12
	H := 1.0
	Tair := 300.0
	Fric := 0.1
	T := 300.0
	rho := 1.2
	gamma := 1.4
	for i := int64(0); i < length; i++ {
		rhotmp = append(rhotmp, rho)
		rhoUtmp = append(rhoUtmp, 0.0)
		rhoEtmp = append(rhoEtmp, Cv*T)
		rhotmpi = append(rhotmpi, rho)
		rhoUtmpi = append(rhoUtmpi, 0.0)
		rhoEtmpi = append(rhoEtmpi, Cv*T)
		utmp = append(utmp, 0.0)
		Ttmp = append(Ttmp, T)
		Etmp = append(Etmp, Cv*T)
		Ptmp = append(Ptmp, rho*gasConstant*T)
		Ctmp = append(Ctmp, math.Sqrt(gamma*gasConstant*T))
		A = append(A, 1.0)
		ID = append(ID, 0.05)
	}
	nGrid := Grid{Points: length, Area: A, ID: ID, Dz: 0.0000001, Rho: rhotmp,
		RhoU: rhoUtmp, RhoE: rhoEtmp, U: utmp, T: Ttmp, E: Etmp, P: Ptmp,
		Tair: Tair, Fric: Fric, H: H, Cv: Cv, R: gasConstant, C: Ctmp, Gamma: gamma,
		Tmax: Ttmp[length/2], Pmax: Ptmp[length/2], Rhomax: rhotmp[length/2], Umax: utmp[length/2],
		Tmin: Ttmp[length/2], Pmin: Ptmp[length/2], Rhomin: rhotmp[length/2], Umin: utmp[length/2],
		Rhoi: rhotmpi, RhoUi: rhoUtmpi, RhoEi: rhoEtmpi}
	return &nGrid
}

func UpdateSubStep(i int, tGrid *Grid, ki float64, li float64, ni float64) {
	tGrid.Rho[i] = tGrid.Rho[i] + 0.5*ki
	tGrid.RhoU[i] = tGrid.RhoU[i] + 0.5*li
	tGrid.RhoE[i] = tGrid.RhoE[i] + 0.5*ni
	tGrid.U[i] = tGrid.RhoU[i] / tGrid.Rho[i]
	tGrid.E[i] = tGrid.RhoE[i] / tGrid.Rho[i]
	tGrid.T[i] = (1.0 / tGrid.Cv) * (tGrid.E[i] - math.Pow(tGrid.U[i], 2)/2.0)
	tGrid.P[i] = tGrid.Rho[i] * tGrid.R * tGrid.T[i]
	tGrid.C[i] = math.Sqrt(tGrid.Gamma * tGrid.R * tGrid.T[i])
	/*
		if i == 999 {
			fmt.Println(tGrid.Rho[i], tGrid.RhoU[i], tGrid.RhoE[i])

		}
	*/
}

func UpdateStep(i int, tGrid *Grid, ki [4]float64, li [4]float64, ni [4]float64) {
	tGrid.Rho[i] = tGrid.Rhoi[i] + (1.0/6.0)*(ki[0]+2.0*ki[1]+2.0*ki[2]+ki[3])
	tGrid.RhoU[i] = tGrid.RhoUi[i] + (1.0/6.0)*(li[0]+2.0*li[1]+2.0*li[2]+li[3])
	tGrid.RhoE[i] = tGrid.RhoEi[i] + (1.0/6.0)*(ni[0]+2.0*ni[1]+2.0*ni[2]+ni[3])
	tGrid.U[i] = tGrid.RhoU[i] / tGrid.Rho[i]
	tGrid.E[i] = tGrid.RhoE[i] / tGrid.Rho[i]
	tGrid.T[i] = (1.0 / tGrid.Cv) * (tGrid.E[i] - math.Pow(tGrid.U[i], 2)/2.0)
	tGrid.P[i] = tGrid.Rho[i] * tGrid.R * tGrid.T[i]
	tGrid.C[i] = math.Sqrt(tGrid.Gamma * tGrid.R * tGrid.T[i])
	if tGrid.Rho[i] > tGrid.Rhomax {
		tGrid.Rhomax = tGrid.Rho[i]
		tGrid.Rhomaxi = i
	}
	if tGrid.Rho[i] < tGrid.Rhomin {
		tGrid.Rhomin = tGrid.Rho[i]
		tGrid.Rhomini = i
	}
	if tGrid.U[i] > tGrid.Umax {
		tGrid.Umax = tGrid.U[i]
		tGrid.Umaxi = i
	}
	if tGrid.U[i] < tGrid.Umin {
		tGrid.Umin = tGrid.U[i]
		tGrid.Umini = i
	}
	if tGrid.T[i] > tGrid.Tmax {
		tGrid.Tmax = tGrid.T[i]
		tGrid.Tmaxi = i
	}
	if tGrid.T[i] < tGrid.Tmin {
		tGrid.Tmin = tGrid.T[i]
		tGrid.Tmini = i
	}
	if tGrid.P[i] > tGrid.Pmax {
		tGrid.Pmax = tGrid.P[i]
		tGrid.Pmaxi = i
	}
	if tGrid.P[i] < tGrid.Pmin {
		tGrid.Pmin = tGrid.P[i]
		tGrid.Pmini = i
	}
}
