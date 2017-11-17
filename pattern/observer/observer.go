package observer

import "math/rand"

// Interfaces
// Define interface Observer
type objObserver interface {
	update() int
}

// Define interface number
type number interface {
	getNumber() int
}

// Define struct for NumberGenerator
type numGenerator struct {
	observers []objObserver
}

// Define struct RandomNumberGenerator
type randNumberGenerator struct {
	*numGenerator
}
// Define struct DigitObserver
type DigitObserver struct {
	generator number
}


// Methods for numGenerator:
// AddObserver: add observer to  observers
func (ng *numGenerator) AddObserver (observer objObserver) {
	ng.observers = append(ng.observers, observer)

}

func (ng *numGenerator) notifyObservers() []int {
	var result []int
	for _, objObserver := range ng.observers {
		result = append(result, objObserver.update())
	}
	return result
}

// Methods for randNumberGenerator:
func NewRandomNumberGenerator() *randNumberGenerator {
	return &randNumberGenerator{&numGenerator{}}
}

func (rn *randNumberGenerator) getNumber() int {
	return rand.Intn(50)
}

func (rn *randNumberGenerator) Execute() []int {
	return rn.notifyObservers()
}

// Methods for DigitObserver:
// Get number generator
func (do *DigitObserver) update() int {
	return do.generator.getNumber()
}