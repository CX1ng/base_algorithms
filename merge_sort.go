package main

// MergeSort 归并排序
// 最差时间复杂度：O(nlogn)
// 平均时间复杂度：O(nlogn)
// 最佳时间复杂度：O(nlogn)
// 稳定性：稳定
func MergeSort(intSlice []int, min, max int) {
	// 归: 分解切片，当切片元素大小为1时，代表已经有序，开始进行并操作
	if min+1 >= max {
		return
	}

	mid := (max + min) / 2
	MergeSort(intSlice, min, mid)
	MergeSort(intSlice, mid, max)

	// 并：已经无法分解为更小的粒度
	tmp := make([]int, max-min)
	i := min
	j := mid
	offset := 0
	for {
		if i >= mid {
			copy(tmp[offset:], intSlice[j:max])
			break
		} else if j >= max {
			copy(tmp[offset:], intSlice[i:mid])
			break
		} else if intSlice[i] < intSlice[j] {
			tmp[offset] = intSlice[i]
			i++
		} else {
			tmp[offset] = intSlice[j]
			j++
		}
		offset++
	}
	copy(intSlice[min:max], tmp)
}
