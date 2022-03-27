package statspkg

import (
	"time"

	"math"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/wailsapp/wails"
)

type MemStats struct {
	log *wails.CustomLogger
}

type MemUsage struct {
	Avg int `json:"avg"`
}

func (s *MemStats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")

	go func() {
		for {
			u, err := s.GetMemUsage()
			if err != nil {
				s.log.Errorf("Error getting Mem usage: %v", err)
			}
			runtime.Events.Emit("mem", u)
			time.Sleep(time.Second)
		}
	}()

	return nil
}

func (s *MemStats) GetMemUsage() (MemUsage, error) {
	avg, _ := mem.VirtualMemory()
	memp := avg.UsedPercent
	return MemUsage{
		Avg: int(math.Round(memp)),
	}, nil
}
