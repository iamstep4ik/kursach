package lib

import (
	"math"
	"math/rand"
	"time"
)

type RndNumNormal struct {
	Mu    float64
	Sigma float64
	Z     float64
}
type RndNumUniform struct {
	A float64
	B float64
}

func GenZ() float64 {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float64()
	result := math.Sqrt(-2*math.Log(r)) * math.Cos(2*math.Pi*r)

	return result
}

func RndNum(conclusion []string, resultsNormal []NormalStatistics, resultsUniform []UniformStatistics) ([]RndNumNormal, []RndNumUniform) {
	rndNumNormals := make([]RndNumNormal, 0)
	rndNumUniforms := make([]RndNumUniform, 0)
	for i := 0; i < len(conclusion); i++ {
		if conclusion[i] == "Нормальное" {
			z := GenZ()

			rndNumNormals = append(rndNumNormals, RndNumNormal{
				Mu:    resultsNormal[i].Mu,
				Sigma: resultsNormal[i].Sigma,
				Z:     z,
			})
		} else if conclusion[i] == "Равномерное" {
			rndNumUniforms = append(rndNumUniforms, RndNumUniform{
				A: resultsUniform[i].A,
				B: resultsUniform[i].B,
			})
		}
	}

	return rndNumNormals, rndNumUniforms
}
