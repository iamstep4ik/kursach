package lib

import "math"

func CalculateA(E []float64, sd []float64) []float64 {
	a := make([]float64, len(E))
	for i, _ := range E {
		a[i] = E[i] - math.Sqrt(3)*sd[i]
	}
	return a
}

func CalculateB(a []float64, E []float64) []float64 {
	b := make([]float64, len(E))
	for i, _ := range E {
		b[i] = E[i] + math.Sqrt(3)*a[i]
	}
	return b
}

func UniO1(data map[string][]int, a []float64) []float64 {
	o1 := make([]float64, len(data))
	k := 0
	for _, arr := range data {
		counter := 0
		for _, value := range arr {
			if float64(value) <= a[k] {
				counter++
			}
		}
		o1[k] = float64(counter)
		k++
	}
	return o1
}

func UniO2(data map[string][]int, a []float64, E []float64) []float64 {
	o2 := make([]float64, len(data))
	k := 0
	for _, arr := range data {
		counter := 0
		for _, value := range arr {
			if float64(value) > a[k] && float64(value) <= E[k] {
				counter++
			}
		}
		o2[k] = float64(counter)
		k++
	}
	return o2
}

func UniO3(data map[string][]int, b []float64, E []float64) []float64 {
	o3 := make([]float64, len(data))
	k := 0
	for _, arr := range data {
		counter := 0
		for _, value := range arr {
			if float64(value) < b[k] && float64(value) >= E[k] {
				counter++
			}
		}
		o3[k] = float64(counter)
		k++
	}
	return o3
}

func UniO4(data map[string][]int, b []float64) []float64 {
	o4 := make([]float64, len(data))
	k := 0
	for _, arr := range data {
		counter := 0
		for _, value := range arr {
			if float64(value) > b[k] {
				counter++
			}
		}
		o4[k] = float64(counter)
		k++
	}
	return o4
}
