package main

import "github.com/Siryu/tasty/internal/data"

func main() {
	// wierd hash to get count of how many nodes

	// but for now...
	mainLattice := &data.Lattice{
		ID:     0,
		Key:    []int{0},
		Values: []data.Point{},
	}

	secondLattice := mainLattice.Create()
	thirdLattice := mainLattice.Create()
	secondLattice.Add(5)
	thirdLattice.Add(10)
	secondLattice.Add(10)
	mainLattice.Add(1)
	thirdLattice.Add(55)

	mainLattice.Print()
	secondLattice.Print()
	thirdLattice.Print()
}
