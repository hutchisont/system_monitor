package main

import (
	"fmt"

	SystemInfo "github.com/hutchisont/system_monitor/internal/system_info"
)

func main() {
	ram := SystemInfo.RAM{}
	ram.UpdateRAMReading()
	fmt.Println(ram)
	cpu := SystemInfo.CPU{}
	cpu.UpdateCPUReading()
	fmt.Println(cpu)
}
