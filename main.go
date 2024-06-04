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
	conclusion := lib.ConclusionRes(resultsNormal, resultsUniform)
	fmt.Println("Random numbers:")
	rndNumNormals, rndNumUniforms := lib.RndNum(conclusion)
	lib.PrintRndNumTable(conclusion, rndNumNormals, rndNumUniforms)
	fmt.Println("Random weighted net:")
	net := lib.RandWeightenedNetwork(data, conclusion)
	lib.OutputTableNet(net)
	fmt.Println("Distance matrix:")
	distm := lib.DistanceMatrix(net)
	outer, inner := lib.FindOuterAndInnerRadius(distm)
	lib.OutputDistanceMatrix(distm, outer, inner)
	fmt.Println("Results of modelling:")
	minOuter, minInner, sum := lib.ResultModel(outer, inner)
	lib.OutputResultModelTable(minOuter, minInner, sum)
}
