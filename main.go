package main

import (
	"fmt"
	"time"

	SystemInfo "github.com/hutchisont/system_monitor/internal/system_info"
)

func main() {

	i := test(100)
	fmt.Println("did", 100, "readings in 1", i)
	i = test(1_000)
	fmt.Println("did", 1_000, "readings in", i)
	i = test(10_000)
	fmt.Println("did", 10_000, "readings in", i)
	i = test(100_000)
	fmt.Println("did", 100_000, "readings in", i)
	i = test(1_000_000)
	fmt.Println("did", 1_000_000, "readings in", i)
}

func test(iters int) (seconds time.Duration) {
	sysInfo := SystemInfo.SystemInfo{}

	start := time.Now()
	for i := 0; i < iters; i++ {
		sysInfo.UpdateAllReadings()
	}
	return time.Since(start)
}
