package models

type Map struct {
	Node []Node
}

type Node struct {
	Location *Location
	Weight   float64
}

type Location struct {
	Lat float64
	Lon float64
}

type Sector struct {
	Border *[]Node
	House  []House
}

type House struct {
	Sector   *Sector
	Location *Location
}
