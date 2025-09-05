package collectors

import (
	"context"
	"github.com/shirou/gopsutil/v3/cpu"
	"main/pkg/types"
	"time"
)

type CPUCollector struct{}

func (c CPUCollector) Name() string {
	return "cpu"
}

func (c *CPUCollector) Collect(ctx context.Context) (types.Metric, error) {
	percent, err := cpu.Percent(0, false)
	if err != nil || len(percent) == 0 {
		return types.Metric{}, err
	}
	return types.Metric{
		CPUUsage:  percent[0],
		Timestamp: time.Now(),
	}, nil
}
