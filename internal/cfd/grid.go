package cfd

import "math"

type Grid struct {
	Points int64
	Area   []float64
	ID     []float64
	Rho    []float64
	RhoU   []float64
	RhoE   []float64
	P      []float64
	E      []float64
	T      []float64
	U      []float64
	Dz     float64
	Tair   float64
	Fric   float64
	H      float64
	Cv     float64
	R      float64
}

func NewGrid(length int64) *Grid {
	nGrid := Grid{Points: length}
	return &nGrid
}

func UpdateSubStep(i int, tGrid Grid, ki float64, li float64, ni float64) {
	tGrid.Rho[i] = tGrid.Rho[i] + 0.5*ki
	tGrid.RhoU[i] = tGrid.RhoU[i] + 0.5*li
	tGrid.RhoE[i] = tGrid.RhoE[i] + 0.5*ni
	tGrid.U[i] = tGrid.RhoU[i] / tGrid.Rho[i]
	tGrid.E[i] = tGrid.RhoE[i] / tGrid.Rho[i]
	tGrid.T[i] = (1.0 / tGrid.Cv) * (tGrid.E[i] - math.Pow(tGrid.U[i], 2))
	tGrid.P[i] = tGrid.Rho[i] * tGrid.R * tGrid.T[i]
}

func UpdateStep(i int, tGrid Grid, ki [4]float64, li [4]float64, ni [4]float64) {
	tGrid.Rho[i] = tGrid.Rho[i] + (1.0/6.0)*(ki[0]+2.0*ki[1]+2.0*ki[2]+ki[3])
	tGrid.RhoU[i] = tGrid.RhoU[i] + (1.0/6.0)*(li[0]+2.0*li[1]+2.0*li[2]+li[3])
	tGrid.RhoE[i] = tGrid.RhoE[i] + (1.0/6.0)*(ni[0]+2.0*ni[1]+2.0*ni[2]+ni[3])
	tGrid.U[i] = tGrid.RhoU[i] / tGrid.Rho[i]
	tGrid.E[i] = tGrid.RhoE[i] / tGrid.Rho[i]
	tGrid.T[i] = (1.0 / tGrid.Cv) * (tGrid.E[i] - math.Pow(tGrid.U[i], 2))
	tGrid.P[i] = tGrid.Rho[i] * tGrid.R * tGrid.T[i]
}
