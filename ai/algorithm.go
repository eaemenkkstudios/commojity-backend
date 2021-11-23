package ai

import (
	"fmt"
	"math/rand"
)

// Spawn k individuals
func SpawnPopulation(k uint) []GeneBuffer {
	population := make([]GeneBuffer, k)
	for i := range population {
		rand.Read(population[i][:])
	}
	return population
}

// Perform a tournament and select the best individual
func Tournament(scenario *Scenario, genes []Gene, k uint) uint {
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
	return winner
}

// Permute function
func permute(k uint64, result *[]Gene, units []Unit) {
	if k == 1 {
		g := Gene{}
		g.DecodeUnits(units)
		*result = append(*result, g)
	} else {
		permute(k-1, result, units)
		for i := uint64(0); i < k-1; i++ {
			if k&1 == 0 {
				a := units[i]
				units[i] = units[k-1]
				units[k-1] = a
			} else {
				a := units[0]
				units[0] = units[k-1]
				units[k-1] = a
			}
			permute(k-1, result, units)
		}
	}
}

// Crossover function
func Crossover(genes ...Gene) []Gene {
	if len(genes) < UNIT_COUNT {
		panic(fmt.Errorf("not enough individuals to crossover (received %d, but must be at least %d)", len(genes), UNIT_COUNT))
	}
	parents := []Unit{}
	for _, g := range genes {
		parents = append(parents, g[:]...)
	}
	population := []Gene{}
	permute(UNIT_COUNT, &population, parents)
	return population
}
