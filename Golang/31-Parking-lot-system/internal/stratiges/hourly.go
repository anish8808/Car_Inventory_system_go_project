package stratiges

import (
	"parkinglot/internal/model"
	"time"
)

type HourlyStrategy struct {
	Rates map[model.VehicleType]float64
}

func (h *HourlyStrategy) CalculateCost(entry, exit time.Time, vType model.VehicleType) float64 {
	duration := exit.Sub(entry).Hours()
	return h.Rates[vType] * duration
}
