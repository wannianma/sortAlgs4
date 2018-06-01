package algs

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/**
 * Quick 快速排序
 *
 */
func Quick(data sort.Interface) {
	shuffle(data)
	fmt.Println("shuffled")
	qsort(data, 0, data.Len()-1)
}

func qsort(data sort.Interface, lo int, hi int) {
	if hi > lo {
		j := partition(data, lo, hi)
		qsort(data, lo, j-1)
		qsort(data, j+1, hi)
	}
}

func partition(data sort.Interface, lo int, hi int) int {
	var i, j = lo + 1, hi

	for true {
		for data.Less(i, lo) {
			if i == hi {
				break
			}
			i++
		}
		for data.Less(lo, j) {
			if j == lo {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		data.Swap(i, j)
	}
	data.Swap(lo, j)
	return j
}

func shuffle(data sort.Interface) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	n := data.Len()
	for n > 0 {
		randIndex := r.Intn(n)
		data.Swap(n-1, randIndex)
		n--
	}
}
