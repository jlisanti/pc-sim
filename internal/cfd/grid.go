package cfd

import (
	"math"
)

type Grid struct {
	Points       int64
	U            []conservative
	Un           []conservative
	W            []primitive
	WorkingFluid fluid
	Area         []float64
	ID           []float64
	Dz           float64
	//MaxMin       varExtremes
}

type conservative struct {
	Rho  float64
	RhoV float64
	RhoE float64
}

type primitive struct {
	P float64
	E float64
	T float64
	V float64
	C float64
}

type fluid struct {
	Tair  float64
	Fric  float64
	H     float64
	Cv    float64
	R     float64
	Gamma float64
}

/*
type varExtremes struct {
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
*/

func NewGrid(length int64) *Grid {
	var primitivetmp []primitive
	var conservativetmp []conservative
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
	fluidtmp := fluid{
		Tair:  Tair,
		Fric:  Fric,
		H:     H,
		Cv:    Cv,
		R:     gasConstant,
		Gamma: gamma,
	}
	for i := int64(0); i < length; i++ {
		primitivei := primitive{
			P: rho * gasConstant * T,
			T: T,
			E: Cv * T,
			V: 0.0,
			C: math.Sqrt(gamma * gasConstant * T)}
		conservativei := conservative{
			Rho:  rho,
			RhoV: 0.0,
			RhoE: Cv * T * rho,
		}

		primitivetmp = append(primitivetmp, primitivei)
		conservativetmp = append(conservativetmp, conservativei)

		A = append(A, 1.0)
		ID = append(ID, 0.05)
	}
	nGrid := Grid{
		Points:       length,
		U:            conservativetmp,
		Un:           conservativetmp,
		W:            primitivetmp,
		WorkingFluid: fluidtmp,
		Area:         A,
		ID:           ID,
		Dz:           0.000001,
	}
	return &nGrid
}

func UpdateSubStep(i int, tGrid *Grid, ki float64, li float64, ni float64) {
	tGrid.U[i].Rho = tGrid.U[i].Rho + 0.5*ki
	tGrid.U[i].RhoV = tGrid.U[i].RhoV + 0.5*li
	tGrid.U[i].RhoE = tGrid.U[i].RhoE + 0.5*ni
	tGrid.W[i].V = tGrid.U[i].RhoV / tGrid.U[i].Rho
	tGrid.W[i].E = tGrid.U[i].RhoE / tGrid.U[i].Rho
	tGrid.W[i].T = (1.0 / tGrid.WorkingFluid.Cv) * (tGrid.W[i].E - math.Pow(tGrid.W[i].V, 2)/2.0)
	tGrid.W[i].P = tGrid.U[i].Rho * tGrid.WorkingFluid.R * tGrid.W[i].T
	tGrid.W[i].C = math.Sqrt(tGrid.WorkingFluid.Gamma * tGrid.WorkingFluid.R * tGrid.W[i].T)
	/*
		if i == 999 {
			fmt.Println(tGrid.Rho[i], tGrid.RhoU[i], tGrid.RhoE[i])

		}
	*/
}

func UpdateStep(i int, tGrid *Grid, ki [4]float64, li [4]float64, ni [4]float64) {
	tGrid.U[i].Rho = tGrid.Un[i].Rho + (1.0/6.0)*(ki[0]+2.0*ki[1]+2.0*ki[2]+ki[3])
	tGrid.U[i].RhoV = tGrid.Un[i].RhoV + (1.0/6.0)*(li[0]+2.0*li[1]+2.0*li[2]+li[3])
	tGrid.U[i].RhoE = tGrid.Un[i].RhoE + (1.0/6.0)*(ni[0]+2.0*ni[1]+2.0*ni[2]+ni[3])
	tGrid.W[i].V = tGrid.U[i].RhoV / tGrid.U[i].Rho
	tGrid.W[i].E = tGrid.U[i].RhoE / tGrid.U[i].Rho
	tGrid.W[i].T = (1.0 / tGrid.WorkingFluid.Cv) * (tGrid.W[i].E - math.Pow(tGrid.W[i].V, 2)/2.0)
	tGrid.W[i].P = tGrid.U[i].Rho * tGrid.WorkingFluid.R * tGrid.W[i].T
	tGrid.W[i].C = math.Sqrt(tGrid.WorkingFluid.Gamma * tGrid.WorkingFluid.R * tGrid.W[i].T)
	/*
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
	*/
}
