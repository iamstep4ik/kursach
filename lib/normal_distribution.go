package lib

import (
	"math"
)

const (
	E1 = 0.16 * float64((25))
	E2 = 0.34 * float64(25)
	E3 = E2
	E4 = E1
)

func CalculateE(data map[string][]int) []float64 {
	averages := make([]float64, len(data))
	k := 0
	for _, arr := range data {
		sum := 0
		for _, value := range arr {
			sum += value
		}
		averages[k] = float64(sum) / float64(len(arr)) // Use length of array
		k++
	}
	return averages
}

func CalculateStandardDeviation(data map[string][]int, E []float64) []float64 {
	sd := make([]float64, len(data))
	k := 0
	n := 25
	for _, arr := range data {
		sum := 0.0
		for _, value := range arr {
			sum += math.Pow(float64(value)-E[k], 2)
		}
		sd[k] = math.Sqrt(sum / float64(n-1))
		k++
	}
	return sd
}

func CalculateDifferenceESD(E []float64, sd []float64) []float64 {
	difference := make([]float64, len(sd))
	for i, value := range sd {
		difference[i] = E[i] - value
	}
	return difference
}

func CalculateSumESD(E []float64, sd []float64) []float64 {
	sum := make([]float64, len(sd))
	for i, value := range sd {
		sum[i] = E[i] + value
	}
	return sum
}

func CalculateO1(data map[string][]int, difference []float64) []float64 {
	o1 := make([]float64, len(data))
	k := 0
	for _, arr := range data {
		counter := 0
		for _, value := range arr {
			if value <= int(difference[k]) {
				counter++
			}
		}
		o1[k] = float64(counter)
		k++
	}
	return o1
}

func CalculateO2(data map[string][]int, difference []float64, E []float64) []float64 {
	o2 := make([]float64, len(data))
	k := 0
	for _, arr := range data {
		counter := 0
		for _, value := range arr {

			if float64(value) > difference[k] && float64(value) <= E[k] {
				counter++
			}
		}
		o2[k] = float64(counter)
		k++
	}
	return o2
}

func CalculateO3(data map[string][]int, sum []float64, E []float64) []float64 {
	o3 := make([]float64, len(data))
	k := 0
	for _, arr := range data {
		counter := 0
		for _, value := range arr {

			if float64(value) <= sum[k] && float64(value) > E[k] {
				counter++
			}
		}
		o3[k] = float64(counter)
		k++
	}
	return o3
}

func CalculateO4(data map[string][]int, sum []float64) []float64 {
	o4 := make([]float64, len(data))
	k := 0
	for _, arr := range data {
		counter := 0
		for _, value := range arr {

			if float64(value) > sum[k] {
				counter++
			}
		}
		o4[k] = float64(counter)
		k++
	}
	return o4
}

func CalculateX2(o1, o2, o3, o4 []float64) []float64 {
	x2 := make([]float64, len(o1))

	for i := 0; i < len(x2); i++ {
		x2[i] = math.Pow(o1[i]-E1, 2)/E1 + math.Pow(o2[i]-E2, 2)/E2 + math.Pow(o3[i]-E3, 2)/E3 + math.Pow(o4[i]-E4, 2)/E4
	}

	return x2
}

func chiSquarePDF(x float64, df int) float64 {
	if x <= 0 || df <= 0 {
		return 0
	}
	return math.Pow(x, 0.5*float64(df)-1) * math.Exp(-0.5*x) / math.Gamma(0.5*float64(df))
}

func chiSquareCDFApprox(x float64, df int, numSteps int) float64 {
	stepSize := x / float64(numSteps)
	cdf := 0.0
	for i := 0; i <= numSteps; i++ {
		xVal := stepSize * float64(i)
		pdfVal := chiSquarePDF(xVal, df)
		if i == 0 || i == numSteps {
			cdf += pdfVal * stepSize * 0.5
		} else {
			cdf += pdfVal * stepSize
		}
	}
	return cdf
}

func ChiSquareRightTailProbability(x2 []float64, df int) []float64 {
	probabilities := make([]float64, len(x2))
	for i := 0; i < len(probabilities); i++ {
		probabilities[i] = 1 - chiSquareCDFApprox(x2[i], df, 100)
	}
	return probabilities
}
