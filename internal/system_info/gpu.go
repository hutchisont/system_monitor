package SystemInfo

import "fmt"

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
	return fmt.Sprintf("GPU Usage: %.2f\nVRAM Usage: %.2f\nVRAM Total: %.2f\nVRAM Used: %.2f\nVRAM Available: %.2f",
		g.gpuBusyPercent, g.vramBusyPercent, g.vramTotal, g.vramUsed, g.vramAvailable)
}

func (g *GPU) UpdateGPUReading() {

}

// going to have to get data from multiple files. maybe create and return a map?
// might be more performant to just return a [][]byte and know for myself
// what index means for each
func getGPUInfoData() {

}
