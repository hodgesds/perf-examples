package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"

	"github.com/hodgesds/perf-utils"
	"golang.org/x/sys/unix"
)

func randInt64s(n int) []int64 {
	ints := make([]int64, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Int63()
	}
	return ints
}

func randInt64Ps(n int) []*int64 {
	ints := make([]*int64, n)
	for i := 0; i < n; i++ {
		v := rand.Int63()
		ints[i] = &v
	}
	return ints
}

func compareInts(entries []int) {
	for _, entry := range entries {

		int64s := randInt64s(entry)
		profileValue, err := perf.L1Data(
			unix.PERF_COUNT_HW_CACHE_OP_READ,
			unix.PERF_COUNT_HW_CACHE_RESULT_ACCESS,
			func() error {
				sort.SliceStable(int64s, func(i, j int) bool {
					return int64s[i] < int64s[j]
				})
				return nil
			})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("L1 Data Read Hit  sort []int64  size %d: %+v\n", entry, profileValue)

		int64s = randInt64s(entry)
		profileValue, err = perf.L1Data(
			unix.PERF_COUNT_HW_CACHE_OP_READ,
			unix.PERF_COUNT_HW_CACHE_RESULT_MISS,
			func() error {
				sort.SliceStable(int64s, func(i, j int) bool {
					return int64s[i] < int64s[j]
				})
				return nil
			})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("L1 Data Read Miss sort []int64  size %d: %+v\n", entry, profileValue)

		int64Ps := randInt64Ps(entry)
		profileValue, err = perf.L1Data(
			unix.PERF_COUNT_HW_CACHE_OP_READ,
			unix.PERF_COUNT_HW_CACHE_RESULT_ACCESS,
			func() error {
				sort.SliceStable(int64Ps, func(i, j int) bool {
					return *int64Ps[i] < *int64Ps[j]
				})
				return nil
			})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("L1 Data Read Hit  sort []*int64 size %d: %+v\n", entry, profileValue)
		int64Ps = randInt64Ps(entry)
		profileValue, err = perf.L1Data(
			unix.PERF_COUNT_HW_CACHE_OP_READ,
			unix.PERF_COUNT_HW_CACHE_RESULT_MISS,
			func() error {
				sort.SliceStable(int64s, func(i, j int) bool {
					return *int64Ps[i] < *int64Ps[j]
				})
				return nil
			})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("L1 Data Read Miss sort []*int64 size %d: %+v\n", entry, profileValue)
		println("")
	}
}
