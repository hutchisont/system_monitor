package SystemInfo

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// We are reading from /proc/meminfo to get information on our RAM.
// https://github.com/torvalds/linux/blob/master/Documentation/filesystems/proc.rst#meminfo

type RAM struct {
	totalRAM     float64
	availableRAM float64
}

func (r *RAM) UpdateRAMReading() {
	data := getMeminfo()
	r.updateMeminfoFromData(data)
}

func (r RAM) String() string {
	return fmt.Sprintf("Total RAM: %.2fGB\nAvailable RAM: %.2fGB", r.totalRAM, r.availableRAM)
}

var memTotalRegex = regexp.MustCompile("MemTotal:\\s*(?P<Total>\\d*)")
var memAvailableRegex = regexp.MustCompile("MemAvailable:\\s*(?P<Available>\\d*)")

func (r *RAM) updateMeminfoFromData(data []byte) {
	totalMatches := memTotalRegex.FindSubmatch(data)
	total := string(totalMatches[memTotalRegex.SubexpIndex("Total")])
	memAvailableMatches := memAvailableRegex.FindSubmatch(data)
	avail := string(memAvailableMatches[memAvailableRegex.SubexpIndex("Available")])
	totalkB, err := strconv.ParseFloat(total, 64)
	if err != nil {
		log.Fatalln(err)
	}
	availkB, err := strconv.ParseFloat(avail, 64)
	if err != nil {
		log.Fatalln(err)
	}

	r.totalRAM = kiloByteTogigaByte(totalkB)
	r.availableRAM = kiloByteTogigaByte(availkB)
}

func getMeminfo() (data []byte) {
	const meminfoFilePath = "/proc/meminfo"
	data, err := os.ReadFile(meminfoFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
