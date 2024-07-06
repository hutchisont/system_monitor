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
	Model     string
	Count     int
	Frequency []float64
}

func (c CPU) String() string {
	return fmt.Sprintf("Model: %v\nThread Count: %v\nPer Thread Frequency: %v",
		c.Model, c.Count, c.Frequency)
}

func (c *CPU) updateCPUReading() {
	data := readCPUInfo()
	c.updateCPUinfoFromData(data)
}

var modelNameRegex = regexp.MustCompile("model name\\s*:\\s*(?P<ModelName>.*)")
var coreCountRegex = regexp.MustCompile("siblings\\s*:\\s*(?P<CoreCount>\\d*)")
var frequencyRegex = regexp.MustCompile("cpu MHz\\s*:\\s*(?P<Frequency>\\d*)")

func (c *CPU) updateCPUinfoFromData(data []byte) {
	if c.Model == "" {
		matches := modelNameRegex.FindSubmatch(data)
		c.Model = string(matches[modelNameRegex.SubexpIndex("ModelName")])
	}

	if c.Count == 0 {
		matches := coreCountRegex.FindSubmatch(data)
		count, err := strconv.
			ParseInt(string(matches[coreCountRegex.SubexpIndex("CoreCount")]), 10, 64)
		if err != nil {
			count = -1
		}
		c.Count = int(count)
	}

	if len(c.Frequency) == 0 {
		c.Frequency = make([]float64, c.Count)
	}

	if c.Count > 0 {
		matches := frequencyRegex.FindAllSubmatch(data, c.Count)
		for i := 0; i < c.Count; i++ {
			frequency, err := strconv.
				ParseFloat(string(matches[i][frequencyRegex.SubexpIndex("Frequency")]), 64)
			if err != nil {
				frequency = -1
			}
			c.Frequency[i] = frequency
		}

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
