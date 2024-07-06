package SystemInfo

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// TODO(Tyler): This will have to figure out the pci device number of the GPU
// and then read from /sys/bus/pci/devices/<id>/
// I don't know what it exactly needs to read yet but some likely ones will be
// mem_busy_percent
// mem_info_vram_total
// mem_info_vram_used
// gpu_busy_percent
//
// would be really nice to get name as well, will have to figure out where to
// get that from

const pciDevicesBasePath = "/sys/bus/pci/devices/"

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
	const baseLocation = "/sys/bus/pci/devices/0000:0b:00.0/"
	gpuBusyPercent, err := os.ReadFile(baseLocation + "gpu_busy_percent")
	if err != nil {
		log.Fatalln(err)
	}
	memBusyPercent, err := os.ReadFile(baseLocation + "mem_busy_percent")
	if err != nil {
		log.Fatalln(err)
	}
	vramTotal, err := os.ReadFile(baseLocation + "mem_info_vram_total")
	if err != nil {
		log.Fatalln(err)
	}
	vramUsed, err := os.ReadFile(baseLocation + "mem_info_vram_used")
	if err != nil {
		log.Fatalln(err)
	}

	gpuBusyPercentF, err = strconv.ParseFloat(strings.TrimSpace(string(gpuBusyPercent)), 64)
	if err != nil {
		log.Fatalln(err)
	}
	vramBusyPercentF, err = strconv.ParseFloat(strings.TrimSpace(string(memBusyPercent)), 64)
	if err != nil {
		log.Fatalln(err)
	}
	vramTotalF, err = strconv.ParseFloat(strings.TrimSpace(string(vramTotal)), 64)
	if err != nil {
		log.Fatalln(err)
	}
	vramUsedF, err = strconv.ParseFloat(strings.TrimSpace(string(vramUsed)), 64)
	if err != nil {
		log.Fatalln(err)
	}

	vramAvailableF = vramTotalF - vramUsedF

	return gpuBusyPercentF, vramBusyPercentF, byteTogigaByte(vramTotalF),
		byteTogigaByte(vramUsedF), byteTogigaByte(vramAvailableF)

}
