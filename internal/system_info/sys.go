package SystemInfo

type SystemInfo struct {
	CPU
	RAM
}

func (s *SystemInfo) UpdateAllReadings() {
	s.RAM.updateRAMReading()
	s.CPU.updateCPUReading()
}
