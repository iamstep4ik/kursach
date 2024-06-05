package lib

import (
	"fmt"
	"os"

	"github.com/wcharczuk/go-chart"
)

func FormData() []float64 {
	dataset := make([]float64, 7)
	n := 10000
	data := Parse("data.json")
	resultsNormal := CalculateAllNormal(data)
	resultsUniform := CalculateUniAll(data)
	conclusion := ConclusionRes(resultsNormal, resultsUniform)
	for i := 0; i < n; i++ {
		net := RandWeightenedNetwork(data, conclusion)
		distm := DistanceMatrix(net)
		outer, inner := FindOuterAndInnerRadius(distm)
		res := CalculateResultModel(outer, inner)
		min := FindMinSumModel(res)
		switch min.Index {
		case 1:
			dataset[0]++
		case 2:
			dataset[1]++
		case 3:
			dataset[2]++
		case 4:
			dataset[3]++
		case 5:
			dataset[4]++
		case 6:
			dataset[5]++
		case 7:
			dataset[6]++
		}
	}
	return dataset
}

func CreateHistogram(dataset []float64) (string, error) {
	graph := chart.BarChart{
		Title: "Histogram",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
			Ticks: []chart.Tick{
				{Value: 0.0, Label: "0"},
				{Value: 2000.0, Label: "2000"},
				{Value: 4000.0, Label: "4000"},
				{Value: 6000.0, Label: "6000"},
				{Value: 8000.0, Label: "8000"},
				{Value: 10000.0, Label: "10000"},
			},
		},
		Bars: []chart.Value{
			{Value: dataset[0], Label: "1"},
			{Value: dataset[1], Label: "2"},
			{Value: dataset[2], Label: "3"},
			{Value: dataset[3], Label: "4"},
			{Value: dataset[4], Label: "5"},
			{Value: dataset[5], Label: "6"},
			{Value: dataset[6], Label: "7"},
		},
	}

	filePath := "pics/histogram.png"
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	if err := graph.Render(chart.PNG, file); err != nil {
		return "", fmt.Errorf("error rendering graph: %v", err)
	}

	return filePath, nil
}
