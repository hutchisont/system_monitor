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

func (r *RAM) updateMeminfoFromData(data []byte) {
	// TODO(Tyler): this is really not ideal, if the order of mem total and mem
	// available ever change then this will break
	reg := regexp.
		MustCompile("MemTotal:\\s*(?P<Total>\\d*)|MemAvailable:\\s*(?P<Available>\\d*)")
	matches := reg.FindAllSubmatch(data, -1)
	total := string(matches[0][reg.SubexpIndex("Total")])
	avail := string(matches[1][reg.SubexpIndex("Available")])
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
