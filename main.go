package main

import (
	"fmt"
	"sortalgs4/algs"
)

func main() {
	timeInsert := float64(algs.TimeRandomInput("Insertion", 10000, 50))
	timeSelect := float64(algs.TimeRandomInput("Shell", 10000, 50))

	fmt.Printf("time insert %.3fs, time Shell %.3fs\n", (timeInsert / (1e9)), (timeSelect / (1e9)))

}
