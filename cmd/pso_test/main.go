package main

import (
	"fmt"
	"math"

	"github.com/ChristianTaidi/PSO_TEST/pkg/entity"
)

func main() {
	optFunction := func(values []float64) float64 {
		comp1 := math.Pow((1.5 - values[0] + values[0]*values[1]), 2)
		comp2 := math.Pow(2.25-values[0]+values[0]*math.Pow(values[1], 2), 2)
		comp3 := math.Pow(2.625-values[0]+values[0]*math.Pow(values[1], 3), 2)
		return comp1 + comp2 + comp3
	}
	swarm := entity.NewSwarm(50, 2, optFunction)
	iterations := 200
	fmt.Printf("Starting swarm %+v\n", swarm)
	for i := 0; i < iterations; i++ {
		swarm.Update()
	}
	fmt.Println("Finished swarm")
	bestPos, resultValue := swarm.BestPosResult()
	fmt.Printf("Result %v\n", resultValue)
	fmt.Printf("Best Position %+v\n", bestPos)
}
