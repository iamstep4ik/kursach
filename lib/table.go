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
