package main

import (
	"fmt"
	"sortAlgs4/lib"
)

func main() {
	// timeInsert := float64(algs.TimeRandomInput("Insertion", 10000, 50))
	// timeSelect := float64(algs.TimeRandomInput("Shell", 10000, 50))

	// fmt.Printf("time insert %.3fs, time Shell %.3fs\n", (timeInsert / (1e9)), (timeSelect / (1e9)))

	// timeQuick := float64(TimeRandomInput("HeapOfficial", 10, 1))

	// fmt.Printf("time Heap %.3f", (timeQuick / (1e9)))

	timeQuick := float64(lib.TimeRandomInput("Quick", 1000000, 50))

	timeHeap := float64(lib.TimeRandomInput("Heap", 1000000, 50))

	fmt.Printf("time quick %.3fs, time heap %.3fs\n", (timeQuick / (1e9)), (timeHeap / (1e9)))

}
