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
	conclusion := lib.Conclusion(resultsNormal, resultsUniform)
	fmt.Printf("Conclusion: %v\n", conclusion)
	rndNumNormals, rndNumUniforms := lib.RndNum(conclusion, resultsNormal, resultsUniform)
	lib.PrintRndNumTable(conclusion, rndNumNormals, rndNumUniforms)
}
