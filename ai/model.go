package ai

import (
	"errors"
	"math"
	"math/rand"
	"unsafe"
)

// Gene struct
type Gene struct {
	Inputs      uint16
	Maintenance uint16
	Contracts   uint16
	Transport   uint16
}

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

type GeneBuffer [unsafe.Sizeof(Gene{})]byte

// Decode gene data into gene
func (g *Gene) Decode(data []byte) error {
	if data == nil {
		return errors.New("cannot decode nil data into gene")
	}
	// Allocate enough bytes to fill model
	plate := GeneBuffer{}
	// Get smallest value
	min := len(plate)
	if len(data) < min {
		min = len(data)
	}
	// Fill plate with input data and fill gene
	for i := 0; i < min; i++ {
		plate[i] = data[i]
	}
	g.Inputs = uint16(plate[0])<<010 | uint16(plate[1])
	g.Maintenance = uint16(plate[2])<<010 | uint16(plate[3])
	g.Contracts = uint16(plate[4])<<010 | uint16(plate[5])
	g.Transport = uint16(plate[6])<<010 | uint16(plate[7])
	return nil
}

// Decode gene data into gene
func (g *Gene) Encode() (data *[]byte) {
	// Allocate slice of bytes as result
	result := []byte{}
	// Insert gene values in order
	*data = append(result,
		byte(g.Inputs>>010), byte(g.Inputs&0xff),
		byte(g.Maintenance>>010), byte(g.Maintenance&0xff),
		byte(g.Contracts>>010), byte(g.Contracts&0xff),
		byte(g.Transport>>010), byte(g.Transport&0xff),
	)
	return data
}

// Get gene stats as percentage
func Stats(gene Gene) (result GeneStats) {
	// Total value
	total := uint64(gene.Inputs) + uint64(gene.Maintenance) + uint64(gene.Transport) + uint64(gene.Contracts)
	// Calculate percentage
	result.Inputs = float64(gene.Inputs) / float64(total)
	result.Maintenance = float64(gene.Maintenance) / float64(total)
	result.Contracts = float64(gene.Contracts) / float64(total)
	result.Transport = float64(gene.Transport) / float64(total)
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
		yield += totalGrains
	}
	return yield
}
