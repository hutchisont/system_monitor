package SystemInfo

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// We are reading from /proc/meminfo to get information on our RAM.
// https://github.com/torvalds/linux/blob/master/Documentation/filesystems/proc.rst#meminfo

type RAM struct {
	TotalRAM     float64
	AvailableRAM float64
}

func (r *RAM) updateRAMReading() {
	data := getMeminfo()
	r.updateMeminfoFromData(data)
}

func (r RAM) String() string {
	return fmt.Sprintf("Total RAM: %.2f GB\nAvailable RAM: %.2f GB",
		r.TotalRAM, r.AvailableRAM)
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
		totalkB = -1
	}
	availkB, err := strconv.ParseFloat(avail, 64)
	if err != nil {
		availkB = -1
	}

	r.TotalRAM = kiloByteTogigaByte(totalkB)
	r.AvailableRAM = kiloByteTogigaByte(availkB)
}

func getMeminfo() (data []byte) {
	const meminfoFilePath = "/proc/meminfo"
	data, _ = os.ReadFile(meminfoFilePath)
	return data
}
