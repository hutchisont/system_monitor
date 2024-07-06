package main

import (
	"fmt"
	"time"

	SystemInfo "github.com/hutchisont/system_monitor/internal/system_info"
)

func main() {
	sysInfo := SystemInfo.SystemInfo{}
	sysInfo.UpdateAllReadings()
	fmt.Println(sysInfo)

	benchmark()
}

func benchmark() {

	fmt.Print("\nStarting benchmark\n\n")

	i := test(100)
	fmt.Printf("did 100 readings in %.5f seconds\n", i.Seconds())
	i = test(1_000)
	fmt.Printf("did 1,000 readings in %.5f seconds\n", i.Seconds())
	i = test(10_000)
	fmt.Printf("did 10,000 readings in %.5f seconds\n", i.Seconds())
	i = test(100_000)
	fmt.Printf("did 100,000 readings in %.5f seconds\n", i.Seconds())
	i = test(1_000_000)
	fmt.Printf("did 1,000,000 readings in %.5f seconds\n", i.Seconds())

	fmt.Print("\nBenchmark Finished")
}

func test(iters int) (seconds time.Duration) {
	sysInfo := SystemInfo.SystemInfo{}

	start := time.Now()
	for i := 0; i < iters; i++ {
		sysInfo.UpdateAllReadings()
	}
	return time.Since(start)
}
