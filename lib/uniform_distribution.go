package lib

import (
	"math"
)

const (
	uE1 = 0.01 * float64(25)
	uE2 = 0.49 * float64(25)
	uE3 = uE2
	uE4 = uE1
)

func InitializeMapUniform(m []map[string]float64) []map[string]float64 {
	var keys = []string{"mu", "sigma", "a", "b", "o1", "o2", "o3", "o4", "e1", "e2", "e3", "e4", "chi2", "chi2pr"}
	Um := make([]map[string]float64, 9)
	for i := range Um {
		Um[i] = make(map[string]float64)
		for _, key := range keys {
			Um[i][key] = 0.0
			Um[i]["e1"] = uE1
			Um[i]["e2"] = uE2
			Um[i]["e3"] = uE3
			Um[i]["e4"] = uE4
			Um[i]["mu"] = m[i]["mu"]
			Um[i]["sigma"] = m[i]["sigma"]
		}
	}
	return Um
}

func CalculateAB(m []map[string]float64) {
	for i := range m {
		m[i]["a"] = m[i]["mu"] - (math.Sqrt(3) * m[i]["sigma"])
		m[i]["b"] = m[i]["mu"] - (math.Sqrt(3) * m[i]["a"])
	}
}

func CalculateUniO(data map[string][]int, m []map[string]float64) {
	k := 0
	for _, arr := range data {
		counter1 := 0
		counter2 := 0
		counter3 := 0
		counter4 := 0
		for _, value := range arr {
			if float64(value) <= m[k]["a"] {
				counter1++
			} else if float64(value) > m[k]["a"] && float64(value) <= m[k]["mu"] {
				counter2++
			} else if float64(value) < m[k]["b"] && float64(value) >= m[k]["mu"] {
				counter3++
			} else if float64(value) > m[k]["b"] {
				counter4++
			}
		}
		m[k]["o1"] = float64(counter1)
		m[k]["o2"] = float64(counter2)
		m[k]["o3"] = float64(counter3)
		m[k]["o4"] = float64(counter4)
		k++
	}
}

func CalculateUniChi2(m []map[string]float64) {
	for i := range m {
		o1, o2, o3, o4 := m[i]["o1"], m[i]["o2"], m[i]["o3"], m[i]["o4"]
		m[i]["chi2"] = math.Pow(o1-uE1, 2)/uE1 + math.Pow(o2-uE2, 2)/uE2 + math.Pow(o3-uE3, 2)/uE3 + math.Pow(o4-uE4, 2)/uE4
	}
}

func CalculateUniAll(data map[string][]int, m []map[string]float64) {
	CalculateAB(m)
	CalculateUniO(data, m)
	CalculateUniChi2(m)
	CalculateChiSquarePValueForMap(m)
}
