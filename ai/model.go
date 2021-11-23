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

// Decode gene data into gene
func (g *Gene) Decode(data []byte) error {
	if data == nil {
		return errors.New("cannot decode nil data into gene")
	}
	// Allocate enough bytes to fill model
	buffer := GeneBuffer{}
	// Fill buffer with input data and fill gene
	copy(buffer[:], data)
	g[0] = Unit(buffer[0])<<010 | Unit(buffer[1])
	g[1] = Unit(buffer[2])<<010 | Unit(buffer[3])
	g[2] = Unit(buffer[4])<<010 | Unit(buffer[5])
	g[3] = Unit(buffer[6])<<010 | Unit(buffer[7])
	return nil
}

// Decode gene data into gene
func (g *Gene) Encode() (data []byte) {
	// Allocate slice of bytes as result
	result := []byte{}
	// Insert gene values in order
	data = append(result,
		byte(g[0]>>010), byte(g[0]&0xff),
		byte(g[1]>>010), byte(g[1]&0xff),
		byte(g[2]>>010), byte(g[2]&0xff),
		byte(g[3]>>010), byte(g[3]&0xff),
	)
	return data
}

// Get gene stats as percentage
func Stats(gene Gene) (result GeneStats) {
	// Total value
	total := uint64(gene[0]) + uint64(gene[1]) + uint64(gene[2]) + uint64(gene[3])
	// Calculate percentage
	result.Inputs = float64(gene[0]) / float64(total)
	result.Maintenance = float64(gene[1]) / float64(total)
	result.Contracts = float64(gene[2]) / float64(total)
	result.Transport = float64(gene[3]) / float64(total)
	return
}

// Fitness function
func (g *GeneStats) Fitness(s *Scenario) float64 {
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
		vehicleCount := math.Floor(math.Log(budgetContracts) / s.Contracts[m])
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
	return yield
}
