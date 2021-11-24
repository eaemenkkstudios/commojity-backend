package ai

import (
	"errors"
	"math"
	"math/rand"
	"unsafe"
)

// Unit count in gene
const UNIT_COUNT = 4

// Gene unit
type Unit uint16

// Gene buffer type
type GeneBuffer [UNIT_COUNT * unsafe.Sizeof(Unit(0))]byte

// Gene type
type Gene [UNIT_COUNT]Unit

// Gene stats struct
type GeneStats struct {
	Inputs      float64 `json:"inputs"`
	Maintenance float64 `json:"maintenance"`
	Contracts   float64 `json:"contracts"`
	Transport   float64 `json:"transport"`
}

// Months range
const MONTHS = 12

// Floats for every month in a year
type Calendar [MONTHS]float64

// Decode gene data (in units) into gene
func (g *Gene) DecodeUnits(data []Unit) error {
	if data == nil {
		return errors.New("cannot decode nil data into gene")
	}
	// Fill buffer with input data and fill gene
	copy(g[:], data)
	return nil
}

// Decode gene buffer into gene
func (g *Gene) DecodeBuffer(data GeneBuffer) error {
	// Fill buffer with input data and fill gene
	for i := range g {
		g[i] = Unit(data[i*2])<<010 | Unit(data[i*2+1])
	}
	return nil
}

// Decode gene data into gene
func (g *Gene) Decode(data []byte) error {
	if data == nil {
		return errors.New("cannot decode nil data into gene")
	}
	// Allocate enough bytes to fill model
	buffer := GeneBuffer{}
	// Fill buffer with input data and fill gene
	copy(buffer[:], data)
	g.DecodeBuffer(buffer)
	return nil
}

// Decode gene data into gene
func (g *Gene) Encode() (data GeneBuffer) {
	// Insert gene values in order
	for i := range data {
		if i&1 == 0 {
			data[i] = byte(g[i/2] >> 010)
		} else {
			data[i] = byte(g[i/2] & 0xff)
		}
	}
	return data
}

// Get gene stats as percentage
func Stats(g Gene) (result GeneStats) {
	// Total value
	total := uint64(g[0]) + uint64(g[1]) + uint64(g[2]) + uint64(g[3])
	// Calculate percentage
	result.Inputs = float64(g[0]) / float64(total)
	result.Maintenance = float64(g[1]) / float64(total)
	result.Contracts = float64(g[2]) / float64(total)
	result.Transport = float64(g[3]) / float64(total)
	return
}

// Mutate gene
func (g *Gene) Mutate(p float64) {
	chance := math.Min(math.Max(p, 0), 1)
	for i := range g {
		g[i] = Unit(float64(g[i]) * (rand.Float64()*chance + 1))
	}
}

// Fitness function
func (g GeneStats) Fitness(s Scenario) float64 {
	// Total profit (starts with 100% loss)
	yield := -s.Budget
	// Budget monthly spent with Inputs
	budgetInputs := s.Budget * g.Inputs / MONTHS
	// Budget monthly spent with Maintenance
	budgetMaintenance := s.Budget * g.Maintenance / MONTHS
	// Budget monthly spent with Contracts
	budgetContracts := s.Budget * g.Contracts / MONTHS
	// Budget monthly spent with Transport
	budgetTransport := s.Budget * g.Transport / MONTHS
	// Perform calculations
	for m := 0; m < MONTHS; m++ {
		// Grains farmed this month
		grainProduction := s.Grains[m]
		// Potential farm outcome
		potentialFarmOutcome := grainProduction * math.Pow(budgetInputs, 2) / 2
		// Maintenance rate
		maintenanceRate := math.Min(budgetMaintenance/s.Maintenance[m], 1)
		// Actual farm outcome
		farmOutcome := maintenanceRate * potentialFarmOutcome
		// Random farm incident
		if maintenanceRate < .75 && rand.Float32() < float32(s.Harvest[m]) {
			farmOutcome -= grainProduction * .1
		}
		// Grain-transporting vehicles count
		vehicleCount := 0.0
		if budgetContracts > 0 {
			vehicleCount = math.Floor(math.Log(budgetContracts) / s.Contracts[m])
		}
		// Capacity of grains that a vehicle can transport
		vehicleCapacity := math.Floor(budgetTransport / s.Transport[m])
		// Grains that were initially transported
		totalGrains := math.Min(vehicleCount*vehicleCapacity, farmOutcome)
		// Random transport incident
		for t := 0; t < int(vehicleCount); t++ {
			if rand.Float32() < float32(s.Route[m]) {
				// Grains loss
				totalGrains -= vehicleCapacity * .23
			}
		}
		// Sell grains
		yield += totalGrains * s.Price[m]
	}
	return math.Max(yield, -s.Budget)
}
