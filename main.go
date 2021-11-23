package main

import (
	"github.com/eaemenkkstudios/commojity-backend/api"
)

func main() {
	r := api.InitializeEngine()
	api.RegisterRoutes(r)
	r.Run()
}
