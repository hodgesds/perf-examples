package main

import (
	"fmt"
	"log"

	"github.com/hodgesds/perf-utils"
	"golang.org/x/sys/unix"
)

func main() {
	profileValue, err := perf.CPUInstructions(func() error {
		testLL(1000)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CPU instructions: %+v\n", profileValue)
	println("")

	entries := []int{1000, 4000, 8000, 16000}
	for _, entry := range entries {
		profileValue, err = perf.MinorPageFaults(func() error {
			testLL(entry)
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Minor Page Faults (%d) entries: %+v\n", entry, profileValue)
	}

	println("")

	sizes := []int{512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304}
	for _, size := range sizes {
		l := testLL(size)
		profileValue, err = perf.L1Data(
			unix.PERF_COUNT_HW_CACHE_OP_READ,
			unix.PERF_COUNT_HW_CACHE_RESULT_MISS,
			func() error {
				l.removeDuplicates()
				return nil
			})
		fmt.Printf("L1 Data Read Miss Remove Duplicates size %d: %+v\n", size, profileValue)
		profileValue, err = perf.L1Data(
			unix.PERF_COUNT_HW_CACHE_OP_READ,
			unix.PERF_COUNT_HW_CACHE_RESULT_ACCESS,
			func() error {
				l.removeDuplicates()
				return nil
			})
		fmt.Printf("L1 Data Read Hit  Remove Duplicates size %d: %+v\n", size, profileValue)
		println("")
	}

	compareInts([]int{256, 512, 1024, 2048, 4096, 8192})
}
