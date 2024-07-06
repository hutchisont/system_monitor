package SystemInfo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// would be really nice to get name as well, will have to figure out where to
// get that from

type GPU struct {
	gpuPCIDeviceID  string
	gpuBusyPercent  float64
	vramBusyPercent float64
	vramTotal       float64
	vramUsed        float64
	vramAvailable   float64
}

func (g GPU) String() string {
	return fmt.Sprintf("GPU Usage: %.0f%%\nVRAM Usage: %.0f%%\nVRAM Total: %.2f GB\nVRAM Used: %.2f GB\nVRAM Available: %.2f GB",
		g.gpuBusyPercent, g.vramBusyPercent, g.vramTotal, g.vramUsed, g.vramAvailable)
}

func (g *GPU) updateGPUReading() {
	g.gpuBusyPercent, g.vramBusyPercent, g.vramTotal, g.vramUsed, g.vramAvailable = getGPUInfoData()
}

func getGPUInfoData() (gpuBusyPercentF, vramBusyPercentF, vramTotalF, vramUsedF, vramAvailableF float64) {
	// TODO(Tyler): Figure out how to dynamically determine the correct ID here
	const gpuID = "0000:0b:00.0"
	const baseLocation = "/sys/bus/pci/devices/" + gpuID + "/"
	gpuBusyPercent, err := os.ReadFile(baseLocation + "gpu_busy_percent")
	if err != nil {
		gpuBusyPercent = []byte("-1")
	}
	memBusyPercent, err := os.ReadFile(baseLocation + "mem_busy_percent")
	if err != nil {
		memBusyPercent = []byte("-1")
	}
	vramTotal, err := os.ReadFile(baseLocation + "mem_info_vram_total")
	if err != nil {
		vramTotal = []byte("-1")
	}
	vramUsed, err := os.ReadFile(baseLocation + "mem_info_vram_used")
	if err != nil {
		vramUsed = []byte("-1")
	}

	gpuBusyPercentF, err = strconv.ParseFloat(strings.TrimSpace(string(gpuBusyPercent)), 64)
	if err != nil {
		gpuBusyPercentF = -1
	}
	vramBusyPercentF, err = strconv.ParseFloat(strings.TrimSpace(string(memBusyPercent)), 64)
	if err != nil {
		vramBusyPercentF = -1
	}
	vramTotalF, err = strconv.ParseFloat(strings.TrimSpace(string(vramTotal)), 64)
	if err != nil {
		vramTotalF = -1
	}
	vramUsedF, err = strconv.ParseFloat(strings.TrimSpace(string(vramUsed)), 64)
	if err != nil {
		vramUsedF = -1
	}

	if vramUsedF == -1 || vramTotalF == -1 {
		vramAvailableF = -1
	} else {
		vramAvailableF = vramTotalF - vramUsedF
	}

	return gpuBusyPercentF, vramBusyPercentF, byteTogigaByte(vramTotalF),
		byteTogigaByte(vramUsedF), byteTogigaByte(vramAvailableF)

}
