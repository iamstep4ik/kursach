package lib

func Conclusion(resN []NormalStatistics, resU []UniformStatistics) []string {
	conclusion := make([]string, len(resN))
	for i := range resN {
		if resN[i].Chi2Pr > resU[i].Chi2Pr {
			conclusion[i] = "Нормальное"
		} else {
			conclusion[i] = "Равномерное"
		}
	}
	return conclusion
}
