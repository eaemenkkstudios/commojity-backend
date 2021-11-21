package ai

import "math/rand"

// Spawn k individuals
func SpawnPopulation(k uint) []GeneBuffer {
	population := make([]GeneBuffer, k)
	for i := range population {
		rand.Read(population[i][:])
	}
	return population
}

// Perform a tournament and select the best individual
func Tournament(scenario *Scenario, genes []GeneBuffer, k uint) uint {
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
		gene := Gene{}
		gene.Decode(genes[selected][:])
		geneStats := Stats(gene)
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
	return winner
}
