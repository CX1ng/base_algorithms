package main

// CountingSort 计数排序 序列中的元素必须是非负整数
// 最差时间复杂度：O(n)
// 平均时间复杂度：O(n)
// 最佳时间复杂度：O(n)
// 空间复杂度：O(n)
// 稳定性：稳定
func CountingSort(intSlice []int) {
	length := len(intSlice)
	var min, max int
	for i := 0; i < length; i++ {
		if intSlice[i] > max {
			if max < intSlice[i] {
				max = intSlice[i]
			}
			if min > intSlice[i] {
				min = intSlice[i]
			}
		}
	}

	// 按照范围申请空间，不适用于跨度很大的序列排序
	counting := make([]int, max-min+1)
	tmp := make([]int, length)

	// 1
	for i := 0; i < length; i++ {
		offset := intSlice[i] - min // 当前元素在整个去重序列中的偏移量
		counting[offset]++
	}

	// 2
	for i := 1; i < len(counting); i++ {
		counting[i] += counting[i-1] // 不比元素(i+m)大的元素个数
	}

	// 3
	for i := length - 1; i >= 0; i-- {
		offset := intSlice[i] - min
		counting[offset] -= 1 //-1考虑存在相同元素的情况。先减而不是后减：数组从0开始，而计数却是从1开始
		tmp[counting[offset]] = intSlice[i]
	}

	copy(intSlice, tmp)
}
