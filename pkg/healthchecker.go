package pkg

import (
	"math"
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func CheckHealth() (HealthStatus, error) {

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return HealthStatus{}, err
	}

	cpuPercentages, err := cpu.Percent(0, false)
	if err != nil {
		return HealthStatus{}, err
	}

	cpu := runtime.NumCPU()
	cpuPercentage := math.Round(cpuPercentages[0]*100) / 100

	ram := vmStat.Total
	ramMax := ram + vmStat.Free
	ramPercentage := vmStat.UsedPercent

	result := HealthStatus{
		CPU:           float64(cpu),
		CPUPercentage: float64(cpuPercentage),
		RAM:           float64(ram),
		RAMMax:        float64(ramMax),
		RAMPercentage: float64(ramPercentage),
	}

	return result, nil
}
