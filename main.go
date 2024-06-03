package main

import (
	"fmt"
	"kursach/lib"
)

func main() {
	data := lib.Parse("data.json")
	resultsNormal := lib.CalculateAllNormal(data)
	fmt.Println("Normal Distribution:")
	lib.PrintSortedResultsNormal(resultsNormal)

	resultsUniform := lib.CalculateUniAll(data)
	fmt.Println("Uniform Distribution:")
	lib.PrintSortedResultsUniform(resultsUniform)
	fmt.Println(lib.Conclusion(resultsNormal, resultsUniform))
}
