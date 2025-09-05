package collectors

import (
	"context"
	"main/pkg/types"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
)

type RamCollector struct{}

func (r RamCollector) Name() string {
	return "ram"
}

func (r *RamCollector) Collect(ctx context.Context) (types.Metric, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return types.Metric{}, err
	}
	return types.Metric{
		RAMUsage:  vm.UsedPercent,
		Timestamp: time.Now(),
	}, nil
}
