package lib

import (
	"fmt"
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

func PrintRndNumTable(conclusion []string, rndNumNormals []RndNumNormal, rndNumUniforms []RndNumUniform) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	defer w.Flush()

	// Print header
	fmt.Fprintln(w, "Path\tMu (Normal)\tSigma (Normal)\tZ (Normal)\tA (Uniform)\tB (Uniform)\t")

	normalIndex, uniformIndex := 0, 0

	// Iterate over conclusion and print RndNum results
	for i, status := range conclusion {
		fmt.Fprintf(w, "%d\t", i)

		if status == "Нормальное" {
			if normalIndex < len(rndNumNormals) {
				fmt.Fprintf(w, "%f\t%f\t%f\t-\t-\t\n",
					rndNumNormals[normalIndex].Mu,
					rndNumNormals[normalIndex].Sigma,
					rndNumNormals[normalIndex].Z,
				)
				normalIndex++
			} else {
				fmt.Fprintln(w, "-\t-\t-\t-\t-\t")
			}
		} else if status == "Равномерное" {
			if uniformIndex < len(rndNumUniforms) {
				fmt.Fprintf(w, "-\t-\t-\t%f\t%f\t\n",
					rndNumUniforms[uniformIndex].A,
					rndNumUniforms[uniformIndex].B,
				)
				uniformIndex++
			} else {
				fmt.Fprintln(w, "-\t-\t-\t-\t-\t")
			}
		}
	}
}
