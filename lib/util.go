package lib

import (
	"math/rand"
	"sort"
	"sortAlgs4/algs"
	"time"
)

// MyDoubleArr self use type
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

// Push heap.Interface only method Push
func (c *MyDoubleArr) Push(x interface{}) {
	*c = append(*c, x.(float32))
}

// Pop heap.Interface only method Pop()
func (c *MyDoubleArr) Pop() interface{} {
	old := *c
	N := len(old)
	x := old[N-1]
	*c = old[0 : N-1]
	return x
}

// Insert insert a new Item to arr
func (c *MyDoubleArr) Insert(x interface{}) {
	*c = append(*c, x.(float32))
}

// DelMax
func (c *MyDoubleArr) DelMax() interface{} {
	old := *c
	N := len(old)
	x := old[N-1]
	*c = old[0 : N-1]
	return x
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
func timeUsed(alg string, data MyDoubleArr) int64 {
	now := time.Now().UnixNano()
	switch alg {
	case "Insertion":
		algs.Insertion(data)
		break
	case "Selection":
		algs.Selection(data)
		break
	case "Shell":
		algs.Shell(data)
		break
	case "Quick":
		algs.Quick(data)
		break
	case "HeapOffical":
		algs.HeapOffical(&data)
		break
	case "Heap":
		algs.Heap(&data)
		break
	default:
		return 0
	}
	return time.Now().UnixNano() - now
}

// TimeRandomInput Get random input time for alg
// alg [string] algorithm name
// N [int] Input arr size
// T [int] sort run time
func TimeRandomInput(alg string, N int, T int) int64 {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make(MyDoubleArr, N)
	var total int64 = 0
	for t := 0; t < T; t++ {
		for i := 0; i < N; i++ {
			arr[i] = rand.Float32()
		}
		total += timeUsed(alg, arr)
	}
	return total
}
