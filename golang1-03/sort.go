package main

import "fmt"

func main() {
	var arr1 []int = []int{5, 3, 6, 8, 1, 2}
	var arr2 []int = []int{5, 3, 6, 8, 1, 2}
	fmt.Println(bubbleSort(arr1))
	fmt.Println(insertionSort(arr2))

}

func bubbleSort(arr []int) []int {
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			j--
		}
	}
	return arr
}
