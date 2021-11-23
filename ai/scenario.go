package ai

import "math/rand"

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
