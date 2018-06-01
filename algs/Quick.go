package algs

import (
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
	var i, j = lo, hi + 1

	for true {
		for i = i + 1; data.Less(i, lo); i++ {
			if i == hi {
				break
			}
		}
		for j = j - 1; data.Less(lo, j); j-- {
			if j == lo {
				break
			}
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
