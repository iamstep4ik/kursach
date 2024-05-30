package main

import (
	"fmt"
	"kursach/lib"
)

func main() {
	data := lib.Parse("test.json")
	e := lib.CalculateE(data)
	sd := lib.CalculateStandardDeviation(data, e)
	diff := lib.CalculateDifferenceESD(e, sd)
	sum := lib.CalculateSumESD(e, sd)
	// fmt.Println(e)
	fmt.Printf("data: %v\n", data)
	fmt.Println()
	fmt.Printf("E: %v\n", e)
	fmt.Println()
	fmt.Printf("SD: %v\n", sd)
	fmt.Println()
	fmt.Printf("Difference: %v\n", diff)
	fmt.Println()
	fmt.Printf("Sum: %v\n", lib.CalculateSumESD(e, sd))
	fmt.Println()
	fmt.Printf("O1: %v\n", lib.CalculateO1(data, diff))
	fmt.Println()
	fmt.Printf("O2: %v\n", lib.CalculateO2(data, diff, e))
	fmt.Println()
	fmt.Printf("O3: %v\n", lib.CalculateO3(data, sum, e))
	fmt.Println()
	fmt.Printf("O4: %v\n", lib.CalculateO4(data, sum))

}
