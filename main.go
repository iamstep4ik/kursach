package main

import (
	"fmt"
	"kursach/lib"
)

func main() {
	data := lib.Parse("test.json")
	m := lib.InitializeMapNormal()
	lib.CalulateAll(data, m)
	fmt.Println("Normal Distribution")
	lib.OutputTableNormal(m)
	Um := lib.InitializeMapUniform(m)
	lib.CalculateUniAll(data, Um)
	fmt.Println("Uniform Distribution")
	lib.OutputTableUniform(Um)

}
