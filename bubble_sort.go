package main

// BubbleSort 冒泡排序
// 最差时间复杂度：O(n²)
// 平均时间复杂度：O(n²)
// 最佳时间复杂度：O(n)
// 空间复杂度：O(1)
// 稳定性：稳定
func BubbleSort(intSlice []int) {
	length := len(intSlice)

	for i := 1; i < length; i++ { //控制有序元素个数
		for j := 0; j < length-i; j++ { // 控制需要比较的元素下标 (i vs i+1)
			if intSlice[j] > intSlice[j+1] {
				intSlice[j], intSlice[j+1] = intSlice[j+1], intSlice[j]
			}
		}
	}
}
