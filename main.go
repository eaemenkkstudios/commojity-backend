package main

import (
	"fmt"

	"github.com/eaemenkkstudios/commojity-backend/ai"
	"github.com/eaemenkkstudios/commojity-backend/api"
)

func main() {
	a := ai.Crossover(ai.Gene{1, 2, 3, 4}, ai.Gene{5, 6, 7, 8}, ai.Gene{9, 10, 11, 12}, ai.Gene{13, 14, 15, 16})
	fmt.Println(a)
	return
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
	for _, p := range population {
		g := ai.Gene{}
		g.Decode(p[:])
		stats := ai.Stats(g)
		fmt.Println(stats.Fitness(&scenario))
	}
	return
	r := api.InitializeEngine()
	api.RegisterRoutes(r)
	r.Run()
}
