package lib

import "sort"

// PQInterface 是一个基于最大堆的优先队列
// 包括 sort.Interface中的Len, Swap, Less三个方法外，
// 还有：
// 		Insert 插入元素方法
//		DelMax 删除最大元素方法
type PQInterface interface {
	sort.Interface
	Insert(x interface{}) // add x as element Len()
	DelMax() interface{}  // remove and return element Len() - 1.
}

func PQInit() {
}

// swim 由下至上有序化
func swim(h PQInterface, k int) {
	for k > 0 && h.Less(k/2, k) {
		h.Swap(k/2, k)
		k = k / 2
	}
}

// sink 由上之下的堆有序化
func sink(h PQInterface, k int) {
	N := h.Len()
	for 2*k <= N {
		j := 2 * k
		if j <= N && h.Less(j, j+1) {
			j++
		}
		if !h.Less(k, j) {
			break
		}
		h.Swap(k, j)
		k = j
	}
}
