package ai

import (
	"math/rand"
)

// Scenario input struct
type Scenario struct {
	Budget      float64  `json:"budget"`
	Grains      Calendar `json:"grains"`
	Inputs      Calendar `json:"inputs"`
	Maintenance Calendar `json:"maintenance"`
	Harvest     Calendar `json:"harvest"`
	Contracts   Calendar `json:"contracts"`
	Transport   Calendar `json:"transport"`
	Route       Calendar `json:"route"`
	Price       Calendar `json:"price"`
}

// Perform a tournament and select the best individual
func (scenario Scenario) Tournament(genes []Gene, k uint) Gene {
	// Participants map
	participants := map[uint]float64{}
	// Select k individuals
	for i := uint(0); i < k; i++ {
		// Pick random individual
	Generate:
		selected := uint(rand.Intn(len(genes)))
		if _, ok := participants[selected]; ok == true {
			goto Generate
		}
		// Perform fitness function
		geneStats := Stats(genes[selected])
		participants[selected] = geneStats.Fitness(scenario)
	}
	// Perform tournament
	result := 0.0
	winner := uint(rand.Intn(len(genes)))
	for k, f := range participants {
		if f > result {
			result = f
			winner = k
		}
	}
	return genes[winner]
}

// Processing function
func (s Scenario) Process(population, generations, parents, competitors uint, mutation float64) (Gene, float64) {
	// Perform parent check
	if parents > population {
		panic("parents must be less than or equal to population count")
	}
	// Spawn initial population
	p := SpawnPopulation(population)
	// Pass generations
	for i := uint(0); i < generations; i++ {
		// Log winners (so there are no repeated winners)
		winners := []Gene{}
		// Pick parents
		for j := uint(0); j < parents; j++ {
		PerformTournament:
			w := s.Tournament(p, competitors)
			// Check if winner has won previously
			for _, k := range winners {
				if k == w {
					goto PerformTournament
				}
			}
			// Add to winners
			winners = append(winners, w)
		}
		// Perform crossover
		p = Crossover(winners...)
		for j := range p {
			p[j].Mutate(mutation)
		}
	}

	// Final winner checking
	var best Gene
	var bestFitness float64
	for i, g := range p {
		fitness := Stats(g).Fitness(s)
		if i == 0 {
			best = g
			bestFitness = fitness
		} else if fitness > bestFitness {
			best = g
			bestFitness = fitness
		}
	}

	return best, bestFitness
}
