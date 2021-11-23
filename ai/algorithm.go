package ai

import (
	"math/rand"
)

// Spawn k individuals
func SpawnPopulation(k uint) []Gene {
	population := make([]Gene, k)
	for i := range population {
		buffer := GeneBuffer{}
		rand.Read(buffer[:])
		population[i].DecodeBuffer(buffer)
	}
	return population
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
	parents := []Unit{}
	for _, g := range genes {
		parents = append(parents, g[:]...)
	}
	population := []Gene{}
	permute(UNIT_COUNT, &population, parents)
	return population
}
