package lib

import (
	"math"
	"math/rand"
	"time"
)

type RndNumNormal struct {
	Path  string
	Mu    float64
	Sigma float64
	Z     float64
}

type RndNumUniform struct {
	Path string
	A    float64
	B    float64
}

func GenZ() float64 {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float64()
	return math.Sqrt(-2*math.Log(r)) * math.Cos(2*math.Pi*r)
}

func RndNum(conclusions []Conclusion) ([]RndNumNormal, []RndNumUniform) {
	rndNumNormals := make([]RndNumNormal, 0)
	rndNumUniforms := make([]RndNumUniform, 0)
	for _, conclusion := range conclusions {
		if conclusion.Type == "Нормальное" {
			z := GenZ()
			rndNumNormals = append(rndNumNormals, RndNumNormal{
				Path:  conclusion.Path,
				Mu:    conclusion.Mu,
				Sigma: conclusion.Sigma,
				Z:     z,
			})
		} else if conclusion.Type == "Равномерное" {
			rndNumUniforms = append(rndNumUniforms, RndNumUniform{
				Path: conclusion.Path,
				A:    conclusion.A,
				B:    conclusion.B,
			})
		}
	}
	return rndNumNormals, rndNumUniforms
}
