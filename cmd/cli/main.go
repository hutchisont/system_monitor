package main

import (
	"fmt"

	SystemInfo "github.com/hutchisont/system_monitor/internal/system_info"
)

func main() {
	sysInfo := SystemInfo.SystemInfo{}
	sysInfo.UpdateAllReadings()
	fmt.Println(sysInfo)
}
