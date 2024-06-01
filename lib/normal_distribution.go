package lib

import (
	"math"
)

const (
	E1 = 0.16 * float64(25)
	E2 = 0.34 * float64(25)
	E3 = E2
	E4 = E1
)

func InitializeMapNormal() []map[string]float64 {
	var keys = []string{"mu", "sigma", "mu-sigma", "mu+sigma", "o1", "o2", "o3", "o4", "e1", "e2", "e3", "e4", "chi2", "chi2pr"}
	m := make([]map[string]float64, 9)
	for i := range m {
		m[i] = make(map[string]float64)
		for _, key := range keys {
			m[i][key] = 0.0
		}
		m[i]["e1"] = E1
		m[i]["e2"] = E2
		m[i]["e3"] = E3
		m[i]["e4"] = E4
	}
	return m
}

func CalculateE(data map[string][]int, m []map[string]float64) {
	k := 0
	for _, arr := range data {
		sum := 0
		for _, value := range arr {
			sum += value
		}
		m[k]["mu"] = float64(sum) / float64(len(arr))
		k++
	}
}

func CalculateStandardDeviation(data map[string][]int, m []map[string]float64) {
	k := 0
	n := 25
	for _, arr := range data {
		E := m[k]["mu"]
		sum := 0.0
		for _, value := range arr {
			sum += math.Pow(float64(value)-E, 2)
		}
		m[k]["sigma"] = math.Sqrt(sum / float64(n-1))
		k++
	}
}

func CalculateDifferenceAndSumESD(m []map[string]float64) {
	for i := range m {
		m[i]["mu-sigma"] = m[i]["mu"] - m[i]["sigma"]
		m[i]["mu+sigma"] = m[i]["mu"] + m[i]["sigma"]
	}
}

func CalculateO(data map[string][]int, m []map[string]float64) {
	k := 0
	for _, arr := range data {
		o1Counter, o2Counter, o3Counter, o4Counter := 0, 0, 0, 0
		muMinusSigma := m[k]["mu-sigma"]
		mu := m[k]["mu"]
		muPlusSigma := m[k]["mu+sigma"]

		for _, value := range arr {
			if float64(value) <= muMinusSigma {
				o1Counter++
			} else if float64(value) > muMinusSigma && float64(value) <= mu {
				o2Counter++
			} else if float64(value) > mu && float64(value) <= muPlusSigma {
				o3Counter++
			} else if float64(value) > muPlusSigma {
				o4Counter++
			}
		}

		m[k]["o1"] = float64(o1Counter)
		m[k]["o2"] = float64(o2Counter)
		m[k]["o3"] = float64(o3Counter)
		m[k]["o4"] = float64(o4Counter)
		k++
	}
}

func CalculateChi2(m []map[string]float64) {
	for i := range m {
		o1, o2, o3, o4 := m[i]["o1"], m[i]["o2"], m[i]["o3"], m[i]["o4"]
		m[i]["chi2"] = math.Pow(o1-E1, 2)/E1 + math.Pow(o2-E2, 2)/E2 + math.Pow(o3-E3, 2)/E3 + math.Pow(o4-E4, 2)/E4
	}
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
	const epsilon = 1e-4
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

func CalculateChiSquarePValueForMap(m []map[string]float64) {
	for i := range m {
		m[i]["chi2pr"] = chiSquarePValue(m[i]["chi2"], 1)
	}
}

func CalulateAll(data map[string][]int, m []map[string]float64) {
	CalculateE(data, m)
	CalculateStandardDeviation(data, m)
	CalculateDifferenceAndSumESD(m)
	CalculateO(data, m)
	CalculateChi2(m)
	CalculateChiSquarePValueForMap(m)
}
