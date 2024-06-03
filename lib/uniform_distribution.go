package lib

import "math"

const (
	uE1 = 0.01 * 25.0
	uE2 = 0.49 * 25.0
	uE3 = 0.49 * 25.0
	uE4 = 0.01 * 25.0
)

type UniformStatistics struct {
	Path   string
	Mu     float64
	Sigma  float64
	A      float64
	B      float64
	O1     float64
	O2     float64
	O3     float64
	O4     float64
	E1     float64
	E2     float64
	E3     float64
	E4     float64
	Chi2   float64
	Chi2Pr float64
}

func InitializeUniformStatistics() UniformStatistics {
	return UniformStatistics{
		E1: uE1,
		E2: uE2,
		E3: uE3,
		E4: uE4,
	}
}

func CalculateAB(data []int, Mu, Sigma float64) (float64, float64) {
	A := Mu - (math.Sqrt(3) * Sigma)
	B := Mu + (math.Sqrt(3) * A)
	return A, B
}

func CalculateUniformO(data []int, A, Mu, B float64) (float64, float64, float64, float64) {
	var o1, o2, o3, o4 float64
	for _, value := range data {
		switch {
		case float64(value) <= A:
			o1++
		case float64(value) > A && float64(value) <= Mu:
			o2++
		case float64(value) >= Mu && float64(value) < B:
			o3++
		case float64(value) > B:
			o4++
		}
	}
	return o1, o2, o3, o4
}

func CalculateUniChi2(o1, o2, o3, o4 float64) float64 {
	return (math.Pow(o1-uE1, 2) / uE1) + (math.Pow(o2-uE2, 2) / uE2) + (math.Pow(o3-uE3, 2) / uE3) + (math.Pow(o4-uE4, 2) / uE4)
}

func CalculateUniAll(data map[string][]int) []UniformStatistics {
	var results []UniformStatistics

	for key, values := range data {
		stat := InitializeUniformStatistics()
		stat.Path = key
		stat.Mu = CalculateMu(values)
		stat.Sigma = CalculateSigma(values, stat.Mu)
		stat.A, stat.B = CalculateAB(values, stat.Mu, stat.Sigma)
		stat.O1, stat.O2, stat.O3, stat.O4 = CalculateUniformO(values, stat.A, stat.Mu, stat.B)
		stat.Chi2 = CalculateUniChi2(stat.O1, stat.O2, stat.O3, stat.O4)
		stat.Chi2Pr = chiSquarePValue(stat.Chi2, 1)
		results = append(results, stat)
	}

	return results
}
