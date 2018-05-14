package main

// SelectionSort 选择排序
// 最差时间复杂度：O(n²)
// 平均时间复杂度：O(n²)
// 最佳时间复杂度：O(n²)
// 空间复杂度：O(1)
// 稳定性：不稳定 [3, 2, 3, 1]
func SelectionSort(intSlice []int) {
	length := len(intSlice)

	for i := 0; i < length-1; i++ { // 控制当前遍历中最小元素的存放下标
		for j := i + 1; j < length; j++ { // 前i个为有序元素，从无序元素中查找最小元素放到i指向的位置
			if intSlice[j] < intSlice[i] {
				intSlice[i], intSlice[j] = intSlice[j], intSlice[i]
			}
		}
	}
}
