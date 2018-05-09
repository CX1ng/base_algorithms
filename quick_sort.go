package main

// QuickSort 快速排序
// 最差时间复杂度：O(n²)
// 平均时间复杂度：O(nlogn)
// 最佳时间复杂度：O(nlogn)
// 空间复杂度：O(1)
// 稳定性：不稳定(?)
func QuickSort(intSlice []int, min, max int) {
	// intSlice[min:max] 一个元素已经有序，至少有两个元素
	if (max - min) < 2 {
		return
	}

	// intSlice[min:cur]的元素小于cur
	// intSlice[curl:i]的元素等于或大于cur
	// intSlice[i:max]的元素等待排序
	cur := min
	for i := cur + 1; i < max; i++ {
		if intSlice[i] < intSlice[cur] {
			intSlice[i], intSlice[cur] = intSlice[cur], intSlice[i]
			cur++
		}
	}

	// 经过一次排序后，cur左侧的元素都比它小，右侧的元素都比它大
	// cur所指向的元素已经处于有序位置，之后的递归不需要再对cur指向的元素进行排序
	QuickSort(intSlice, min, cur)
	QuickSort(intSlice, cur+1, max)
}
