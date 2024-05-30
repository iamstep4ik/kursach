package lib

import (
	"math"
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
