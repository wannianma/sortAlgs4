package algs

import (
	"sort"
)

/*
 * Selection 选择排序
 * 算法思路：
 * 首先，找到数组中最小的那个元素，其次，将它和数组中的第一个元素交换位置。
 * 再次，在剩下的元素中找到最小的元素，将它与数组中的第二个元素交换位置。
 * 如此往复，直到整个数组排序。
 */
func Selection(data sort.Interface) {
	N := data.Len()

	for i := 0; i < N; i++ {
		var Min = i
		for j := i + 1; j < N; j++ {
			if data.Less(j, Min) {
				Min = j
			}
		}
		data.Swap(i, Min)
	}
}
