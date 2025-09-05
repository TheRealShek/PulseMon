package types

import (
	"context"
	"time"
)

type Metric struct {
	CPUUsage  float64
	RAMUsage  float64
	Timestamp time.Time
}

// Collector are things that know how to collect metrics
// This is a interface with 2 Methods - Collect and Name
type Collector interface {
	Collect(ctx context.Context) (Metric, error)
	Name() string
}
