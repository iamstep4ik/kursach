package lib

type Conclusion struct {
	Path  string
	Type  string
	Mu    float64
	Sigma float64
	A     float64
	B     float64
}

func ConclusionRes(resN []NormalStatistics, resU []UniformStatistics) []Conclusion {
	conclusions := make([]Conclusion, len(resN))
	for i := range resN {
		if resN[i].Chi2Pr > resU[i].Chi2Pr {
			conclusions[i] = Conclusion{
				Path:  resN[i].Path,
				Type:  "Нормальное",
				Mu:    resN[i].Mu,
				Sigma: resN[i].Sigma,
			}
		} else {
			conclusions[i] = Conclusion{
				Path: resU[i].Path,
				Type: "Равномерное",
				A:    resU[i].A,
				B:    resU[i].B,
			}
		}
	}
	return conclusions
}
