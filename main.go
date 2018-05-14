package main

import "fmt"

func main() {
	// 冒泡排序
	input := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1}
	BubbleSort(input)
	fmt.Printf("Bubble Sort Result:%v\n", input)
	input = []int{9, 8, 4, 6, 6, 2, 0, 1, 4, 7, 2, 3}
	BubbleSort(input)
	fmt.Printf("Bubble Sort Result:%v\n", input)

	// 选择排序
	input = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1}
	SelectionSort(input)
	fmt.Printf("Selection Sort Result:%v\n", input)
	input = []int{9, 8, 4, 6, 6, 2, 0, 1, 4, 7, 2, 3}
	SelectionSort(input)
	fmt.Printf("Selection Sort Result:%v\n", input)

	// 插入排序
	input = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1}
	InsertionSort(input)
	fmt.Printf("Insertion Sort Result:%v\n", input)
	input = []int{9, 8, 4, 6, 6, 2, 0, 1, 4, 7, 2, 3}
	InsertionSort(input)
	fmt.Printf("Insertion Sort Result:%v\n", input)

	// 快速排序
	input = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1}
	QuickSort(input, 0, len(input))
	fmt.Printf("Quick Sort Result:%v\n", input)
	input = []int{9, 8, 4, 6, 6, 2, 0, 1, 4, 7, 2, 3}
	QuickSort(input, 0, len(input))
	fmt.Printf("Quick Sort Result:%v\n", input)

	// 归并排序
	input = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1}
	MergeSort(input, 0, len(input))
	fmt.Printf("Merge Sort Result:%v\n", input)
	input = []int{9, 8, 4, 6, 6, 2, 0, 1, 4, 7, 2, 3}
	MergeSort(input, 0, len(input))
	fmt.Printf("Merge Sort Result:%v\n", input)

	// 计数排序
	input = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1}
	CountingSort(input)
	fmt.Printf("Counting Sort Result:%v\n", input)
	input = []int{9, 8, 4, 6, 6, 2, 0, 1, 4, 7, 2, 3}
	CountingSort(input)
	fmt.Printf("Counting Sort Result:%v\n", input)

	// 堆排序
	input = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1}
	HeapSort(input)
	fmt.Printf("Heap Sort Result:%v\n", input)
	input = []int{9, 8, 4, 6, 6, 2, 0, 1, 4, 7, 2, 3}
	HeapSort(input)
	fmt.Printf("Heap Sort Result:%v\n", input)
}
