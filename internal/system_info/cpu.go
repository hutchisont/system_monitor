package SystemInfo

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

type CPU struct {
	model     string
	count     int
	usage     []float64
	frequency []float64
}

func (c CPU) String() string {
	return fmt.Sprintf("Model: %v\nCore Count: %v\nPer Core Usage: %v\nPer Core Frequency: %v\n",
		c.model, c.count, c.usage, c.frequency)
}

func (c *CPU) UpdateCPUReading() {
	data := readCPUInfo()
	c.updateMeminfoFromData(data)

}

var modelNameRegex = regexp.MustCompile("model name\\s*:\\s(?P<ModelName>.*$)")
var coreCountRegex = regexp.MustCompile("siblings\\s*:\\s(?P<ModelName>\\d*$)")

func (c *CPU) updateMeminfoFromData(data []byte) {
	panic("not yet implemented")
}

func readCPUInfo() (data []byte) {
	const cpuInfoFilePath = "/proc/cpuinfo"
	data, err := os.ReadFile(cpuInfoFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
