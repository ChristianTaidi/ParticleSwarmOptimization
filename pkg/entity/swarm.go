package entity

type Swarm struct {
	particles            []*Particle
	bestPos              []float64
	optimizationFunction func([]float64) float64
}

func NewSwarm(nParticles, dimension int, optimization func([]float64) float64) *Swarm {
	particles := make([]*Particle, nParticles)
	for i := range particles {
		particles[i] = &Particle{}
		particles[i].New(dimension)
	}
	return &Swarm{
		particles:            particles,
		bestPos:              make([]float64, dimension),
		optimizationFunction: optimization,
	}
}

func (s *Swarm) Update() {
	for _, particle := range s.particles {
		particle.Update(s.bestPos)
		pBestPos := particle.UpdateBestPos(s.optimizationFunction)
		if s.optimizationFunction(pBestPos) < s.optimizationFunction(s.bestPos) {
			s.bestPos = make([]float64, len(pBestPos))
			copy(s.bestPos, pBestPos)
		}
	}
}

func (s *Swarm) BestPosResult() ([]float64, float64) {
	return s.bestPos, s.optimizationFunction(s.bestPos)
}
