package SystemInfo

import "fmt"

type SystemInfo struct {
	CPU
	RAM
	GPU
}

func (s SystemInfo) String() string {
	return fmt.Sprintf("%v\n\n%v\n\n%v", s.CPU, s.RAM, s.GPU)
}

func (s *SystemInfo) UpdateAllReadings() {
	s.RAM.updateRAMReading()
	s.CPU.updateCPUReading()
	s.GPU.updateGPUReading()
}
