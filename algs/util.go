package algs

import (
	"math/rand"
	"sort"
	"time"
)

type MyDoubleArr []float32

func (c MyDoubleArr) Len() int {
	return len(c)
}

func (c MyDoubleArr) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c MyDoubleArr) Less(i, j int) bool {
	return c[i] < c[j]
}

// IsSorted reports whether data is sorted.
func IsSorted(data sort.Interface) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

// Get sort time
func timeUsed(alg string, data sort.Interface) int64 {
	now := time.Now().UnixNano()
	switch alg {
	case "Insertion":
		Insertion(data)
		break
	case "Selection":
		Selection(data)
		break
	case "Shell":
		Shell(data)
		break
	default:
		return 0
	}
	return time.Now().UnixNano() - now
}

/* TimeRandomInput ...
 *  Get random input time for alg
 * alg [string] algorithm name
 * N [int] Input arr size
 * T [int] sort run time
 */
func TimeRandomInput(alg string, N int, T int) int64 {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var arr MyDoubleArr = make(MyDoubleArr, N)
	var total int64 = 0
	for t := 0; t < T; t++ {
		for i := 0; i < N; i++ {
			arr[i] = rand.Float32()
		}
		total += timeUsed(alg, arr)
	}
	return total
}