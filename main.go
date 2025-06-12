package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func benchmarkLookup(maxN, step, trials int) ([]int, []float64, []float64) {
	ns := []int{}
	arrayTimes := []float64{}
	mapTimes := []float64{}

	for n := 1; n <= maxN; n += step {
		dataSlice := make([]int, n)
		dataMap := make(map[int]bool, n)
		for i := range n {
			dataSlice[i] = i
			dataMap[i] = true
		}
		target := n - 1 // worst-case for array

		// Time slice lookup (linear scan)
		start := time.Now()
		for range trials {
			_ = contains(dataSlice, target)
		}
		duration := time.Since(start)
		arrayTimes = append(arrayTimes, duration.Seconds()/float64(trials))

		// Time map lookup
		start = time.Now()
		for i := 0; i < trials; i++ {
			_ = dataMap[target]
		}
		duration = time.Since(start)
		mapTimes = append(mapTimes, duration.Seconds()/float64(trials))

		ns = append(ns, n)
	}

	return ns, arrayTimes, mapTimes
}

func contains(slice []int, target int) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}

func writeCSV(filename string, ns []int, arrayTimes, mapTimes []float64) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"n", "array_lookup", "map_lookup"})

	for i := range ns {
		writer.Write([]string{
			strconv.Itoa(ns[i]),
			fmt.Sprintf("%.10f", arrayTimes[i]),
			fmt.Sprintf("%.10f", mapTimes[i]),
		})
	}

	return nil
}

func main() {
	maxN := 40
	step := 1
	trials := 1000000

	ns, arrayTimes, mapTimes := benchmarkLookup(maxN, step, trials)
	err := writeCSV("benchmark.csv", ns, arrayTimes, mapTimes)
	if err != nil {
		fmt.Println("Error writing CSV:", err)
	} else {
		fmt.Println("Benchmark results saved to benchmark.csv")
	}
}
