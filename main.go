package main

import (
	"fmt"

	"github.com/eaemenkkstudios/commojity-backend/ai"
	"github.com/eaemenkkstudios/commojity-backend/api"
)

func main() {
	population := ai.SpawnPopulation(100)
	for _, p := range population {
		g := ai.Gene{}
		g.Decode(p[:])
		fmt.Println(ai.Stats(g))
	}
	return
	r := api.InitializeEngine()
	api.RegisterRoutes(r)
	r.Run()
}
