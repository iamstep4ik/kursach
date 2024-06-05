package lib

import (
	"fmt"
	"math"
	"os"
	"sort"
	"text/tabwriter"
)

type PathStatistics struct {
	Path       string
	Statistics []NormalStatistics
}

func outputTableNormal(results []NormalStatistics) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	defer w.Flush()

	fmt.Fprintln(w, "Path\tMu\tSigma\tMu-Sigma\tMu+Sigma\tO1\tO2\tO3\tO4\tE1\tE2\tE3\tE4\tChi2\tChi2Pr\t")

	for _, result := range results {
		fmt.Fprintf(w, "%s\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t\n",
			result.Path,
			result.Mu,
			result.Sigma,
			result.MuSigma,
			result.MuPlusSigma,
			result.O1,
			result.O2,
			result.O3,
			result.O4,
			result.E1,
			result.E2,
			result.E3,
			result.E4,
			result.Chi2,
			result.Chi2Pr)
	}
}

func PrintSortedResultsNormal(results []NormalStatistics) {
	sort.Slice(results, func(i, j int) bool {
		return results[i].Path < results[j].Path
	})

	outputTableNormal(results)
}

func outputTableUniform(results []UniformStatistics) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	defer w.Flush()

	fmt.Fprintln(w, "Path\tMu\tSigma\tA\tB\tO1\tO2\tO3\tO4\tE1\tE2\tE3\tE4\tChi2\tChi2Pr\t")

	for _, result := range results {
		fmt.Fprintf(w, "%s\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t\n",
			result.Path,
			result.Mu,
			result.Sigma,
			result.A,
			result.B,
			result.O1,
			result.O2,
			result.O3,
			result.O4,
			result.E1,
			result.E2,
			result.E3,
			result.E4,
			result.Chi2,
			result.Chi2Pr)
	}
}

func PrintSortedResultsUniform(results []UniformStatistics) {
	sort.Slice(results, func(i, j int) bool {
		return results[i].Path < results[j].Path
	})

	outputTableUniform(results)
}

func PrintRndNumTable(conclusions []Conclusion, rndNumNormals []RndNumNormal, rndNumUniforms []RndNumUniform) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	// Print header
	fmt.Fprintln(w, "Path\tType\tMu (Normal)\tSigma (Normal)\tZ (Normal)\tA (Uniform)\tB (Uniform)\t")

	normalMap := make(map[string]RndNumNormal)
	uniformMap := make(map[string]RndNumUniform)

	// Store RndNumNormal structs in a map for easy access
	for _, rndNum := range rndNumNormals {
		normalMap[rndNum.Path] = rndNum
	}

	// Store RndNumUniform structs in a map for easy access
	for _, rndNum := range rndNumUniforms {
		uniformMap[rndNum.Path] = rndNum
	}

	// Iterate over conclusions and print RndNum results
	for _, conclusion := range conclusions {
		// Print path and distribution type
		fmt.Fprintf(w, "%s\t%s\t", conclusion.Path, conclusion.Type)

		// Print normal distribution values if available
		if conclusion.Type == "Нормальное" {
			rndNumStruct, ok := normalMap[conclusion.Path]
			if ok {
				fmt.Fprintf(w, "%.4f\t%.4f\t%.4f\t-\t-\t\n", rndNumStruct.Mu, rndNumStruct.Sigma, rndNumStruct.Z)
			} else {
				fmt.Fprintln(w, "-\t-\t-\t-\t-")
			}
		} else if conclusion.Type == "Равномерное" {
			// Print uniform distribution values if available
			rndNumStruct, ok := uniformMap[conclusion.Path]
			if ok {
				fmt.Fprintf(w, "-\t-\t-\t%.4f\t%.4f\n", rndNumStruct.A, rndNumStruct.B)
			} else {
				fmt.Fprintln(w, "-\t-\t-\t-\t-")
			}
		}
	}

	w.Flush()
}

func OutputTableNet(matrix [][]float64) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	// Print header
	fmt.Fprintln(w, "Path\t1\t2\t3\t4\t5\t6\t7\t")

	// Iterate over each row in the matrix
	for i, row := range matrix {
		fmt.Fprintf(w, "%d\t", i+1)

		// Print the values in the row
		for _, value := range row {
			if math.IsInf(value, 1) {
				fmt.Fprintf(w, "0.0000\t")
			} else {
				fmt.Fprintf(w, "%.4f\t", value)
			}
		}

		fmt.Fprintln(w)
	}

	w.Flush()
}

func OutputDistanceMatrix(distances [][]float64, outer, inner []float64) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	defer w.Flush()

	// Print header row
	fmt.Fprintf(w, "Path\t")
	for i := range distances {
		fmt.Fprintf(w, "%d\t", i+1)
	}
	fmt.Fprintf(w, "Outer\n")

	// Print distance matrix and outer values
	for i, row := range distances {
		fmt.Fprintf(w, "%d\t", i+1)
		for _, d := range row {
			if math.IsInf(d, 1) {
				fmt.Fprintf(w, "0.0000\t")
			} else {
				fmt.Fprintf(w, "%.4f\t", d)
			}
		}
		fmt.Fprintf(w, "%.4f\n", outer[i])
	}

	// Print inner values
	fmt.Fprintf(w, "Inner\t")
	for _, v := range inner {
		fmt.Fprintf(w, "%.4f\t", v)
	}
	fmt.Fprintf(w, "\n")
}

func OutputResultTable(results []ResultModel) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	defer w.Flush()

	// Print header row
	fmt.Fprintf(w, "Node\tOuter Radius\tInner Radius\tSum\n")

	// Print each result
	for _, result := range results {
		fmt.Fprintf(w, "%d\t%.4f\t%.4f\t%.4f\n",
			result.Index, result.OuterRadius, result.InnerRadius, result.Sum)
	}
}

func OutputMinSumModel(minModel ResultModel) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	defer w.Flush()

	// Print header row
	fmt.Fprintf(w, "Node\tOuter Radius\tInner Radius\tSum\n")

	// Print the minModel
	fmt.Fprintf(w, "%d\t%.4f\t%.4f\t%.4f\n",
		minModel.Index, minModel.OuterRadius, minModel.InnerRadius, minModel.Sum)
}
