package algs

import (
	"container/heap"
	"fmt"
	"sortAlgs4/lib"
)

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
 * 1 堆构造阶段：
 *   使用官方heap.Init构建堆
 * 2 下沉获取最小元素
 * 	 使用heap.Pop得到最小元素
 */
func Heap(data lib.PQInterface) {

}
