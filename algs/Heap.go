package algs

import (
	"container/heap"
	"fmt"
	"sort"
)

// PQInit 构造有序堆
func PQInit(h sort.Interface) {
	N := h.Len()
	for k := N/2 - 1; k >= 0; k-- {
		sink(h, k, N)
	}
}

// swim 由下至上有序化
func swim(h sort.Interface, k int) {
	for k > 0 && h.Less(k/2, k) {
		h.Swap(k/2, k)
		k = k / 2
	}
}

// sink 由上之下的堆有序化
func sink(h sort.Interface, k int, N int) {
	// 数组边界问题
	for 2*k+1 < N-1 {
		j := 2*k + 1
		if j < N-1 && h.Less(j, j+1) {
			j++
		}
		if !h.Less(k, j) {
			break
		}
		h.Swap(k, j)
		k = j
	}
}

/*
 * 堆排序 Heap Sort
 *
 * 堆排序分为两个阶段，在堆得构造阶段，将原始数组重新组织安排进一个堆中；
 * 然后是下沉排序阶段，从堆中按递减顺序取出所有元素并取得排序结果。
 *
 * 1 堆构造阶段：
 *   使用官方heap.Init构建堆
 * 2 下沉获取最小元素
 * 	 使用heap.Pop得到最小元素
 */
func HeapOffical(data heap.Interface) {
	heap.Init(data)
	for data.Len() > 0 {
		fmt.Printf("%f ", heap.Pop(data))
	}
	fmt.Printf("\n")
}

/*
 * 堆排序 Heap Sort
 *
 * 堆排序分为两个阶段，在堆得构造阶段，将原始数组重新组织安排进一个堆中；
 * 然后是下沉排序阶段，从堆中按递减顺序取出所有元素并取得排序结果。
 *
 */
func Heap(data sort.Interface) {
	// 堆构造阶段
	PQInit(data)
	N := data.Len()
	for N > 0 {
		data.Swap(0, N-1)
		N--
		sink(data, 0, N)
	}
}
