package main

import (
	"testing"
)

func TestBubbleSortOrder(t *testing.T) {
	var arr []int = []int{5, 3, 6}
	sarr := bubbleSort(arr)
	if sarr[0] != 3 && sarr[1] != 5 && sarr[2] != 6 {
		t.Errorf("Sorting order in bubble sort is not correct")
	}
}

func TestBubbleSortLen(t *testing.T) {
	var arr []int = []int{5, 3, 6}
	sarr := bubbleSort(arr)
	if len(sarr) != 3 {
		t.Errorf("The output len is not correct")
	}
}

func TestInsertionSortOrder(t *testing.T) {
	var arr []int = []int{5, 3, 6}
	sarr := insertionSort(arr)
	if sarr[0] != 3 && sarr[1] != 5 && sarr[2] != 6 {
		t.Errorf("Sorting order in insertion sort is not correct")
	}
}

func TestInsertionSortLen(t *testing.T) {
	var arr []int = []int{5, 3, 6}
	sarr := bubbleSort(arr)
	if len(sarr) != 3 {
		t.Errorf("The output len is not correct")
	}
}
