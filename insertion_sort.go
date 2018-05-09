package main

// InsertionSort 插入排序
// 最差时间复杂度：O(n²)
// 平均时间复杂度：O(n²)
// 最佳时间复杂度：O(n)
// 空间复杂度：O(1)
// 稳定性：稳定
func InsertionSort(intSlice []int) {
	length := len(intSlice)

	for i := 1; i < length; i++ { // i指向待排序元素，左侧为有序元素，右侧为待排序元素
		for j := i; j > 0; j-- { // j为有序元素的最大值，i指向的元素向左比较，直到比j指向的元素大
			if intSlice[j] < intSlice[j-1] {
				intSlice[j-1], intSlice[j] = intSlice[j], intSlice[j-1]
			} else {
				break
			}
		}
	}
}
