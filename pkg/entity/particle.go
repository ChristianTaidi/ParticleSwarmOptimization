package entity

import (
	"fmt"
	"math/rand"
)

type Particle struct {
	// All properties are set as slice to handle n variable optimization
	// each variable represents a position and velocity component
	position []float64
	velocity []float64
	bestPos  []float64
}

func (p *Particle) New(dimension int) {
	p.position = make([]float64, dimension)
	p.velocity = make([]float64, dimension)
	p.bestPos = make([]float64, dimension)
	for i := range p.position {
		p.position[i] = rand.Float64()*10 - 5
		p.velocity[i] = rand.Float64()*10 - 5
	}
	copy(p.bestPos, p.position)
}

func (p *Particle) Update(bestSwarmPosition []float64) {
	for i := range p.velocity {
		r1 := rand.Float64()
		r2 := rand.Float64()
		fmt.Printf("PPOS %+v\n", p.position)
		fmt.Printf("PVEL %+v\n", p.velocity)
		fmt.Printf("PBPOS %+v\n", p.bestPos)
		fmt.Printf("SBPOS %+v\n", bestSwarmPosition)

		p.velocity[i] = 0.5*p.velocity[i] + 2.0*r1*(p.bestPos[i]-p.position[i]) + 2.0*r2*(bestSwarmPosition[i]-p.position[i])
		p.position[i] = p.position[i] + p.velocity[i]
	}
}

func (p *Particle) UpdateBestPos(optimizationFunction func([]float64) float64) []float64 {
	if optimizationFunction(p.position) < optimizationFunction(p.bestPos) {
		p.bestPos = make([]float64, len(p.position))
		copy(p.bestPos, p.position)
	}
	return p.bestPos
}
