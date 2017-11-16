package data

import (
	"log"
)

type Lattice struct {
	ID     int
	Key    []int
	Values []Point
}

type Point struct {
	Value int
	Snap  []int
}

func (l *Lattice) Merge(other Lattice) {
	// create a delta array from key

	for index := range other.Key {
		_, ok := l.Key[index] // figure out index thing, this isn't a map derrr
		if !ok {
			l.Key = make([]int, index+1)
			continue
		}
		// deltaArray[index] = difference from key and other
	}

	// pull the values from other for the amount of delta on each key

	return
}

func (l *Lattice) Create() (newLattice *Lattice) {
	l.Key = make([]int, len(l.Key)+1)
	newLattice = &Lattice{
		ID:     len(l.Key) - 1,
		Key:    make([]int, len(l.Key)),
		Values: l.Values,
	}
	return
}

func (l *Lattice) Add(value int) {
	l.Key[l.ID] = l.Key[l.ID] + 1

	newPoint := Point{
		Value: value,
		Snap:  l.Key,
	}
	l.Values = append(l.Values, newPoint)
	return
}

func (l *Lattice) Print() {
	var total int
	for _, value := range l.Values {
		total = total + value.Value
	}
	log.Printf(`
		ID: %v
		Length of Values: %v
		Snap of Keys: %v
		Total: %v`,
		l.ID,
		len(l.Values),
		l.Key,
		total,
	)
}
