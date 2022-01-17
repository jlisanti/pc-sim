package cfd

import "fmt"

func Forward1(i int, scalar []float64, dz float64) float64 {
	var difference float64
	if i > len(scalar)-1 {
		fmt.Println("Boundary point, can't compute forward difference")

	} else {
		difference = (scalar[i+1] - scalar[i]) / dz
	}
	return difference
}

func Backward1(i int, scalar []float64, dz float64) float64 {
	var difference float64
	if i < 1 {
		fmt.Println("Boundary point, can't compute backward difference")

	} else {
		difference = (scalar[i] - scalar[i-1]) / dz
	}
	return difference
}

func Center1(i int, scalar []float64, dz float64) float64 {
	var difference float64
	if i < 1 && i > len(scalar)-1 {
		fmt.Println("Boundary point, can't compute center difference")

	} else {
		difference = (scalar[i+1] - scalar[i-1]) / (2.0 * dz)
	}
	return difference
}

func Forward2(i int, scalar []float64, dz float64) float64 {
	var difference float64
	if i > len(scalar)-1 {
		fmt.Println("Boundary point, can't compute 2nd order forward difference")

	} else {
		difference = (-3.0*scalar[i] + 4.0*scalar[i+1] - scalar[i+2]) / (2.0 * dz)
	}
	return difference
}

func Backward2(i int, scalar []float64, dz float64) float64 {
	var difference float64
	if i < 2 {
		fmt.Println("Boundary point, can't compute 2nd order backward difference")

	} else {
		difference = (scalar[i-2] - 4.0*scalar[i-1] + 3.0*scalar[i]) / (2.0 * dz)
	}
	return difference
}

func Center4(i int, scalar []float64, dz float64) float64 {
	var difference float64
	if i < 3 && i > len(scalar)-3 {
		fmt.Println("Boundary point, can't compute 4th order center difference")

	} else {
		difference = (scalar[i-2] - 8.0*scalar[i-1] + 8.0*scalar[i+1] - scalar[i+2]) / (12.0 * dz)
	}
	return difference
}
