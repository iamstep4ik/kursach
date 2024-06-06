package lib

import (
	"fmt"
	"html"
	"log"
	"math"
	"sort"
)

type PathStatistics struct {
	Path       string
	Statistics []NormalStatistics
}

func PrintTableHTML(data Data) string {
	// Extract and sort keys
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Generate HTML table
	htmlContent := "<h2>Исходные данные</h2>"
	htmlContent += "<table border='1'>"
	htmlContent += "<tr><th>Path</th>"
	for i := 0; i < 5; i++ {
		htmlContent += "<th>9:00</th><th>12:00</th><th>15:00</th><th>18:00</th><th>21:00</th>"
	}
	htmlContent += "</tr>"

	// Write each sorted key and its values
	for _, key := range keys {
		values := data[key]
		if len(values) != 25 {
			log.Fatalf("Error: expected 25 values for key %s, got %d", key, len(values))
		}
		htmlContent += fmt.Sprintf("<tr><td>%s</td>", html.EscapeString(key))
		for i := 0; i < 25; i += 5 {
			htmlContent += fmt.Sprintf("<td>%d</td><td>%d</td><td>%d</td><td>%d</td><td>%d</td>", values[i], values[i+1], values[i+2], values[i+3], values[i+4])
		}
		htmlContent += "</tr>"
	}

	htmlContent += "</table>"
	return htmlContent
}

func outputTableNormal(results []NormalStatistics) string {
	htmlContent := "<h2>Нормальное Распределение</h2>"
	htmlContent += "<table border='1'>"
	htmlContent += "<tr><th>Path</th><th>Mu</th><th>Sigma</th><th>Mu-Sigma</th><th>Mu+Sigma</th><th>O1</th><th>O2</th><th>O3</th><th>O4</th><th>E1</th><th>E2</th><th>E3</th><th>E4</th><th>Chi2</th><th>Chi2Pr</th></tr>"

	for _, result := range results {
		htmlContent += fmt.Sprintf(
			"<tr><td>%s</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td></tr>",
			html.EscapeString(result.Path), result.Mu, result.Sigma, result.MuSigma, result.MuPlusSigma,
			result.O1, result.O2, result.O3, result.O4, result.E1, result.E2, result.E3, result.E4, result.Chi2, result.Chi2Pr)
	}

	htmlContent += "</table>"
	return htmlContent
}

func PrintSortedResultsNormal(results []NormalStatistics) string {
	sort.Slice(results, func(i, j int) bool {
		return results[i].Path < results[j].Path
	})

	return outputTableNormal(results)
}

func outputTableUniform(results []UniformStatistics) string {
	htmlContent := "<h2>Равномерное Распределение</h2>"
	htmlContent += "<table border='1'>"
	htmlContent += "<tr><th>Path</th><th>Mu</th><th>Sigma</th><th>A</th><th>B</th><th>O1</th><th>O2</th><th>O3</th><th>O4</th><th>E1</th><th>E2</th><th>E3</th><th>E4</th><th>Chi2</th><th>Chi2Pr</th></tr>"

	for _, result := range results {
		htmlContent += fmt.Sprintf(
			"<tr><td>%s</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td><td>%.4f</td></tr>",
			html.EscapeString(result.Path), result.Mu, result.Sigma, result.A, result.B,
			result.O1, result.O2, result.O3, result.O4, result.E1, result.E2, result.E3, result.E4, result.Chi2, result.Chi2Pr)
	}

	htmlContent += "</table>"
	return htmlContent
}

func PrintSortedResultsUniform(results []UniformStatistics) string {
	sort.Slice(results, func(i, j int) bool {
		return results[i].Path < results[j].Path
	})

	return outputTableUniform(results)
}

func PrintRndNumTable(conclusions []Conclusion, rndNumNormals []RndNumNormal, rndNumUniforms []RndNumUniform) string {
	htmlContent := "<h2>Генерация случайного числа</h2>"
	htmlContent += "<table border='1'>"
	htmlContent += "<tr><th>Path</th><th>Type</th><th>Mu (Normal)</th><th>Sigma (Normal)</th><th>Z (Normal)</th><th>A (Uniform)</th><th>B (Uniform)</th></tr>"

	// Create maps for normal and uniform random numbers
	normalMap := make(map[string]RndNumNormal)
	uniformMap := make(map[string]RndNumUniform)

	for _, rndNum := range rndNumNormals {
		normalMap[rndNum.Path] = rndNum
	}

	for _, rndNum := range rndNumUniforms {
		uniformMap[rndNum.Path] = rndNum
	}

	// Sort conclusions by path
	sort.Slice(conclusions, func(i, j int) bool {
		return conclusions[i].Path < conclusions[j].Path
	})

	// Iterate through sorted conclusions and generate table rows
	for _, conclusion := range conclusions {
		htmlContent += fmt.Sprintf("<tr><td>%s</td><td>%s</td>", html.EscapeString(conclusion.Path), conclusion.Type)

		if conclusion.Type == "Нормальное" {
			if rndNum, ok := normalMap[conclusion.Path]; ok {
				htmlContent += fmt.Sprintf("<td>%.4f</td><td>%.4f</td><td>%.4f</td><td>-</td><td>-</td></tr>",
					rndNum.Mu, rndNum.Sigma, rndNum.Z)
			} else {
				htmlContent += "<td>-</td><td>-</td><td>-</td><td>-</td><td>-</td></tr>"
			}
		} else if conclusion.Type == "Равномерное" {
			if rndNum, ok := uniformMap[conclusion.Path]; ok {
				htmlContent += fmt.Sprintf("<td>-</td><td>-</td><td>-</td><td>%.4f</td><td>%.4f</td></tr>",
					rndNum.A, rndNum.B)
			} else {
				htmlContent += "<td>-</td><td>-</td><td>-</td><td>-</td><td>-</td></tr>"
			}
		}
	}

	htmlContent += "</table>"
	return htmlContent
}
func OutputTableNet(matrix [][]float64) string {
	htmlContent := "<h2>Случайная взвешенная сеть</h2>"
	htmlContent += "<table border='1'>"
	htmlContent += "<tr><th>Path</th><th>1</th><th>2</th><th>3</th><th>4</th><th>5</th><th>6</th><th>7</th></tr>"

	for i, row := range matrix {
		htmlContent += fmt.Sprintf("<tr><td>%d</td>", i+1)

		for _, value := range row {
			if math.IsInf(value, 1) {
				htmlContent += "<td>0.0000</td>"
			} else {
				htmlContent += fmt.Sprintf("<td>%.4f</td>", value)
			}
		}

		htmlContent += "</tr>"
	}

	htmlContent += "</table>"
	return htmlContent
}

func OutputDistanceMatrix(distances [][]float64, outer, inner []float64) string {
	htmlContent := "<h2>Матрица Расстояний</h2>"
	htmlContent += "<table border='1'>"
	htmlContent += "<tr><th>Path</th>"
	for i := range distances {
		htmlContent += fmt.Sprintf("<th>%d</th>", i+1)
	}
	htmlContent += "<th>Outer</th></tr>"

	for i, row := range distances {
		htmlContent += fmt.Sprintf("<tr><td>%d</td>", i+1)
		for _, d := range row {
			if math.IsInf(d, 1) {
				htmlContent += "<td>0.0000</td>"
			} else {
				htmlContent += fmt.Sprintf("<td>%.4f</td>", d)
			}
		}
		htmlContent += fmt.Sprintf("<td>%.4f</td></tr>", outer[i])
	}

	htmlContent += "<tr><td>Inner</td>"
	for _, v := range inner {
		htmlContent += fmt.Sprintf("<td>%.4f</td>", v)
	}
	htmlContent += "</tr></table>"
	return htmlContent
}

func OutputResultTable(results []ResultModel) string {
	htmlContent := "<h2>Результирующая таблица</h2>"
	htmlContent += "<table border='1'>"
	htmlContent += "<tr><th>Node</th><th>Outer Radius</th><th>Inner Radius</th><th>Sum</th></tr>"

	for _, result := range results {
		htmlContent += fmt.Sprintf("<tr><td>%d</td><td>%.4f</td><td>%.4f</td><td>%.4f</td></tr>",
			result.Index, result.OuterRadius, result.InnerRadius, result.Sum)
	}

	htmlContent += "</table>"
	return htmlContent
}

func OutputMinSumModel(minModel ResultModel) string {
	htmlContent := "<h2>Результат моделирования</h2>"
	htmlContent += "<table border='1'>"
	htmlContent += "<tr><th>Node</th><th>Outer Radius</th><th>Inner Radius</th><th>Sum</th></tr>"

	htmlContent += fmt.Sprintf("<tr><td>%d</td><td>%.4f</td><td>%.4f</td><td>%.4f</td></tr>",
		minModel.Index, minModel.OuterRadius, minModel.InnerRadius, minModel.Sum)

	htmlContent += "</table>"
	return htmlContent
}
