package main

import (
	"fmt"

	"github.com/eaemenkkstudios/commojity-backend/ai"
	"github.com/eaemenkkstudios/commojity-backend/api"
)

func main() {
	scenario := ai.Scenario{
		Budget:      100000,
		Grains:      ai.Calendar{1000, 900, 800, 700, 800, 500, 900, 1000, 1000, 200, 300, 400},
		Inputs:      ai.Calendar{9, 8, 7, 8, 5, 9, 10, 10, 2, 3, 4, 10},
		Maintenance: ai.Calendar{80, 70, 80, 50, 90, 100, 100, 20, 30, 40, 100, 90},
		Harvest:     ai.Calendar{.1, .2, .3, .1, .3, .3, .2, .1, .4, .1, .2, .3},
		Contracts:   ai.Calendar{1, 2, 9, 10, 9, 8, 7, 10, 10, 2, 3, 4},
		Transport:   ai.Calendar{8, 5, 9, 10, 10, 2, 3, 4, 10, 9, 8, 7},
		Route:       ai.Calendar{.2, .2, .1, .3, .1, .2, .1, .1, .1, .3, .1, .2},
		Price:       ai.Calendar{100, 180, 200, 200, 200, 180, 100, 200, 250, 50, 80, 100},
	}

	population := ai.SpawnPopulation(10)
	for _, g := range population {
		fmt.Println(g, ai.Stats(g).Fitness(scenario))
	}

	g1 := scenario.Tournament(population, 4)
	fmt.Println("G1:", g1, ai.Stats(g1).Fitness(scenario))

FindG2:
	g2 := scenario.Tournament(population, 4)
	if g2 == g1 {
		goto FindG2
	}
	fmt.Println("G2:", g2, ai.Stats(g2).Fitness(scenario))

FindG3:
	g3 := scenario.Tournament(population, 4)
	if g3 == g2 || g3 == g1 {
		goto FindG3
	}
	fmt.Println("G3:", g3, ai.Stats(g3).Fitness(scenario))

FindG4:
	g4 := scenario.Tournament(population, 4)
	if g4 == g3 || g4 == g2 || g4 == g1 {
		goto FindG4
	}
	fmt.Println("G4:", g4, ai.Stats(g4).Fitness(scenario))

	population = ai.Crossover(g1, g2, g3, g4)
	for i, g := range population {
		population[i].Mutate(.2)
		fmt.Println(g, ai.Stats(g).Fitness(scenario))
	}

	return
	r := api.InitializeEngine()
	api.RegisterRoutes(r)
	r.Run()
}
