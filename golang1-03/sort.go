package main

import "fmt"

func main() {
	var arr []int = []int{5, 3, 6, 8, 1, 2}
	fmt.Println(bubbleSort(arr))
	fmt.Println(insertionSort(arr))
	fmt.Println(arr)

}

func bubbleSort(arr []int) []int {
	n := len(arr)
	nums := make([]int, n)
	copy(nums, arr)
	for i := n; i > 0; i-- {
		for j := 1; j < i; j++ {
			if nums[j-1] > nums[j] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}

func insertionSort(arr []int) []int {
	n := len(arr)
	nums := make([]int, len(arr))
	copy(nums, arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if nums[j-1] > nums[j] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}
