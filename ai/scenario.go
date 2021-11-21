package ai

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
