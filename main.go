package main

import (
	"fmt"
	"kursach/lib"
	"net/http"
)

func main() {
	// Parse data and perform calculations
	data := lib.Parse("data.json")
	resultsNormal := lib.CalculateAllNormal(data)
	resultsUniform := lib.CalculateUniAll(data)
	conclusion := lib.ConclusionRes(resultsNormal, resultsUniform)
	dataset := lib.FormData()
	fmt.Printf("Dataset: %v\n", dataset)
	histogramFile, err := lib.CreateHistogram(dataset)
	if err != nil {
		fmt.Printf("Failed to create histogram: %v\n", err)
		return
	}

	// HTTP handler to serve the HTML content
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rndNumNormals, rndNumUniforms := lib.RndNum(conclusion)
		net := lib.RandWeightenedNetwork(data, conclusion)
		distm := lib.DistanceMatrix(net)
		outer, inner := lib.FindOuterAndInnerRadius(distm)
		res := lib.CalculateResultModel(outer, inner)
		min := lib.FindMinSumModel(res)

		htmlContent := "<html><head><style>"
		htmlContent += `
			table {
				width: 100%;
				border-collapse: collapse;
			}
			th, td {
				padding: 8px;
				text-align: left;
				border: 1px solid #ddd;
			}
			tr:nth-child(even) {
				background-color: #f2f2f2;
			}
			th {
				background-color: #4CAF50;
				color: white;
			}
		`
		htmlContent += "</style></head><body>"
		htmlContent += "<h2>Карта</h2>"
		htmlContent += fmt.Sprint("<img src='pics/map.png' alt='Map'/>")
		htmlContent += "<h2>Граф</h2>"
		htmlContent += fmt.Sprint("<img src='pics/graph.png' alt='Graph'/>")
		htmlContent += lib.PrintTableHTML(data)
		htmlContent += "<br><br>"
		htmlContent += lib.PrintSortedResultsNormal(resultsNormal)
		htmlContent += "<br><br>"
		htmlContent += lib.PrintSortedResultsUniform(resultsUniform)
		htmlContent += "<br><br>"
		htmlContent += lib.PrintRndNumTable(conclusion, rndNumNormals, rndNumUniforms)
		htmlContent += "<br><br>"
		htmlContent += lib.OutputTableNet(net)
		htmlContent += "<br><br>"
		htmlContent += lib.OutputDistanceMatrix(distm, outer, inner)
		htmlContent += "<br><br>"
		htmlContent += lib.OutputResultTable(res)
		htmlContent += "<br><br>"
		htmlContent += lib.OutputMinSumModel(min)
		htmlContent += "<br><br>"
		htmlContent += "<h2>Гистограмма</h2>"
		htmlContent += fmt.Sprintf("<img src='%s' alt='Histogram'/>", histogramFile)
		htmlContent += "</body></html>"

		fmt.Fprintf(w, "%s", htmlContent)
	})

	http.Handle("/pics/", http.StripPrefix("/pics/", http.FileServer(http.Dir("pics"))))

	// Start the server
	fmt.Println("Serving on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
