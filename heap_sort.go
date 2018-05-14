package main

// HeapSort 堆排序(大顶堆)
// 最差时间复杂度：O(nlogn)
// 平均时间复杂度：O(nlogn)
// 最佳时间复杂度：O(nlogn)
// 空间复杂度：O(1)
// 稳定性：不稳定
func HeapSort(intSlice []int) {
	length := len(intSlice)

	// 根节点为intSlice[0]
	for i := length/2 - 1; i >= 0; i-- {
		maxHeapify(intSlice, length, i)
	}

	for i := length - 1; i > 0; i-- {
		intSlice[0], intSlice[i] = intSlice[i], intSlice[0] // 根节点代表的值为整个树的最大值，每次将根节点与最后一个叶子节点互换
		maxHeapify(intSlice, i, 0)                          // 有序性已经打乱，重新下沉根节点，将根节点的值放到正确位置
	}
}

func maxHeapify(intSlice []int, length, cur int) {
	if cur >= length {
		return
	}
	father := cur
	child := cur*2 + 1 //左子节点:i*2+1 右子节点:i*2+2

	for {
		if child >= length {
			break
		}
		// 右子树存在，且右子树比左子树大，则交换保证左子树是最大的
		if ((child + 1) < length) && (intSlice[child] < intSlice[child+1]) {
			intSlice[child], intSlice[child+1] = intSlice[child+1], intSlice[child]
		}
		// 父节点与左子树进行比较，父节点小则交换
		if intSlice[father] < intSlice[child] {
			intSlice[father], intSlice[child] = intSlice[child], intSlice[father]
		}
		father = child
		child = child*2 + 1
	}
}
