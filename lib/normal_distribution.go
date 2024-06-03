package lib

import (
	"math"
)

const (
	E1 = 0.16 * 25.0
	E2 = 0.34 * 25.0
	E3 = 0.34 * 25.0
	E4 = 0.16 * 25.0
)

type NormalStatistics struct {
	Path        string
	Mu          float64
	Sigma       float64
	MuSigma     float64
	MuPlusSigma float64
	O1          float64
	O2          float64
	O3          float64
	O4          float64
	E1          float64
	E2          float64
	E3          float64
	E4          float64
	Chi2        float64
	Chi2Pr      float64
}

func InitializeStatistics() NormalStatistics {
	return NormalStatistics{
		E1: E1,
		E2: E2,
		E3: E3,
		E4: E4,
	}
}

func CalculateMu(data []int) float64 {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return float64(sum) / float64(len(data))
}

func CalculateSigma(data []int, mu float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += math.Pow(float64(value)-mu, 2)
	}
	return math.Sqrt(sum / float64(len(data)-1))
}

func CalculateO(data []int, muMinusSigma, mu, muPlusSigma float64) (o1, o2, o3, o4 float64) {
	for _, value := range data {
		switch {
		case float64(value) <= muMinusSigma:
			o1++
		case float64(value) > muMinusSigma && float64(value) <= mu:
			o2++
		case float64(value) > mu && float64(value) <= muPlusSigma:
			o3++
		case float64(value) > muPlusSigma:
			o4++
		}
	}
	return o1, o2, o3, o4
}

func CalculateChi2(o1, o2, o3, o4 float64) float64 {
	return (math.Pow(o1-E1, 2)/E1 +
		math.Pow(o2-E2, 2)/E2 +
		math.Pow(o3-E3, 2)/E3 +
		math.Pow(o4-E4, 2)/E4)
}

const PI = math.Pi

func gammaFunction(z float64) float64 {
	if z < 0.5 {
		return PI / (math.Sin(PI*z) * gammaFunction(1-z))
	}

	z -= 1
	x := 0.99999999999980993
	coefficients := []float64{
		676.5203681218851, -1259.1392167224028, 771.32342877765313,
		-176.61502916214059, 12.507343278686905, -0.13857109526572012,
		9.9843695780195716e-6, 1.5056327351493116e-7,
	}

	for i, coef := range coefficients {
		x += coef / (z + float64(i) + 1)
	}

	t := z + float64(len(coefficients)) - 0.5
	return math.Sqrt(2*PI) * math.Pow(t, z+0.5) * math.Exp(-t) * x
}

func integrand(t, s float64) float64 {
	return math.Pow(t, s-1) * math.Exp(-t)
}

func incompleteGamma(s, x float64) float64 {
	const epsilon = 1e-5
	result := 0.0
	step := epsilon
	for t := 0.0; t <= x; t += step {
		num := integrand(t, s) * step
		if math.IsInf(num, 0) {
			num = 0
		}
		result += num
	}
	return result
}

func chiSquareCDF(x, k float64) float64 {
	s := k / 2.0
	gammaS := gammaFunction(s)
	gammaIncomplete := incompleteGamma(s, x/2.0)
	return gammaIncomplete / gammaS
}

func chiSquarePValue(x, k float64) float64 {
	cdf := chiSquareCDF(x, k)
	return 1 - cdf
}

func CalculateAllNormal(data map[string][]int) []NormalStatistics {
	results := make([]NormalStatistics, len(data))
	i := 0

	for key, values := range data {
		stat := InitializeStatistics()
		stat.Path = key
		stat.Mu = CalculateMu(values)
		stat.Sigma = CalculateSigma(values, stat.Mu)
		stat.MuSigma = stat.Mu - stat.Sigma
		stat.MuPlusSigma = stat.Mu + stat.Sigma

		stat.O1, stat.O2, stat.O3, stat.O4 = CalculateO(values, stat.MuSigma, stat.Mu, stat.MuPlusSigma)

		stat.Chi2 = CalculateChi2(stat.O1, stat.O2, stat.O3, stat.O4)
		stat.Chi2Pr = chiSquarePValue(stat.Chi2, 1)

		results[i] = stat
		i++
	}

	return results
}
