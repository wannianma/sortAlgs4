package algs

import (
	"sort"
)

/**
 * Shell 希尔排序
 * 希尔排序是对插入排序算法的一种改进，交换不相邻的元素以对数组的局部进行排序，
 * 是数组中任意间隔为h的元素都是有序的，并最终用插入排序将局部有序的数组排序。
 *
 * 实现希尔排序的一种方法师对于每个h，用插入排序将h个子数组独立的排序。但因为数组是
 * 相互独立的，一种更简单的方法是在H个子数组中将每个元素交换到比它大的元素之前去，这样
 * 只需要将插入排序中的元素移动距离用1改为h即可。
 *
 * 这样希尔排序的实现就转化为一个类似于插入排序但使用不同增量的过程。
 */
func Shell(data sort.Interface) {
	N := data.Len()
	h := 1

	// 递增序列的选取会影响排序效率
	for h < N/3 {
		h = 3*h + 1 // 1, 4, 13, 40
	}

	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j > h && data.Less(j, j-h); j -= h {
				data.Swap(j, j-h)
			}
		}
		h = h / 3
	}
}
