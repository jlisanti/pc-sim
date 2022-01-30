package cfd

import "math"

func Q(i int, cGrid Grid, t float64, width float64, amp float64, offset float64) float64 {

	i_peak := 10
	i_width := 0.1
	freq := 250.0
	spatialScaling := gaussian(1.0, float64(i_peak), i_width, float64(i))
	cycleTime := 1.0 / freq
	numCycles := math.Floor(t / cycleTime)
	timeScaled := t - float64(numCycles)*cycleTime
	start := (((offset / 360.0) * cycleTime) + ((width / 360.0) * cycleTime))

	if t < start {
		return 0.0
	} else {
		return spatialScaling * gaussianCyclic(
			amp/((width/360.0)*cycleTime*math.Sqrt(2.0*math.Pi)),
			(offset/360.0)*cycleTime,
			((offset+360.0)/360.0)*cycleTime,
			(width/360.0)*cycleTime,
			timeScaled)
	}
}

func gaussian(a float64, b float64, c float64, x float64) float64 {
	return a * math.Exp(-math.Pow(x-b, 2)/(2.0*math.Pow(c, 2)))
}

func gaussianCyclic(a float64, b float64, tb float64, c float64, x float64) float64 {
	return a*math.Exp(-math.Pow(x-b, 2)/(2.0*math.Pow(c, 2))) +
		a*math.Exp(-math.Pow(x-tb, 2)/(2.0*math.Pow(c, 2)))
}

func fwhm(c float64) float64 {
	return 2.0 * math.Sqrt(2.0*math.Log(2.0)) * c
}
