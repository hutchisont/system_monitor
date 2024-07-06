package SystemInfo

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// TODO(Tyler): Add getting usage at some point, probably overall and per core
type CPU struct {
	model string
	count int
	// usage     float64
	frequency []float64
}

func (c CPU) String() string {
	return fmt.Sprintf("Model: %v\nThread Count: %v\nPer Core Frequency: %v",
		c.model, c.count, c.frequency)
	// return fmt.Sprintf("Model: %v\nCore Count: %v\nPer Core Usage: %v\nPer Core Frequency: %v\n",
	// 	c.model, c.count, c.usage, c.frequency)
}

func (c *CPU) UpdateCPUReading() {
	data := readCPUInfo()
	c.updateCPUinfoFromData(data)

}

var modelNameRegex = regexp.MustCompile("model name\\s*:\\s*(?P<ModelName>.*)")
var coreCountRegex = regexp.MustCompile("siblings\\s*:\\s*(?P<CoreCount>\\d*)")
var frequencyRegex = regexp.MustCompile("cpu MHz\\s*:\\s*(?P<Frequency>\\d*)")

func (c *CPU) updateCPUinfoFromData(data []byte) {
	if c.model == "" {
		matches := modelNameRegex.FindSubmatch(data)
		c.model = string(matches[modelNameRegex.SubexpIndex("ModelName")])
	}

	if c.count == 0 {
		matches := coreCountRegex.FindSubmatch(data)
		count, err := strconv.
			ParseInt(string(matches[coreCountRegex.SubexpIndex("CoreCount")]), 10, 64)
		if err != nil {
			log.Fatalln("Could not convert core count", err)
		}
		c.count = int(count)
	}

	if len(c.frequency) == 0 {
		c.frequency = make([]float64, c.count)
	}

	matches := frequencyRegex.FindAllSubmatch(data, c.count)
	for i := 0; i < c.count; i++ {
		frequency, err := strconv.
			ParseFloat(string(matches[i][frequencyRegex.SubexpIndex("Frequency")]), 64)
		if err != nil {
			log.Fatalln("Could not convert frequency for core: ", i)
		}
		c.frequency[i] = frequency
	}
}

func readCPUInfo() (data []byte) {
	const cpuInfoFilePath = "/proc/cpuinfo"
	data, err := os.ReadFile(cpuInfoFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
