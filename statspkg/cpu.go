package statspkg

import (
	"time"

	"math"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/wailsapp/wails"
)

type CpuStats struct {
	log *wails.CustomLogger
}

type CPUUsage struct {
	Avg int `json:"avg"`
}

func (s *CpuStats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")

	go func() {
		for {
			u, err := s.GetCPUUsage()
			if err != nil {
				s.log.Errorf("Error getting CPU usage: %v", err)
			}
			runtime.Events.Emit("cpu", u)
			time.Sleep(time.Second)
		}
	}()

	return nil
}

func (s *CpuStats) GetCPUUsage() (CPUUsage, error) {
	avg, err := cpu.Percent(time.Second, false)
	if err != nil {
		s.log.Errorf("Error getting CPU usage: %v", err)
	}
	return CPUUsage{
		Avg: int(math.Round(avg[0])),
	}, nil
}
