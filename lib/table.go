package lib

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func OutputTableNormal(m []map[string]float64) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "Index\tmu\tsigma\tmu-sigma\tmu+sigma\to1\to2\to3\to4\te1\te2\te3\te4\tchi2\tchi2pr\t")

	for i, mapInstance := range m {
		fmt.Fprintf(w, "%d\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t\n",
			i,
			mapInstance["mu"],
			mapInstance["sigma"],
			mapInstance["mu-sigma"],
			mapInstance["mu+sigma"],
			mapInstance["o1"],
			mapInstance["o2"],
			mapInstance["o3"],
			mapInstance["o4"],
			mapInstance["e1"],
			mapInstance["e2"],
			mapInstance["e3"],
			mapInstance["e4"],
			mapInstance["chi2"],
			mapInstance["chi2pr"])
	}

	w.Flush()
}

func OutputTableUniform(m []map[string]float64) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "Index\tmu\tsigma\ta\tb\to1\to2\to3\to4\te1\te2\te3\te4\tchi2\tchi2pr\t")

	for i, mapInstance := range m {
		fmt.Fprintf(w, "%d\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t\n",
			i,
			mapInstance["mu"],
			mapInstance["sigma"],
			mapInstance["a"],
			mapInstance["b"],
			mapInstance["o1"],
			mapInstance["o2"],
			mapInstance["o3"],
			mapInstance["o4"],
			mapInstance["e1"],
			mapInstance["e2"],
			mapInstance["e3"],
			mapInstance["e4"],
			mapInstance["chi2"],
			mapInstance["chi2pr"])
	}

	w.Flush()
}
