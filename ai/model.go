package ai

import (
	"errors"
	"unsafe"

	"github.com/eaemenkkstudios/commojity-backend/data"
)

// Gene struct
type Gene struct {
	Inputs      uint16
	Maintenance uint16
	Contracts   uint16
	Transport   uint16
}

// Gene Stats struct
type GeneStats struct {
	Inputs      float64 `json:"inputs"`
	Maintenance float64 `json:"maintenance"`
	Contracts   float64 `json:"contracts"`
	Transport   float64 `json:"transport"`
}

// Decode gene data into gene
func (g *Gene) Decode(data data.GeneData) error {
	if data == nil {
		return errors.New("cannot decode nil data into gene")
	}
	// Allocate enough bytes to fill model
	plate := [unsafe.Sizeof(Gene{})]byte{}
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
func (g *Gene) Encode() (data *data.GeneData) {
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
func (g *Gene) Stats() (result GeneStats) {
	// Total value
	total := uint64(g.Inputs) + uint64(g.Maintenance) + uint64(g.Transport) + uint64(g.Contracts)
	// Calculate percentage
	result.Inputs = float64(g.Inputs) / float64(total) * 100
	result.Maintenance = float64(g.Maintenance) / float64(total) * 100
	result.Contracts = float64(g.Contracts) / float64(total) * 100
	result.Transport = float64(g.Transport) / float64(total) * 100
	return
}
