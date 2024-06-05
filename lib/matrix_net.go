package lib

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type ResultModel struct {
	Index       int
	OuterRadius float64
	InnerRadius float64
	Sum         float64
}

func RandWeightenedNetwork(data Data, conclusions []Conclusion) [][]float64 {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float64()
	matrix := make([][]float64, 7)
	for i := range matrix {
		matrix[i] = make([]float64, 7)
		for j := range matrix[i] {
			if i == j {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = math.Inf(1)
			}
		}
	}

	for i := 1; i <= 7; i++ {
		for j := i + 1; j <= 7; j++ {
			key1 := fmt.Sprintf("%d,%d", i, j)
			key2 := fmt.Sprintf("%d,%d", j, i)

			var found bool

			if _, found = data[key1]; !found {
				if _, found = data[key2]; !found {
					continue
				}
			}

			// Find the corresponding conclusion
			var conclusion Conclusion
			for _, c := range conclusions {
				if c.Path == key1 || c.Path == key2 {
					conclusion = c
					break
				}
			}

			if conclusion.Type == "Нормальное" {
				matrix[i-1][j-1] = conclusion.Mu + conclusion.Sigma*r
				matrix[j-1][i-1] = conclusion.Mu + conclusion.Sigma*r
			} else if conclusion.Type == "Равномерное" {
				a := conclusion.A
				b := conclusion.B
				matrix[i-1][j-1] = a + (b-a)*r
				matrix[j-1][i-1] = a + (b-a)*r
			}
		}
	}

	return matrix
}
func DistanceMatrix(matrix [][]float64) [][]float64 {
	n := len(matrix)
	distances := make([][]float64, n)

	for i := 0; i < n; i++ {
		distances[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				distances[i][j] = 0
			} else if matrix[i][j] != math.Inf(1) {
				distances[i][j] = matrix[i][j]
			} else {
				distances[i][j] = math.Inf(1)
			}
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if distances[i][k]+distances[k][j] < distances[i][j] {
					distances[i][j] = distances[i][k] + distances[k][j]
				}
			}
		}
	}

	return distances
}

func FindOuterAndInnerRadius(matrix [][]float64) ([]float64, []float64) {
	outer := make([]float64, len(matrix))
	inner := make([]float64, len(matrix[0]))
	for i := range outer {
		outer[i] = -math.MaxFloat64
	}
	for i := range inner {
		inner[i] = -math.MaxFloat64
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] > outer[i] {
				outer[i] = matrix[i][j]
			}
			if matrix[i][j] > inner[j] {
				inner[j] = matrix[i][j]
			}
		}
	}
	return outer, inner
}

func CalculateResultModel(outer, inner []float64) []ResultModel {
	res := make([]ResultModel, len(outer))
	for i := range outer {
		res[i].Index = i + 1
		res[i].OuterRadius = outer[i]
		res[i].InnerRadius = inner[i]
		res[i].Sum = outer[i] + inner[i]
	}

	return res
}

func FindMinSumModel(models []ResultModel) ResultModel {
	minModel := models[0]
	for _, model := range models[1:] {
		if model.Sum < minModel.Sum {
			minModel = model
		}
	}
	return minModel
}
